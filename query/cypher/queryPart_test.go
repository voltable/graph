package cypher_test

import (
	"testing"

	"github.com/RossMerr/Caudex.Graph/expressions"
	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/query/cypher"
	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
	"github.com/RossMerr/Caudex.Graph/query/cypher/ir"
	"github.com/RossMerr/Caudex.Graph/vertices"
)

func Test_ToQueryPath(t *testing.T) {
	edgePatn := &ir.EdgePatn{Body: &ir.EdgeBodyStmt{LengthMinimum: 2, LengthMaximum: 5}}
	vertexPatn := &ir.VertexPatn{Variable: "bar", Edge: edgePatn}
	wherePatn := &ast.WhereStmt{Predicate: ast.NewComparisonExpr(expressions.EQ, &ast.PropertyStmt{Variable: "n", Value: "name"}, &ast.Ident{Data: "foo"})}
	match := &ast.MatchStmt{Pattern: vertexPatn, Next: wherePatn}

	var b bool
	toPredicateVertex := func(*ir.VertexPatn) query.PredicateVertex {
		return func(v *vertices.Vertex) (string, bool) {
			return "", b
		}
	}

	toPredicateEdge := func(patn *ir.EdgePatn) query.PredicateEdge {
		return func(e *vertices.Edge) (string, bool) {
			return "", b
		}
	}

	want := &cypher.Root{}
	vertexPath := &query.PredicateVertexPath{PredicateVertex: toPredicateVertex(vertexPatn)}
	vertexPath.SetNext(&query.PredicateEdgePath{PredicateEdge: toPredicateEdge(edgePatn)})
	want.SetNext(vertexPath)

	parts, _ := cypher.NewParts().ToQueryPart(match)
	partOne := parts[0]

	if partOne.Where == nil {
		t.Errorf("Where statment not matched")
	}

	got := partOne.Path
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

	var tests = []struct {
		c      ast.Stmt
		result bool
	}{
		{
			c:      ast.DeleteStmt{},
			result: true,
		}, {
			c:      ast.CreateStmt{},
			result: true,
		}, {
			c:      ast.OptionalMatchStmt{},
			result: true,
		}, {
			c:      ast.MatchStmt{},
			result: true,
		}, {
			c:      ast.WhereStmt{},
			result: true,
		},
	}

	for i, tt := range tests {
		_, ok := cypher.IsPattern(&tt.c)
		if ok == tt.result {
			t.Errorf("%d. comparison mismatch:\n %v\n\n", i, tt.c)
		}
	}
}

// errstring returns the string representation of an error.
func errstring(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}
