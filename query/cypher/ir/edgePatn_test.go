package ir_test

import (
	"testing"

	"github.com/RossMerr/Caudex.Graph/query/cypher/ir"
	"github.com/RossMerr/Caudex.Graph/vertices"
)

func Test_ToPredicateEdge(t *testing.T) {
	properties := make(map[string]interface{})
	properties["name"] = "test"
	eb := &ir.EdgeBodyStmt{Type: "Person", Properties: properties}
	vp := ir.EdgePatn{Body: eb}
	predicateEdgePath := vp.ToPredicateEdgePath()

	v, _ := vertices.NewEdge()
	v.SetRelationshipType("Person")
	v.SetProperty("name", "test")
	_, result := predicateEdgePath.PredicateEdge(v, 0)

	if result != true {
		t.Errorf("predicate failed")
	}
}

func Test_ToPredicateEdgeTypeFail(t *testing.T) {
	eb := &ir.EdgeBodyStmt{}
	vp := ir.EdgePatn{Body: eb}
	predicateEdgePath := vp.ToPredicateEdgePath()

	v, _ := vertices.NewEdge()
	v.SetRelationshipType("Person")
	_, result := predicateEdgePath.PredicateEdge(v, 0)

	if result == true {
		t.Errorf("predicate failed")
	}
}

func Test_ToPredicateEdgePropertiesFail(t *testing.T) {
	properties := make(map[string]interface{})
	properties["name"] = "test"
	eb := &ir.EdgeBodyStmt{Properties: properties}
	vp := ir.EdgePatn{Body: eb}
	predicateEdgePath := vp.ToPredicateEdgePath()

	v, _ := vertices.NewEdge()
	v.SetProperty("name", "hello")

	_, result := predicateEdgePath.PredicateEdge(v, 0)

	if result == true {
		t.Errorf("predicate failed")
	}
}

func Test_ToPredicateEdgePropertiesFailEmpty(t *testing.T) {
	properties := make(map[string]interface{})
	properties["name"] = "test"
	eb := &ir.EdgeBodyStmt{Properties: properties}
	vp := ir.EdgePatn{Body: eb}
	predicateEdgePath := vp.ToPredicateEdgePath()

	v, _ := vertices.NewEdge()

	_, result := predicateEdgePath.PredicateEdge(v, 0)

	if result == true {
		t.Errorf("predicate failed")
	}
}
