package ast_test

import (
	"testing"

	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
	"github.com/RossMerr/Caudex.Graph/vertices"
)

func Test_ToPredicateVertex(t *testing.T) {
	vp := ast.VertexPatn{}
	predicate := vp.ToPredicateVertex()

	v := &vertices.Vertex{}
	result := predicate(v)

	if result != true {
		t.Errorf("predicate failed")
	}
}
