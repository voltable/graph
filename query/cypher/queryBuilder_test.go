package cypher_test

import (
	"testing"

	"fmt"

	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/query/cypher"
	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
	"github.com/RossMerr/Caudex.Graph/vertices"
)

func Test_ToQueryPath(t *testing.T) {
	v := &ast.VertexPatn{Variable: "you"}
	want := &query.QueryPath{}
	var b bool

	toPredicateVertex := func(*ast.VertexPatn) query.PredicateVertex {
		return func(v *vertices.Vertex) bool {
			return b
		}
	}

	toPredicateEdge := func(patn *ast.EdgePatn) query.PredicateEdge {
		return func(*vertices.Edge) bool {
			return b
		}
	}

	want.SetNext(&query.PredicateVertexPath{PredicateVertex: toPredicateVertex(v)})

	got, _ := cypher.ToQueryPath(&ast.MatchStmt{Pattern: v}, toPredicateVertex, toPredicateEdge)

	if fmt.Sprintf("%v", got.Next()) != fmt.Sprintf("%v", want.Next()) {
		t.Errorf("ToQueryPath() = %w, want %w", got.Next(), want.Next())
	}

}
