package triplestore_test

import (
	"testing"

	"github.com/RossMerr/Caudex.Graph"
	"github.com/RossMerr/Caudex.Graph/query"

	"github.com/RossMerr/Caudex.Graph/triplestore"
)

func Test_ToPredicateEdge(t *testing.T) {
	properties := make(map[string]interface{})
	properties["name"] = "test"
	eb := &triplestore.EdgeBodyStmt{Type: "Person", Properties: properties}
	vp := triplestore.EdgePatn{Body: eb}
	predicateEdgePath := vp.ToPredicateEdgePath()

	v, _ := graph.NewEdge()
	v.SetRelationshipType("Person")
	v.SetProperty("name", "test")
	_, result := predicateEdgePath.PredicateEdge(v, 0)

	if result != query.Matching {
		t.Errorf("predicate failed")
	}
}

func Test_ToPredicateEdgeLengthMinimumFail(t *testing.T) {
	eb := &triplestore.EdgeBodyStmt{LengthMinimum: 2}
	vp := triplestore.EdgePatn{Body: eb}
	predicateEdgePath := vp.ToPredicateEdgePath()

	v, _ := graph.NewEdge()
	v.SetRelationshipType("Person")
	_, result := predicateEdgePath.PredicateEdge(v, 1)

	if result != query.Visiting {
		t.Errorf("predicate failed")
	}
}

func Test_ToPredicateEdgeLengthMaximumFail(t *testing.T) {
	eb := &triplestore.EdgeBodyStmt{LengthMaximum: 0}
	vp := triplestore.EdgePatn{Body: eb}
	predicateEdgePath := vp.ToPredicateEdgePath()

	v, _ := graph.NewEdge()
	v.SetRelationshipType("Person")
	_, result := predicateEdgePath.PredicateEdge(v, 1)

	if result != query.Failed {
		t.Errorf("predicate failed")
	}
}

func Test_ToPredicateEdgeTypeFail(t *testing.T) {
	eb := &triplestore.EdgeBodyStmt{Type: "NotAPerson"}
	vp := triplestore.EdgePatn{Body: eb}
	predicateEdgePath := vp.ToPredicateEdgePath()

	v, _ := graph.NewEdge()
	v.SetRelationshipType("Person")
	_, result := predicateEdgePath.PredicateEdge(v, 0)

	if result != query.Failed {
		t.Errorf("predicate failed")
	}
}

func Test_ToPredicateEdgePropertiesFail(t *testing.T) {
	properties := make(map[string]interface{})
	properties["name"] = "test"
	eb := &triplestore.EdgeBodyStmt{Properties: properties}
	vp := triplestore.EdgePatn{Body: eb}
	predicateEdgePath := vp.ToPredicateEdgePath()

	v, _ := graph.NewEdge()
	v.SetProperty("name", "hello")

	_, result := predicateEdgePath.PredicateEdge(v, 0)

	if result != query.Failed {
		t.Errorf("predicate failed")
	}
}

func Test_ToPredicateEdgePropertiesFailEmpty(t *testing.T) {
	properties := make(map[string]interface{})
	properties["name"] = "test"
	eb := &triplestore.EdgeBodyStmt{Properties: properties}
	vp := triplestore.EdgePatn{Body: eb}
	predicateEdgePath := vp.ToPredicateEdgePath()

	v, _ := graph.NewEdge()

	_, result := predicateEdgePath.PredicateEdge(v, 0)

	if result != query.Failed {
		t.Errorf("predicate failed")
	}
}
