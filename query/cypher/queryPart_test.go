package cypher_test

import (
	"testing"

	"github.com/RossMerr/Caudex.Graph"
	"github.com/RossMerr/Caudex.Graph/expressions"
	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/query/cypher"
	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
	"github.com/RossMerr/Caudex.Graph/query/cypher/ir"
)

func Test_ToQueryPath(t *testing.T) {
	edgePatn := &ir.EdgePatn{Body: &ir.EdgeBodyStmt{LengthMinimum: 2, LengthMaximum: 5}}
	vertexPatn := &ir.VertexPatn{Variable: "bar", Edge: edgePatn}
	wherePatn := &ast.WhereStmt{Predicate: ast.NewComparisonExpr(expressions.EQ, &ast.PropertyStmt{Variable: "n", Value: "name"}, &ast.Ident{Data: "foo"})}
	match := &ast.MatchStmt{Pattern: vertexPatn, Next: wherePatn}

	var b bool
	toPredicateVertex := func(*ir.VertexPatn) query.PredicateVertex {
		return func(v *graph.Vertex) (string, query.Traverse) {
			if b {
				return "", query.Matched
			} else {
				return "", query.Failed

			}
		}
	}

	toPredicateEdge := func(patn *ir.EdgePatn) query.PredicateEdge {
		return func(e *graph.Edge, depth uint) (string, query.Traverse) {
			if b {
				return "", query.Matched
			} else {
				return "", query.Failed

			}
		}
	}

	want := make([]interface{}, 0)

	want = append(want, &query.PredicateVertexPath{PredicateVertex: toPredicateVertex(vertexPatn)})
	want = append(want, &query.PredicateEdgePath{PredicateEdge: toPredicateEdge(edgePatn)})

	parts, _ := cypher.NewParts().ToQueryPart(match)
	partOne := parts[0]

	if partOne.Where == nil {
		t.Errorf("Where statment not matched")
	}

	got := partOne.Predicates
	e := got[0]
	v, _ := e.(*query.PredicateVertexPath)
	if v == nil {
		t.Errorf("VertexNext")
	}

	e = got[1]
	en, _ := e.(*query.PredicateEdgePath)
	if en == nil {
		t.Errorf("EdgeNext")
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
