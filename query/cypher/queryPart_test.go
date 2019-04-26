package cypher_test

import (
	"reflect"
	"testing"

	"github.com/voltable/graph/expressions"
	"github.com/voltable/graph/query/cypher"
	"github.com/voltable/graph/query/cypher/ast"
)

func Test_ToQueryPath(t *testing.T) {
	var patn ast.Patn
	wherePatn := &ast.WhereStmt{Predicate: ast.NewComparisonExpr(expressions.EQ, &ast.PropertyStmt{Variable: "n", Value: "name"}, &ast.Ident{Data: "foo"})}
	match := &ast.MatchStmt{Pattern: patn, Next: wherePatn}

	parts, _ := cypher.NewParts().ToQueryPart(match)
	partOne := parts[0]

	if partOne.Where == nil {
		t.Errorf("Where statment not matched")
	}
}

func Test_ToQueryPath_Pattern(t *testing.T) {
	v := &ast.VertexPatn{Edge: &ast.EdgePatn{}}
	wherePatn := &ast.WhereStmt{Predicate: ast.NewComparisonExpr(expressions.EQ, &ast.PropertyStmt{Variable: "n", Value: "name"}, &ast.Ident{Data: "foo"})}
	match := &ast.MatchStmt{Pattern: v, Next: wherePatn}

	parts, _ := cypher.NewParts().ToQueryPart(match)
	partOne := parts[0]

	if partOne.Where == nil {
		t.Errorf("Where statment not matched")
	}
}

func Test_ToQueryPath_Return(t *testing.T) {
	returnPatn := &ast.ReturnStmt{}
	match := &ast.MatchStmt{Next: returnPatn}

	parts, _ := cypher.NewParts().ToQueryPart(match)
	partOne := parts[0]

	if partOne.Return == nil {
		t.Errorf("Return statment not matched")
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

func TestQueryPart_Predicate(t *testing.T) {
	predicate := ast.NewComparisonExpr(expressions.EQ, &ast.PropertyStmt{Variable: "n", Value: "name"}, &ast.Ident{Data: "foo"})
	wherePatn := &ast.WhereStmt{Predicate: predicate}

	qp := &cypher.QueryPart{}
	qp.Where = wherePatn

	result := qp.Predicate()
	if result != predicate {
		t.Errorf("Where predicate not found")
	}
}

func TestQueryPart_Maps(t *testing.T) {
	maps := make([]*ast.ProjectionMapStmt, 0)
	returnPatn := &ast.ReturnStmt{Maps: maps}

	qp := &cypher.QueryPart{}
	qp.Return = returnPatn

	result := qp.Maps()
	if !reflect.DeepEqual(result, maps) {
		t.Errorf("Return maps not found")
	}
}
