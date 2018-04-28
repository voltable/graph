package triplestore_test

import (
	"testing"

	"github.com/RossMerr/Caudex.Graph"
	"github.com/RossMerr/Caudex.Graph/query"

	"github.com/RossMerr/Caudex.Graph/triplestore"
)

// (n:Person {name: "test"})
func Test_ToPredicateVertex(t *testing.T) {
	properties := make(map[string]interface{})
	properties["name"] = "test"

	vp := &triplestore.VertexPatn{Variable: "n", Label: "Person", Properties: properties}

	predicateVertexPath := vp.ToPredicateVertexPath()

	v, _ := graph.NewVertex()
	v.SetLabel("Person")
	v.SetProperty("name", "test")
	_, result := predicateVertexPath.PredicateVertex(v)

	if result != query.Matched {
		t.Errorf("predicate failed")
	}
}

func Test_ToPredicateVertexLabelFail(t *testing.T) {
	properties := make(map[string]interface{})

	vp := triplestore.VertexPatn{Variable: "n", Label: "World", Properties: properties}
	predicateVertexPath := vp.ToPredicateVertexPath()

	v, _ := graph.NewVertex()
	v.SetLabel("Person")
	_, result := predicateVertexPath.PredicateVertex(v)

	if result == query.Visiting {
		t.Errorf("predicate failed")
	}
}

func Test_ToPredicateVertexPropertiesFail(t *testing.T) {

	properties := make(map[string]interface{})
	properties["name"] = "test"

	vp := triplestore.VertexPatn{Variable: "n", Label: "Person", Properties: properties}
	predicateVertexPath := vp.ToPredicateVertexPath()

	v, _ := graph.NewVertex()
	v.SetLabel("Person")
	v.SetProperty("name", "hello")

	_, result := predicateVertexPath.PredicateVertex(v)

	if result == query.Visiting {
		t.Errorf("predicate failed")
	}
}

func Test_ToPredicateVertexPropertiesFailEmpty(t *testing.T) {

	properties := make(map[string]interface{})
	properties["name"] = "test"

	vp := triplestore.VertexPatn{Variable: "n", Label: "Person", Properties: properties}
	predicateVertexPath := vp.ToPredicateVertexPath()

	v, _ := graph.NewVertex()
	v.SetLabel("Person")

	_, result := predicateVertexPath.PredicateVertex(v)

	if result == query.Visiting {
		t.Errorf("predicate failed")
	}
}
