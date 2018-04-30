package query_test

import (
	"container/list"
	"errors"
	"reflect"
	"testing"

	"github.com/RossMerr/Caudex.Graph"
	"github.com/RossMerr/Caudex.Graph/expressions"
	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/query/cypher/ir"
)

var (
	errRecordNotFound = errors.New("Record Not found")
)

var (
	drw, _ = graph.NewVertex()
	cns, _ = graph.NewVertex()
	asp, _ = graph.NewVertex()
	bne, _ = graph.NewVertex()
	syd, _ = graph.NewVertex()
	cbr, _ = graph.NewVertex()
	mel, _ = graph.NewVertex()
	adl, _ = graph.NewVertex()
	per, _ = graph.NewVertex()
)

// https://neo4j.com/blog/graph-search-algorithm-basics/
// https://s3.amazonaws.com/dev.assets.neo4j.com/wp-content/uploads/20160715114833/Dijkstras-Algorithm-routing11.png
func init() {
	drw.SetLabel("drw")
	cns.SetLabel("cns")
	asp.SetLabel("asp")
	bne.SetLabel("bne")
	syd.SetLabel("syd")
	cbr.SetLabel("cbr")
	mel.SetLabel("mel")
	adl.SetLabel("adl")
	per.SetLabel("per")

	drw.AddEdgeWeight(cns, float64(30))
	drw.AddEdgeWeight(asp, float64(15))
	drw.AddEdgeWeight(per, float64(48))

	cns.AddEdgeWeight(drw, float64(30))
	cns.AddEdgeWeight(asp, float64(24))
	cns.AddEdgeWeight(bne, float64(22))

	asp.AddEdgeWeight(drw, float64(15))
	asp.AddEdgeWeight(cns, float64(24))
	asp.AddEdgeWeight(bne, float64(31))
	asp.AddEdgeWeight(cbr, float64(15))
	asp.AddEdgeWeight(adl, float64(15))

	bne.AddEdgeWeight(cns, float64(22))
	bne.AddEdgeWeight(asp, float64(31))
	bne.AddEdgeWeight(syd, float64(9))

	syd.AddEdgeWeight(bne, float64(9))
	syd.AddEdgeWeight(mel, float64(12))
	syd.AddEdgeWeight(cbr, float64(4))

	cbr.AddEdgeWeight(mel, float64(6))
	cbr.AddEdgeWeight(syd, float64(4))
	cbr.AddEdgeWeight(asp, float64(15))

	mel.AddEdgeWeight(syd, float64(12))
	mel.AddEdgeWeight(cbr, float64(6))
	mel.AddEdgeWeight(adl, float64(8))

	adl.AddEdgeWeight(mel, float64(8))
	adl.AddEdgeWeight(asp, float64(15))
	adl.AddEdgeWeight(per, float64(32))

	per.AddEdgeWeight(adl, float64(32))
	per.AddEdgeWeight(drw, float64(48))
}

func Test_UniformCostSearch(t *testing.T) {
	g := AustraliaGraph()

	toPredicateVertex := toPredicateVertex(t)
	vPath := &query.PredicateVertexPath{PredicateVertex: toPredicateVertex(nil)}

	toPredicateEdge := toPredicateEdge(t)
	ePath := &query.PredicateEdgePath{PredicateEdge: toPredicateEdge(nil)}

	path := make([]interface{}, 0)
	path = append(path, vPath)
	path = append(path, ePath)
	path = append(path, vPath)
	path = append(path, ePath)
	path = append(path, vPath)
	path = append(path, ePath)
	path = append(path, vPath)
	path = append(path, ePath)
	path = append(path, vPath)

	plan := query.NewPlan(g)
	it, err := plan.SearchPlan(g.ForEachTest(), path)
	if err != nil {
		t.Fatalf("Travers failed %+v", err)
	}
	result := ToIterator(it)

	count := len(result)
	if count != 5 {
		t.Fatalf("Expected result count to be %+v but was %+v", 5, count)
	}

	if !reflect.DeepEqual(result[0], syd) {
		t.Fatalf("Expected syd: \n%+v \nbut was \n%+v", syd, result[0])
	}

	if !reflect.DeepEqual(result[1], cbr) {
		t.Fatalf("Expected cbr: \n%+v \nbut was \n%+v", cbr, result[1])
	}

	if !reflect.DeepEqual(result[2], mel) {
		t.Fatalf("Expected mel: \n%+v \nbut was \n%+v", mel, result[2])
	}

	if !reflect.DeepEqual(result[3], adl) {
		t.Fatalf("Expected adl: \n%+v \nbut was \n%+v", adl, result[3])
	}

	if !reflect.DeepEqual(result[4], per) {
		t.Fatalf("Expected per: \n%+v \nbut was \n%+v", per, result[4])
	}
}

func AustraliaGraph() *StorageEngine {
	g := &StorageEngine{vertices: make(map[string]graph.Vertex)}
	g.Create(drw, cns, asp, bne, syd, cbr, mel, adl, per)
	return g
}

type StorageEngine struct {
	vertices  map[string]graph.Vertex
	keys      []string
	traversal query.Traversal
}

// Create adds a array of vertices to the persistence
func (se *StorageEngine) Create(c ...*graph.Vertex) error {
	for _, v := range c {
		se.vertices[v.ID()] = *v
		se.keys = append(se.keys, v.ID())
	}

	return nil
}

// Delete the array of vertices from the persistence
func (se *StorageEngine) Delete(c ...*graph.Vertex) error {
	return nil
}

// Update the array of vertices from the persistence
func (se *StorageEngine) Update(c ...*graph.Vertex) error {
	se.Create(c...)
	return nil
}

func (se *StorageEngine) Query(str string) (*graph.Query, error) {
	return nil, nil
}

func (se *StorageEngine) Fetch(id string) (*graph.Vertex, error) {
	if v, ok := se.vertices[id]; ok {
		return &v, nil
	} else {
		return nil, errRecordNotFound
	}
}

func (se *StorageEngine) Close() {

}

func (se *StorageEngine) ForEach() graph.Iterator {
	position := 0
	length := len(se.keys)
	return func() (item interface{}, ok bool) {
		if position < length {
			key := se.keys[position]
			v := se.vertices[key]
			position = position + 1
			return &v, true
		}
		return nil, false
	}
}

func (se *StorageEngine) ForEachTest() query.IteratorFrontier {
	state := false
	return func() (item *query.Frontier, ok bool) {
		state = expressions.XORSwap(state)
		f := query.NewFrontier(syd, "")
		return &f, state
	}
}

func ToIterator(i query.IteratorFrontier) []*graph.Vertex {
	results := make([]*graph.Vertex, 0)

	for frontier, ok := i(); ok; frontier, ok = i() {
		if frontier.Len() > 0 {
			parts := frontier.OptimalPath()
			for _, i := range parts {
				if v, ok := i.Object.(*graph.Vertex); ok {
					results = append(results, v)
				}
			}
		}
	}
	return results
}

func toPredicateVertex(t *testing.T) func(*ir.VertexPatn) query.PredicateVertex {
	return func(*ir.VertexPatn) query.PredicateVertex {
		return func(v *graph.Vertex) (string, query.Traverse) {
			if v.ID() == per.ID() {
				return "", query.Matched
			} else {
				return "", query.Failed
			}
		}
	}
}

func toPredicateEdge(t *testing.T) func(patn *ir.EdgePatn) query.PredicateEdge {
	return func(patn *ir.EdgePatn) query.PredicateEdge {
		return func(e *graph.Edge, depth uint) (string, query.Traverse) {
			if e.ID() != per.ID() {
				return "", query.Visiting
			} else {
				return "", query.Matching
			}
		}
	}
}

func index(l *list.List, i int) interface{} {
	e := l.Front()
	for index := 1; index < i; index++ {
		e = e.Next()
	}

	return e.Value
}
