package ast_test

import (
	"testing"

	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
	"github.com/RossMerr/Caudex.Graph/vertices"
)

func Test_ToPredicateEdge(t *testing.T) {
	properties := make(map[string]interface{})
	properties["name"] = "test"
	eb := &ast.EdgeBodyStmt{Type: "Person", Properties: properties}
	vp := ast.EdgePatn{Body: eb}
	predicateEdgePath := vp.ToPredicateEdgePath()

	v, _ := vertices.NewEdge()
	v.SetRelationshipType("Person")
	v.SetProperty("name", "test")
	_, result := predicateEdgePath.PredicateEdge(v)

	if result != true {
		t.Errorf("predicate failed")
	}
}

func Test_ToPredicateEdgeTypeFail(t *testing.T) {
	eb := &ast.EdgeBodyStmt{}
	vp := ast.EdgePatn{Body: eb}
	predicateEdgePath := vp.ToPredicateEdgePath()

	v, _ := vertices.NewEdge()
	v.SetRelationshipType("Person")
	_, result := predicateEdgePath.PredicateEdge(v)

	if result == true {
		t.Errorf("predicate failed")
	}
}

func Test_ToPredicateEdgePropertiesFail(t *testing.T) {
	properties := make(map[string]interface{})
	properties["name"] = "test"
	eb := &ast.EdgeBodyStmt{Properties: properties}
	vp := ast.EdgePatn{Body: eb}
	predicateEdgePath := vp.ToPredicateEdgePath()

	v, _ := vertices.NewEdge()
	v.SetProperty("name", "hello")

	_, result := predicateEdgePath.PredicateEdge(v)

	if result == true {
		t.Errorf("predicate failed")
	}
}

func Test_ToPredicateEdgePropertiesFailEmpty(t *testing.T) {
	properties := make(map[string]interface{})
	properties["name"] = "test"
	eb := &ast.EdgeBodyStmt{Properties: properties}
	vp := ast.EdgePatn{Body: eb}
	predicateEdgePath := vp.ToPredicateEdgePath()

	v, _ := vertices.NewEdge()

	_, result := predicateEdgePath.PredicateEdge(v)

	if result == true {
		t.Errorf("predicate failed")
	}
}
