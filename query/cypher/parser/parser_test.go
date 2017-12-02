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
			s:    `MATCH (you) RETURN *`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Variable: "you"}, Next: ast.NewReturnStmt(ast.NewMapProjectionStmt("*", &ast.MapAll{}))},
		},
		{
			s:    `MATCH (:Person) RETURN *`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Label: "Person"}, Next: ast.NewReturnStmt(ast.NewMapProjectionStmt("*", &ast.MapAll{}))},
		},
		{
			s:    `MATCH (you:Person) RETURN *`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Variable: "you", Label: "Person"}, Next: ast.NewReturnStmt(ast.NewMapProjectionStmt("*", &ast.MapAll{}))},
		},
		{
			s:    `MATCH (you:Person {name:"You"}) RETURN *`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Variable: "you", Label: "Person", Properties: map[string]interface{}{"name": "You"}}, Next: ast.NewReturnStmt(ast.NewMapProjectionStmt("*", &ast.MapAll{}))},
		},
		{
			s:    `MATCH (you:Person {name:"foo bar"}) RETURN *`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Variable: "you", Label: "Person", Properties: map[string]interface{}{"name": "foo bar"}}, Next: ast.NewReturnStmt(ast.NewMapProjectionStmt("*", &ast.MapAll{}))},
		},
		{
			s:    `MATCH (you:Person {name:"You",age: 21}) RETURN *`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Variable: "you", Label: "Person", Properties: map[string]interface{}{"name": "You", "age": 21}}, Next: ast.NewReturnStmt(ast.NewMapProjectionStmt("*", &ast.MapAll{}))},
		},
		{
			s:    `MATCH (you:Person {name:"You",age: 21, happy :true}) RETURN *`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Variable: "you", Label: "Person", Properties: map[string]interface{}{"name": "You", "age": 21, "happy": true}}, Next: ast.NewReturnStmt(ast.NewMapProjectionStmt("*", &ast.MapAll{}))},
		},
		{
			s:    `MATCH (:Person)--(:Car) RETURN *`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Label: "Person", Edge: &ir.EdgePatn{Vertex: &ir.VertexPatn{Label: "Car"}}}, Next: ast.NewReturnStmt(ast.NewMapProjectionStmt("*", &ast.MapAll{}))},
		},
		{
			s:    `MATCH (:Person)<--(:Car) RETURN *`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Label: "Person", Edge: &ir.EdgePatn{Relationship: ir.Outbound, Vertex: &ir.VertexPatn{Label: "Car"}}}, Next: ast.NewReturnStmt(ast.NewMapProjectionStmt("*", &ast.MapAll{}))},
		},
		{
			s:    `MATCH (:Person)-->(:Car) RETURN *`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Label: "Person", Edge: &ir.EdgePatn{Relationship: ir.Inbound, Vertex: &ir.VertexPatn{Label: "Car"}}}, Next: ast.NewReturnStmt(ast.NewMapProjectionStmt("*", &ast.MapAll{}))},
		},
		{
			s:    `MATCH (:Person)-[]-(:Car) RETURN *`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Label: "Person", Edge: &ir.EdgePatn{Body: &ir.EdgeBodyStmt{LengthMinimum: 1, LengthMaximum: 1}, Vertex: &ir.VertexPatn{Label: "Car"}}}, Next: ast.NewReturnStmt(ast.NewMapProjectionStmt("*", &ast.MapAll{}))},
		},
		{
			s:    `MATCH (:Person)-[*2]-(:Car) RETURN *`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Label: "Person", Edge: &ir.EdgePatn{Body: &ir.EdgeBodyStmt{LengthMinimum: 2, LengthMaximum: 2}, Vertex: &ir.VertexPatn{Label: "Car"}}}, Next: ast.NewReturnStmt(ast.NewMapProjectionStmt("*", &ast.MapAll{}))},
		},
		{
			s:    `MATCH (:Person)-[*..5]-(:Car) RETURN *`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Label: "Person", Edge: &ir.EdgePatn{Body: &ir.EdgeBodyStmt{LengthMinimum: 1, LengthMaximum: 5}, Vertex: &ir.VertexPatn{Label: "Car"}}}, Next: ast.NewReturnStmt(ast.NewMapProjectionStmt("*", &ast.MapAll{}))},
		},
		{
			s:    `MATCH (:Person)-[*2..]-(:Car) RETURN *`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Label: "Person", Edge: &ir.EdgePatn{Body: &ir.EdgeBodyStmt{LengthMinimum: 2, LengthMaximum: parser.MaxUint}, Vertex: &ir.VertexPatn{Label: "Car"}}}, Next: ast.NewReturnStmt(ast.NewMapProjectionStmt("*", &ast.MapAll{}))},
		},
		{
			s:    `MATCH (:Person)-[*2..5]-(:Car) RETURN *`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Label: "Person", Edge: &ir.EdgePatn{Body: &ir.EdgeBodyStmt{LengthMinimum: 2, LengthMaximum: 5}, Vertex: &ir.VertexPatn{Label: "Car"}}}, Next: ast.NewReturnStmt(ast.NewMapProjectionStmt("*", &ast.MapAll{}))},
		},
		{
			s:    `MATCH (:Person)-[*]-(:Car) RETURN *`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Label: "Person", Edge: &ir.EdgePatn{Body: &ir.EdgeBodyStmt{LengthMinimum: 1, LengthMaximum: parser.MaxUint}, Vertex: &ir.VertexPatn{Label: "Car"}}}, Next: ast.NewReturnStmt(ast.NewMapProjectionStmt("*", &ast.MapAll{}))},
		},
		{
			s:    `MATCH (:Person)-[:Owns*]-(:Car) RETURN *`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Label: "Person", Edge: &ir.EdgePatn{Body: &ir.EdgeBodyStmt{Type: "Owns", LengthMinimum: 1, LengthMaximum: parser.MaxUint}, Vertex: &ir.VertexPatn{Label: "Car"}}}, Next: ast.NewReturnStmt(ast.NewMapProjectionStmt("*", &ast.MapAll{}))},
		},
		{
			s:    `MATCH (:Person)-[:Owns*2..5]-(:Car) RETURN *`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Label: "Person", Edge: &ir.EdgePatn{Body: &ir.EdgeBodyStmt{Type: "Owns", LengthMinimum: 2, LengthMaximum: 5}, Vertex: &ir.VertexPatn{Label: "Car"}}}, Next: ast.NewReturnStmt(ast.NewMapProjectionStmt("*", &ast.MapAll{}))},
		},
		{
			s:    `MATCH (:Person)-[purchased:Owns*]-(:Car) RETURN *`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Label: "Person", Edge: &ir.EdgePatn{Body: &ir.EdgeBodyStmt{Variable: "purchased", Type: "Owns", LengthMinimum: 1, LengthMaximum: parser.MaxUint}, Vertex: &ir.VertexPatn{Label: "Car"}}}, Next: ast.NewReturnStmt(ast.NewMapProjectionStmt("*", &ast.MapAll{}))},
		},
		{
			s:    `MATCH (:Person)-[* {blocked:false}]-(:Car) RETURN *`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Label: "Person", Edge: &ir.EdgePatn{Body: &ir.EdgeBodyStmt{LengthMinimum: 1, LengthMaximum: parser.MaxUint, Properties: map[string]interface{}{"blocked": false}}, Vertex: &ir.VertexPatn{Label: "Car"}}}, Next: ast.NewReturnStmt(ast.NewMapProjectionStmt("*", &ast.MapAll{}))},
		},
		{
			s:    `MATCH (:Person)-[purchased:Owns*2..5 {blocked:false}]-(:Car) RETURN *`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Label: "Person", Edge: &ir.EdgePatn{Body: &ir.EdgeBodyStmt{Variable: "purchased", Type: "Owns", LengthMinimum: 2, LengthMaximum: 5, Properties: map[string]interface{}{"blocked": false}}, Vertex: &ir.VertexPatn{Label: "Car"}}}, Next: ast.NewReturnStmt(ast.NewMapProjectionStmt("*", &ast.MapAll{}))},
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
			s:    `MATCH () RETURN *`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{}, Next: ast.NewReturnStmt(ast.NewMapProjectionStmt("*", &ast.MapAll{}))},
		},
		// {
		// 	s:    `OPTIONAL MATCH () RETURN *`,
		// 	stmt: &ast.OptionalMatchStmt{Pattern: &ir.VertexPatn{}, Next: ast.NewReturnStmt(ast.NewMapProjectionStmt("*", &ast.MapAll{}))},
		// },
		// {
		// 	s:    `CREATE () RETURN *`,
		// 	stmt: &ast.CreateStmt{Pattern: &ir.VertexPatn{}, Next: ast.NewReturnStmt(ast.NewMapProjectionStmt("*", &ast.MapAll{}))},
		// },
		// {
		// 	s:    `DELETE () RETURN *`,
		// 	stmt: &ast.DeleteStmt{Pattern: &ir.VertexPatn{}, Next: ast.NewReturnStmt(ast.NewMapProjectionStmt("*", &ast.MapAll{}))},
		// },
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
			s:    `MATCH (n) RETURN *`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Variable: "n"}, Next: ast.NewReturnStmt(ast.NewMapProjectionStmt("*", &ast.MapAll{}))},
		},
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
			s:    `MATCH (n) RETURN n { .* }`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Variable: "n"}, Next: ast.NewReturnStmt(ast.NewMapProjectionStmt("n", &ast.MapAll{}))},
		},
		{
			s:    `MATCH (n) RETURN n { .number, .name }`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Variable: "n"}, Next: ast.NewReturnStmt(ast.NewMapProjectionStmt("n", &ast.MapProperty{Key: "number"}, &ast.MapProperty{Key: "name"}))},
		},
		{
			s:    `MATCH (n) RETURN n { .number, nrOfMovies }`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Variable: "n"}, Next: ast.NewReturnStmt(ast.NewMapProjectionStmt("n", &ast.MapProperty{Key: "number"}, &ast.MapVariable{Key: "nrOfMovies"}))},
		},
		{
			s:    `MATCH (n) RETURN n { .number, nrOfMovies, .* }`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Variable: "n"}, Next: ast.NewReturnStmt(ast.NewMapProjectionStmt("n", &ast.MapProperty{Key: "number"}, &ast.MapVariable{Key: "nrOfMovies"}, &ast.MapAll{}))},
		},
		{
			s:    `MATCH (n) RETURN n.number`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Variable: "n"}, Next: ast.NewReturnStmt(ast.NewMapProjectionStmt("n", &ast.MapProperty{Key: "number"}))},
		},
		{
			s:    `MATCH (n) RETURN n.number, n.name`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Variable: "n"}, Next: ast.NewReturnStmt(ast.NewMapProjectionStmt("n", &ast.MapProperty{Key: "number"}, &ast.MapProperty{Key: "name"}))},
		},
		{
			s:    "MATCH (n) RETURN `This isn't a common variable`",
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Variable: "n"}, Next: ast.NewReturnStmt(ast.NewMapProjectionStmt(`This isn't a common variable`))},
		},
		{
			s:    "MATCH (n) RETURN `This isn't a common variable`.number",
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Variable: "n"}, Next: ast.NewReturnStmt(ast.NewMapProjectionStmt(`This isn't a common variable`, &ast.MapProperty{Key: "number"}))},
		},
		{
			s:    "MATCH (n) RETURN `This isn't a common variable` { .number }",
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Variable: "n"}, Next: ast.NewReturnStmt(ast.NewMapProjectionStmt(`This isn't a common variable`, &ast.MapProperty{Key: "number"}))},
		},
		{
			s:    `MATCH (n) RETURN n.number AS SomethingTotallyDifferent`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Variable: "n"}, Next: ast.NewReturnStmt(ast.NewMapProjectionStmt("n", &ast.MapProperty{Key: "number", Alias: "SomethingTotallyDifferent"}))},
		},
		{
			s:    `MATCH (n) RETURN n { .number AS SomethingTotallyDifferent}`,
			stmt: &ast.MatchStmt{Pattern: &ir.VertexPatn{Variable: "n"}, Next: ast.NewReturnStmt(ast.NewMapProjectionStmt("n", &ast.MapProperty{Key: "number", Alias: "SomethingTotallyDifferent"}))},
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
