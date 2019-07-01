package traversal_test

import (
	"bytes"
	"container/list"
	"fmt"
	"reflect"
	"testing"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	graph "github.com/voltable/graph"
	"github.com/voltable/graph/expressions"
	"github.com/voltable/graph/query"
	"github.com/voltable/graph/query/cypher"
	"github.com/voltable/graph/query/cypher/ast"
	"github.com/voltable/graph/query/cypher/traversal"
	"github.com/voltable/graph/uuid"
	"github.com/voltable/graph/widecolumnstore"
	"github.com/voltable/graph/widecolumnstore/storage/memorydb"
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

	drw.AddDirectedEdgeWeight(cns, float64(30))
	drw.AddDirectedEdgeWeight(asp, float64(15))
	drw.AddDirectedEdgeWeight(per, float64(48))

	cns.AddDirectedEdgeWeight(drw, float64(30))
	cns.AddDirectedEdgeWeight(asp, float64(24))
	cns.AddDirectedEdgeWeight(bne, float64(22))

	asp.AddDirectedEdgeWeight(drw, float64(15))
	asp.AddDirectedEdgeWeight(cns, float64(24))
	asp.AddDirectedEdgeWeight(bne, float64(31))
	asp.AddDirectedEdgeWeight(cbr, float64(15))
	asp.AddDirectedEdgeWeight(adl, float64(15))

	bne.AddDirectedEdgeWeight(cns, float64(22))
	bne.AddDirectedEdgeWeight(asp, float64(31))
	bne.AddDirectedEdgeWeight(syd, float64(9))

	syd.AddDirectedEdgeWeight(bne, float64(9))
	syd.AddDirectedEdgeWeight(mel, float64(12))
	syd.AddDirectedEdgeWeight(cbr, float64(4))

	cbr.AddDirectedEdgeWeight(mel, float64(6))
	cbr.AddDirectedEdgeWeight(syd, float64(4))
	cbr.AddDirectedEdgeWeight(asp, float64(15))

	mel.AddDirectedEdgeWeight(syd, float64(12))
	mel.AddDirectedEdgeWeight(cbr, float64(6))
	mel.AddDirectedEdgeWeight(adl, float64(8))

	adl.AddDirectedEdgeWeight(mel, float64(8))
	adl.AddDirectedEdgeWeight(asp, float64(15))
	adl.AddDirectedEdgeWeight(per, float64(32))

	per.AddDirectedEdgeWeight(adl, float64(32))
	per.AddDirectedEdgeWeight(drw, float64(48))
}

func Test_UniformCostSearch(t *testing.T) {
	g := AustraliaGraph()
	graph := *g.(*query.Graph)

	fmt.Printf("drw: %+v\n", drw.ID())
	fmt.Printf("cns: %+v\n", cns.ID())
	fmt.Printf("asp: %+v\n", asp.ID())
	fmt.Printf("bne: %+v\n", bne.ID())
	fmt.Printf("syd: %+v\n", syd.ID())
	fmt.Printf("cbr: %+v\n", cbr.ID())
	fmt.Printf("mel: %+v\n", mel.ID())
	fmt.Printf("adl: %+v\n", adl.ID())
	fmt.Printf("per: %+v\n", per.ID())

	id := per.ID()
	targetBytes := id[:]

	goal := func(key widecolumnstore.Key) bool {
		return bytes.Equal(targetBytes, key.ID)
	}

	result, err := traversal.UniformCostSearch2(graph.Storage, syd, goal)
	if err != nil {
		t.Fatalf("Expected err to be nil but was %s", err)
	}

	count := len(result)
	if count != 5 {
		t.Fatalf("Expected result count to be %+v but was %+v", 5, count)
	}

	if !reflect.DeepEqual(result[0], syd.ID()) {
		t.Fatalf("Expected syd: \n%+v \nbut was \n%+v", syd.ID(), result[0])
	}

	if !reflect.DeepEqual(result[1], cbr.ID()) {
		t.Fatalf("Expected cbr: \n%+v \nbut was \n%+v", cbr.ID(), result[1])
	}

	if !reflect.DeepEqual(result[2], mel.ID()) {
		t.Fatalf("Expected mel: \n%+v \nbut was \n%+v", mel.ID(), result[2])
	}

	if !reflect.DeepEqual(result[3], adl.ID()) {
		t.Fatalf("Expected adl: \n%+v \nbut was \n%+v", adl.ID(), result[3])
	}

	if !reflect.DeepEqual(result[4], per.ID()) {
		t.Fatalf("Expected per: \n%+v \nbut was \n%+v", per.ID(), result[4])
	}
}

func AustraliaGraph() graph.Graph {

	cypher.RegisterEngine()
	options := graph.NewOptions(cypher.QueryType, memorydb.StorageType)

	g, err := query.NewGraphEngine(options)
	if err != nil {
		fmt.Errorf("Failed to create the storageEngine %v", err)
	}

	g.Create(drw, cns, asp, bne, syd, cbr, mel, adl, per)

	return g
}

func ForEachTest(se graph.Graph) query.IteratorFrontier {
	ok := false
	return func() (*query.Frontier, bool) {
		ok = expressions.XORSwap(ok)
		if ok {
			kv, _ := query.MarshalKeyValue(syd)
			id, err := query.UUID(kv[0])
			if err != nil {
				log.Error(errors.Wrap(err, "ForEachTest"))
			} else {
				f := query.NewFrontier(&id, "")
				return &f, ok
			}
		}
		return nil, ok
	}
}

func ToIterator(i query.IteratorFrontier) []*uuid.UUID {
	results := make([]*uuid.UUID, 0)

	for frontier, ok := i(); ok; frontier, ok = i() {
		if frontier.Len() > 0 {
			parts := frontier.OptimalPath()
			for _, i := range parts {
				results = append(results, i.UUID)
			}
		}
	}
	return results
}

type queryBuilder struct {
	Storage graph.Graph
}

// Predicate(patterns []ast.Patn) (widecolumnstore.Operator, error)
func (qb *queryBuilder) Predicate([]ast.Patn) (widecolumnstore.Operator, error) {
	// path := make([]query.Predicate, 0)
	// path = append(path, qb.toPredicate())
	// path = append(path, qb.toPredicate())
	// path = append(path, qb.toPredicate())
	// path = append(path, qb.toPredicate())
	// path = append(path, qb.toPredicate())
	// path = append(path, qb.toPredicate())
	// path = append(path, qb.toPredicate())
	// path = append(path, qb.toPredicate())
	// path = append(path, qb.toPredicate())

	return nil, nil
}

func (qb *queryBuilder) toPredicate() query.Predicate {
	return func(from, to uuid.UUID, depth int) (string, query.Traverse) {
		if to == uuid.Nil {
			if from == per.ID() {

				return "", query.Matched
			}

			return "", query.Failed
		}

		if to != per.ID() {
			return "", query.Visiting
		}

		return "", query.Matching
	}

}

func index(l *list.List, i int) interface{} {
	e := l.Front()
	for index := 1; index < i; index++ {
		e = e.Next()
	}

	return e.Value
}
