package cypher_test

import (
	"testing"

	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/query/cypher"
	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
	"github.com/RossMerr/Caudex.Graph/vertices"
)

func Test_ToQueryPath(t *testing.T) {
	edgePatn := &ast.EdgePatn{Body: &ast.EdgeBodyStmt{LengthMinimum: 2, LengthMaximum: 5}}
	vertexPatn := &ast.VertexPatn{Variable: "bar", Edge: edgePatn}
	var b bool

	toPredicateVertex := func(*ast.VertexPatn) query.PredicateVertex {
		return func(v *vertices.Vertex) bool {
			return b
		}
	}

	toPredicateEdge := func(patn *ast.EdgePatn) query.PredicateEdge {
		return func(e *vertices.Edge) bool {
			return b
		}
	}

	want := &query.QueryPath{}
	vertexPath := &query.PredicateVertexPath{PredicateVertex: toPredicateVertex(vertexPatn)}
	vertexPath.SetNext(&query.PredicateEdgePath{PredicateEdge: toPredicateEdge(edgePatn)})
	want.SetNext(vertexPath)

	got, _ := cypher.ToQueryPath(&ast.MatchStmt{Pattern: vertexPatn}, toPredicateVertex, toPredicateEdge)

	v, _ := got.Next().(query.VertexNext)
	if v == nil {
		t.Errorf("VertexNext")
	}

	pv, _ := v.(*query.PredicateVertexPath)

	if pv == nil {
		t.Errorf("PredicateVertexPath")
	}

	e, _ := pv.Next().(query.EdgeNext)
	if e == nil {
		t.Errorf("EdgeNext")
	}

	pe, _ := e.(*query.PredicateEdgePath)
	if pe == nil {
		t.Errorf("PredicateEdgePath")
	}

	if pe.Length().Minimum != 2 {
		t.Errorf("Minimum")
	}

	if pe.Length().Maximum != 5 {
		t.Errorf("Maximum")
	}

	last, _ := pe.Next().(query.VertexNext)
	if last != nil {
		t.Errorf("VertexNext")
	}
}

func Test_IsPattern(t *testing.T) {

	if _, ok := cypher.IsPattern(&ast.DeleteStmt{}); !ok {
		t.Errorf("DeleteStmt")
	}

	if _, ok := cypher.IsPattern(&ast.CreateStmt{}); !ok {
		t.Errorf("CreateStmt")
	}

	if _, ok := cypher.IsPattern(&ast.OptionalMatchStmt{}); !ok {
		t.Errorf("OptionalMatchStmt")
	}

	if _, ok := cypher.IsPattern(&ast.MatchStmt{}); !ok {
		t.Errorf("MatchStmt")
	}

	if _, ok := cypher.IsPattern(&ast.WhereStmt{}); ok {
		t.Errorf("WhereStmt")
	}
}
