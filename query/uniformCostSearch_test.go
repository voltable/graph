package query_test

import (
	"container/list"
	"errors"
	"reflect"
	"testing"

	"github.com/RossMerr/Caudex.Graph/enumerables"
	"github.com/RossMerr/Caudex.Graph/expressions"
	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/query/cypher/ir"
	"github.com/RossMerr/Caudex.Graph/vertices"
)

var (
	errRecordNotFound = errors.New("Record Not found")
)

var (
	drw, _ = vertices.NewVertex()
	cns, _ = vertices.NewVertex()
	asp, _ = vertices.NewVertex()
	bne, _ = vertices.NewVertex()
	syd, _ = vertices.NewVertex()
	cbr, _ = vertices.NewVertex()
	mel, _ = vertices.NewVertex()
	adl, _ = vertices.NewVertex()
	per, _ = vertices.NewVertex()
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

	drw.AddEdgeWeight(cns, float32(30))
	drw.AddEdgeWeight(asp, float32(15))
	drw.AddEdgeWeight(per, float32(48))

	cns.AddEdgeWeight(drw, float32(30))
	cns.AddEdgeWeight(asp, float32(24))
	cns.AddEdgeWeight(bne, float32(22))

	asp.AddEdgeWeight(drw, float32(15))
	asp.AddEdgeWeight(cns, float32(24))
	asp.AddEdgeWeight(bne, float32(31))
	asp.AddEdgeWeight(cbr, float32(15))
	asp.AddEdgeWeight(adl, float32(15))

	bne.AddEdgeWeight(cns, float32(22))
	bne.AddEdgeWeight(asp, float32(31))
	bne.AddEdgeWeight(syd, float32(9))

	syd.AddEdgeWeight(bne, float32(9))
	syd.AddEdgeWeight(mel, float32(12))
	syd.AddEdgeWeight(cbr, float32(4))

	cbr.AddEdgeWeight(mel, float32(6))
	cbr.AddEdgeWeight(syd, float32(4))
	cbr.AddEdgeWeight(asp, float32(15))

	mel.AddEdgeWeight(syd, float32(12))
	mel.AddEdgeWeight(cbr, float32(6))
	mel.AddEdgeWeight(adl, float32(8))

	adl.AddEdgeWeight(mel, float32(8))
	adl.AddEdgeWeight(asp, float32(15))
	adl.AddEdgeWeight(per, float32(32))

	per.AddEdgeWeight(adl, float32(32))
	per.AddEdgeWeight(drw, float32(48))
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
	g := &StorageEngine{vertices: make(map[string]vertices.Vertex)}
	g.Create(drw, cns, asp, bne, syd, cbr, mel, adl, per)
	return g
}

type StorageEngine struct {
	vertices  map[string]vertices.Vertex
	keys      []string
	traversal query.Traversal
}

// Create adds a array of vertices to the persistence
func (se *StorageEngine) Create(c ...*vertices.Vertex) error {
	for _, v := range c {
		se.vertices[v.ID()] = *v
		se.keys = append(se.keys, v.ID())
	}

	return nil
}

// Delete the array of vertices from the persistence
func (se *StorageEngine) Delete(c ...*vertices.Vertex) error {
	return nil
}

// Update the array of vertices from the persistence
func (se *StorageEngine) Update(c ...*vertices.Vertex) error {
	se.Create(c...)
	return nil
}

func (se *StorageEngine) Query(str string) (*query.Query, error) {
	return nil, nil
}

func (se *StorageEngine) Fetch(id string) (*vertices.Vertex, error) {
	if v, ok := se.vertices[id]; ok {
		return &v, nil
	} else {
		return nil, errRecordNotFound
	}
}

func (se *StorageEngine) Close() {

}

func (se *StorageEngine) ForEach() enumerables.Iterator {
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

func (se *StorageEngine) ForEachTest() enumerables.Iterator {
	state := false
	return func() (item interface{}, ok bool) {
		state = expressions.XORSwap(state)
		return syd, state
	}
}

func ToIterator(i query.IteratorFrontier) []*vertices.Vertex {
	results := make([]*vertices.Vertex, 0)

	for frontier, ok := i(); ok != query.Failed; frontier, ok = i() {
		if frontier.Len() > 0 {
			vertices := frontier.OptimalPath()
			for _, v := range vertices {
				results = append(results, v.(*query.FrontierVertex).Vertex)
			}
		}
	}
	return results
}

func toPredicateVertex(t *testing.T) func(*ir.VertexPatn) query.PredicateVertex {
	return func(*ir.VertexPatn) query.PredicateVertex {
		return func(v *vertices.Vertex) (string, query.Traverse) {
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
		return func(e *vertices.Edge, depth uint) (string, query.Traverse) {
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
