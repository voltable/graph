package parser_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/RossMerr/Caudex.Graph/expressions"
	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
	"github.com/RossMerr/Caudex.Graph/query/cypher/ir"
	"github.com/RossMerr/Caudex.Graph/query/cypher/parser"
)

// Ensure the parser can parse the right patterns.
func TestParser_Pattern(t *testing.T) {
	var tests = []struct {
		s    string
		stmt ast.Stmt
		err  string
	}{
		{
			s:    `MATCH (you)`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Variable: "you"}},
		},
		{
			s:    `MATCH (:Person)`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Label: "Person"}},
		},
		{
			s:    `MATCH (you:Person)`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Variable: "you", Label: "Person"}},
		},
		{
			s:    `MATCH (you:Person {name:"You"})`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Variable: "you", Label: "Person", Properties: map[string]interface{}{"name": "You"}}},
		},
		{
			s:    `MATCH (you:Person {name:"foo bar"})`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Variable: "you", Label: "Person", Properties: map[string]interface{}{"name": "foo bar"}}},
		},
		{
			s:    `MATCH (you:Person {name:"You",age: 21})`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Variable: "you", Label: "Person", Properties: map[string]interface{}{"name": "You", "age": 21}}},
		},
		{
			s:    `MATCH (you:Person {name:"You",age: 21, happy :true})`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Variable: "you", Label: "Person", Properties: map[string]interface{}{"name": "You", "age": 21, "happy": true}}},
		},
		{
			s:    `MATCH (:Person)--(:Car)`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Label: "Person", Edge: &ir.EdgePatn{Vertex: &ir.VertexPatn{Label: "Car"}}}},
		},
		{
			s:    `MATCH (:Person)<--(:Car)`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Label: "Person", Edge: &ir.EdgePatn{Relationship: ir.Outbound, Vertex: &ir.VertexPatn{Label: "Car"}}}},
		},
		{
			s:    `MATCH (:Person)-->(:Car)`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Label: "Person", Edge: &ir.EdgePatn{Relationship: ir.Inbound, Vertex: &ir.VertexPatn{Label: "Car"}}}},
		},
		{
			s:    `MATCH (:Person)-[]-(:Car)`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Label: "Person", Edge: &ir.EdgePatn{Body: &ir.EdgeBodyStmt{LengthMinimum: 1, LengthMaximum: 1}, Vertex: &ir.VertexPatn{Label: "Car"}}}},
		},
		{
			s:    `MATCH (:Person)-[*2]-(:Car)`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Label: "Person", Edge: &ir.EdgePatn{Body: &ir.EdgeBodyStmt{LengthMinimum: 2, LengthMaximum: 2}, Vertex: &ir.VertexPatn{Label: "Car"}}}},
		},
		{
			s:    `MATCH (:Person)-[*..5]-(:Car)`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Label: "Person", Edge: &ir.EdgePatn{Body: &ir.EdgeBodyStmt{LengthMinimum: 1, LengthMaximum: 5}, Vertex: &ir.VertexPatn{Label: "Car"}}}},
		},
		{
			s:    `MATCH (:Person)-[*2..]-(:Car)`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Label: "Person", Edge: &ir.EdgePatn{Body: &ir.EdgeBodyStmt{LengthMinimum: 2, LengthMaximum: parser.MaxUint}, Vertex: &ir.VertexPatn{Label: "Car"}}}},
		},
		{
			s:    `MATCH (:Person)-[*2..5]-(:Car)`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Label: "Person", Edge: &ir.EdgePatn{Body: &ir.EdgeBodyStmt{LengthMinimum: 2, LengthMaximum: 5}, Vertex: &ir.VertexPatn{Label: "Car"}}}},
		},
		{
			s:    `MATCH (:Person)-[*]-(:Car)`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Label: "Person", Edge: &ir.EdgePatn{Body: &ir.EdgeBodyStmt{LengthMinimum: 1, LengthMaximum: parser.MaxUint}, Vertex: &ir.VertexPatn{Label: "Car"}}}},
		},
		{
			s:    `MATCH (:Person)-[:Owns*]-(:Car)`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Label: "Person", Edge: &ir.EdgePatn{Body: &ir.EdgeBodyStmt{Type: "Owns", LengthMinimum: 1, LengthMaximum: parser.MaxUint}, Vertex: &ir.VertexPatn{Label: "Car"}}}},
		},
		{
			s:    `MATCH (:Person)-[:Owns*2..5]-(:Car)`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Label: "Person", Edge: &ir.EdgePatn{Body: &ir.EdgeBodyStmt{Type: "Owns", LengthMinimum: 2, LengthMaximum: 5}, Vertex: &ir.VertexPatn{Label: "Car"}}}},
		},
		{
			s:    `MATCH (:Person)-[purchased:Owns*]-(:Car)`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Label: "Person", Edge: &ir.EdgePatn{Body: &ir.EdgeBodyStmt{Variable: "purchased", Type: "Owns", LengthMinimum: 1, LengthMaximum: parser.MaxUint}, Vertex: &ir.VertexPatn{Label: "Car"}}}},
		},
		{
			s:    `MATCH (:Person)-[* {blocked:false}]-(:Car)`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Label: "Person", Edge: &ir.EdgePatn{Body: &ir.EdgeBodyStmt{LengthMinimum: 1, LengthMaximum: parser.MaxUint, Properties: map[string]interface{}{"blocked": false}}, Vertex: &ir.VertexPatn{Label: "Car"}}}},
		},
		{
			s:    `MATCH (:Person)-[purchased:Owns*2..5 {blocked:false}]-(:Car)`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Label: "Person", Edge: &ir.EdgePatn{Body: &ir.EdgeBodyStmt{Variable: "purchased", Type: "Owns", LengthMinimum: 2, LengthMaximum: 5, Properties: map[string]interface{}{"blocked": false}}, Vertex: &ir.VertexPatn{Label: "Car"}}}},
		},
	}

	for i, tt := range tests {
		stmt, err := parser.NewParser().Parse(strings.NewReader(tt.s))
		if !reflect.DeepEqual(tt.err, errstring(err)) {
			t.Errorf("%d. %q: error mismatch:\n  exp=%s\n  got=%s\n\n", i, tt.s, tt.err, err)
		} else if tt.err == "" && !reflect.DeepEqual(tt.stmt, stmt) {
			t.Errorf("%d. %q\n\nstmt mismatch:\n\nexp=%#v\n\ngot=%#v\n\n", i, tt.s, tt.stmt, stmt)
		}
	}
}

// Ensure the parser can parse all the Clauses.
func TestParser_Clauses(t *testing.T) {
	var tests = []struct {
		s    string
		stmt ast.Stmt
		err  string
	}{
		{
			s:    `MATCH ()`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{}},
		},
		{
			s:    `OPTIONAL MATCH ()`,
			stmt: &ast.OptionalMatchStmt{Pattern: &ir.VertexPatn{}},
		},
		{
			s:    `CREATE ()`,
			stmt: &ast.CreateStmt{Pattern: &ir.VertexPatn{}},
		},
		{
			s:    `DELETE ()`,
			stmt: &ast.DeleteStmt{Pattern: &ir.VertexPatn{}},
		},
		// {
		// 	s:    `DETACH DELETE ()`,
		// 	stmt: &cypher.ClauseStmt{Pattern: &cypher.VertexPatn{}, Clause: cypher.DETACH_DELETE},
		// },
		// {
		// 	s:    `MERGE ()`,
		// 	stmt: &cypher.ClauseStmt{Pattern: &cypher.VertexPatn{}, Clause: cypher.MERGE},
		// },
		// {
		// 	s:    `REMOVE ()`,
		// 	stmt: &cypher.ClauseStmt{Pattern: &cypher.VertexPatn{}, Clause: cypher.REMOVE},
		// },
		// {
		// 	s:    `RETURN ()`,
		// 	stmt: &cypher.ClauseStmt{Pattern: &cypher.VertexPatn{}, Clause: cypher.RETURN},
		// },
		// {
		// 	s:    `SET ()`,
		// 	stmt: &cypher.ClauseStmt{Pattern: &cypher.VertexPatn{}, Clause: cypher.SET},
		// },
		// {
		// 	s:    `UNION ()`,
		// 	stmt: &cypher.ClauseStmt{Pattern: &cypher.VertexPatn{}, Clause: cypher.UNION},
		// },
		// {
		// 	s:    `UNWIND ()`,
		// 	stmt: &cypher.ClauseStmt{Pattern: &cypher.VertexPatn{}, Clause: cypher.UNWIND},
		// },
		// {
		// 	s:    `WITH ()`,
		// 	stmt: &cypher.ClauseStmt{Pattern: &cypher.VertexPatn{}, Clause: cypher.WITH},
		// },
	}

	for i, tt := range tests {
		stmt, err := parser.NewParser().Parse(strings.NewReader(tt.s))
		if !reflect.DeepEqual(tt.err, errstring(err)) {
			t.Errorf("%d. %q: error mismatch:\n  exp=%s\n  got=%s\n\n", i, tt.s, tt.err, err)
		} else if tt.err == "" && !reflect.DeepEqual(tt.stmt, stmt) {
			t.Errorf("%d. %q\n\nstmt mismatch:\n\nexp=%#v\n\ngot=%#v\n\n", i, tt.s, tt.stmt, stmt)
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

func TestParser_Where(t *testing.T) {

	var tests = []struct {
		s    string
		stmt ast.Stmt
		err  string
	}{
		{
			s:    `MATCH () WHERE n.number >= 1 AND n.number <= 10`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{}, Next: &ast.WhereStmt{Predicate: ast.NewBooleanExpr(expressions.AND, ast.NewComparisonExpr(expressions.GTE, &ast.PropertyStmt{Variable: "n", Value: "number"}, &ast.Ident{Data: 1}), ast.NewComparisonExpr(expressions.LTE, &ast.PropertyStmt{Variable: "n", Value: "number"}, &ast.Ident{Data: 10}))}},
		},
		{
			s:    `MATCH () WHERE n.name = "john smith"`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{}, Next: &ast.WhereStmt{Predicate: ast.NewComparisonExpr(expressions.EQ, &ast.PropertyStmt{Variable: "n", Value: "name"}, &ast.Ident{Data: "john smith"})}},
		},
	}

	for i, tt := range tests {
		stmt, err := parser.NewParser().Parse(strings.NewReader(tt.s))
		if !reflect.DeepEqual(tt.err, errstring(err)) {
			t.Errorf("%d. %q: error mismatch:\n  exp=%s\n  got=%s\n\n", i, tt.s, tt.err, err)
		} else if tt.err == "" && !reflect.DeepEqual(tt.stmt, stmt) {
			t.Errorf("%d. %q\n\nstmt mismatch:\n\nexp=%#v\n\ngot=%#v\n\n", i, tt.s, tt.stmt, stmt)
		}
	}
}

func TestParser_Return(t *testing.T) {

	var tests = []struct {
		s    string
		stmt ast.Stmt
		err  string
	}{
		{
			s:    `MATCH (n) RETURN n`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Variable: "n"}, Next: ast.NewReturnStmt(ast.NewMapProjectionStmt("n"))},
		},
		{
			s:    `MATCH (n) RETURN n { .number }`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Variable: "n"}, Next: ast.NewReturnStmt(ast.NewMapProjectionStmt("n", &ast.MapProperty{Key: "number"}))},
		},
		{
			s:    `MATCH (n) RETURN n { nrOfMovies }`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Variable: "n"}, Next: ast.NewReturnStmt(ast.NewMapProjectionStmt("n", &ast.MapVariable{Key: "nrOfMovies"}))},
		},
		{
			s:    `MATCH (n) RETURN n { .number, .name }`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Variable: "n"}, Next: ast.NewReturnStmt(ast.NewMapProjectionStmt("n", &ast.MapProperty{Key: "number"}, &ast.MapProperty{Key: "name"}))},
		},
		{
			s:    `MATCH (n) RETURN n { .number, nrOfMovies }`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Variable: "n"}, Next: ast.NewReturnStmt(ast.NewMapProjectionStmt("n", &ast.MapProperty{Key: "number"}, &ast.MapVariable{Key: "nrOfMovies"}))},
		},
	}

	for i, tt := range tests {
		stmt, err := parser.NewParser().Parse(strings.NewReader(tt.s))
		if !reflect.DeepEqual(tt.err, errstring(err)) {
			t.Errorf("%d. %q: error mismatch:\n  exp=%s\n  got=%s\n\n", i, tt.s, tt.err, err)
		} else if tt.err == "" && !reflect.DeepEqual(tt.stmt, stmt) {
			t.Errorf("%d. %q\n\nstmt mismatch:\n\nexp=%#v\n\ngot=%#v\n\n", i, tt.s, tt.stmt, stmt)
		}
	}
}
