package parser_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
	"github.com/RossMerr/Caudex.Graph/query/cypher/parser"
)

// Ensure the parser can parse the right patterns.
func TestParser_Pattern(t *testing.T) {
	var tests = []struct {
		s    string
		stmt ast.Statement
		err  string
	}{
		{
			s:    `MATCH (you)`,
			stmt: &ast.MatchStmt{Pattern: &ast.VertexStmt{Variable: "you"}},
		},
		{
			s:    `MATCH (:Person)`,
			stmt: &ast.MatchStmt{Pattern: &ast.VertexStmt{Label: "Person"}},
		},
		{
			s:    `MATCH (you:Person)`,
			stmt: &ast.MatchStmt{Pattern: &ast.VertexStmt{Variable: "you", Label: "Person"}},
		},
		{
			s:    `MATCH (you:Person {name:"You"})`,
			stmt: &ast.MatchStmt{Pattern: &ast.VertexStmt{Variable: "you", Label: "Person", Properties: map[string]interface{}{"name": "You"}}},
		},
		{
			s:    `MATCH (you:Person {name:"You",age: 21})`,
			stmt: &ast.MatchStmt{Pattern: &ast.VertexStmt{Variable: "you", Label: "Person", Properties: map[string]interface{}{"name": "You", "age": 21}}},
		},
		{
			s:    `MATCH (you:Person {name:"You",age: 21, happy :true})`,
			stmt: &ast.MatchStmt{Pattern: &ast.VertexStmt{Variable: "you", Label: "Person", Properties: map[string]interface{}{"name": "You", "age": 21, "happy": true}}},
		},
		{
			s:    `MATCH (:Person)--(:Car)`,
			stmt: &ast.MatchStmt{Pattern: &ast.VertexStmt{Label: "Person", Edge: &ast.EdgeStmt{Vertex: &ast.VertexStmt{Label: "Car"}}}},
		},
		{
			s:    `MATCH (:Person)<--(:Car)`,
			stmt: &ast.MatchStmt{Pattern: &ast.VertexStmt{Label: "Person", Edge: &ast.EdgeStmt{Relationship: ast.Outbound, Vertex: &ast.VertexStmt{Label: "Car"}}}},
		},
		{
			s:    `MATCH (:Person)-->(:Car)`,
			stmt: &ast.MatchStmt{Pattern: &ast.VertexStmt{Label: "Person", Edge: &ast.EdgeStmt{Relationship: ast.Inbound, Vertex: &ast.VertexStmt{Label: "Car"}}}},
		},
		{
			s:    `MATCH (:Person)-[]-(:Car)`,
			stmt: &ast.MatchStmt{Pattern: &ast.VertexStmt{Label: "Person", Edge: &ast.EdgeStmt{Body: &ast.EdgeBodyStmt{LengthMinimum: 1, LengthMaximum: 1}, Vertex: &ast.VertexStmt{Label: "Car"}}}},
		},
		{
			s:    `MATCH (:Person)-[*2]-(:Car)`,
			stmt: &ast.MatchStmt{Pattern: &ast.VertexStmt{Label: "Person", Edge: &ast.EdgeStmt{Body: &ast.EdgeBodyStmt{LengthMinimum: 2, LengthMaximum: 2}, Vertex: &ast.VertexStmt{Label: "Car"}}}},
		},
		{
			s:    `MATCH (:Person)-[*..5]-(:Car)`,
			stmt: &ast.MatchStmt{Pattern: &ast.VertexStmt{Label: "Person", Edge: &ast.EdgeStmt{Body: &ast.EdgeBodyStmt{LengthMinimum: 1, LengthMaximum: 5}, Vertex: &ast.VertexStmt{Label: "Car"}}}},
		},
		{
			s:    `MATCH (:Person)-[*2..]-(:Car)`,
			stmt: &ast.MatchStmt{Pattern: &ast.VertexStmt{Label: "Person", Edge: &ast.EdgeStmt{Body: &ast.EdgeBodyStmt{LengthMinimum: 2, LengthMaximum: parser.MaxUint}, Vertex: &ast.VertexStmt{Label: "Car"}}}},
		},
		{
			s:    `MATCH (:Person)-[*2..5]-(:Car)`,
			stmt: &ast.MatchStmt{Pattern: &ast.VertexStmt{Label: "Person", Edge: &ast.EdgeStmt{Body: &ast.EdgeBodyStmt{LengthMinimum: 2, LengthMaximum: 5}, Vertex: &ast.VertexStmt{Label: "Car"}}}},
		},
		{
			s:    `MATCH (:Person)-[*]-(:Car)`,
			stmt: &ast.MatchStmt{Pattern: &ast.VertexStmt{Label: "Person", Edge: &ast.EdgeStmt{Body: &ast.EdgeBodyStmt{LengthMinimum: 1, LengthMaximum: parser.MaxUint}, Vertex: &ast.VertexStmt{Label: "Car"}}}},
		},
		{
			s:    `MATCH (:Person)-[:Owns*]-(:Car)`,
			stmt: &ast.MatchStmt{Pattern: &ast.VertexStmt{Label: "Person", Edge: &ast.EdgeStmt{Body: &ast.EdgeBodyStmt{Label: "Owns", LengthMinimum: 1, LengthMaximum: parser.MaxUint}, Vertex: &ast.VertexStmt{Label: "Car"}}}},
		},
		{
			s:    `MATCH (:Person)-[:Owns*2..5]-(:Car)`,
			stmt: &ast.MatchStmt{Pattern: &ast.VertexStmt{Label: "Person", Edge: &ast.EdgeStmt{Body: &ast.EdgeBodyStmt{Label: "Owns", LengthMinimum: 2, LengthMaximum: 5}, Vertex: &ast.VertexStmt{Label: "Car"}}}},
		},
		{
			s:    `MATCH (:Person)-[purchased:Owns*]-(:Car)`,
			stmt: &ast.MatchStmt{Pattern: &ast.VertexStmt{Label: "Person", Edge: &ast.EdgeStmt{Body: &ast.EdgeBodyStmt{Variable: "purchased", Label: "Owns", LengthMinimum: 1, LengthMaximum: parser.MaxUint}, Vertex: &ast.VertexStmt{Label: "Car"}}}},
		},
		{
			s:    `MATCH (:Person)-[* {blocked:false}]-(:Car)`,
			stmt: &ast.MatchStmt{Pattern: &ast.VertexStmt{Label: "Person", Edge: &ast.EdgeStmt{Body: &ast.EdgeBodyStmt{LengthMinimum: 1, LengthMaximum: parser.MaxUint, Properties: map[string]interface{}{"blocked": false}}, Vertex: &ast.VertexStmt{Label: "Car"}}}},
		},
		{
			s:    `MATCH (:Person)-[purchased:Owns*2..5 {blocked:false}]-(:Car)`,
			stmt: &ast.MatchStmt{Pattern: &ast.VertexStmt{Label: "Person", Edge: &ast.EdgeStmt{Body: &ast.EdgeBodyStmt{Variable: "purchased", Label: "Owns", LengthMinimum: 2, LengthMaximum: 5, Properties: map[string]interface{}{"blocked": false}}, Vertex: &ast.VertexStmt{Label: "Car"}}}},
		},
	}

	for i, tt := range tests {
		stmt, err := parser.NewParser(strings.NewReader(tt.s)).Parse()
		if !reflect.DeepEqual(tt.err, errstring(err)) {
			t.Errorf("%d. %q: error mismatch:\n  exp=%s\n  got=%s\n\n", i, tt.s, tt.err, err)
		} else if tt.err == "" && !reflect.DeepEqual(tt.stmt, stmt) {
			t.Errorf("%d. %q\n\nstmt mismatch:\n\nexp=%#v\n\ngot=%#v\n\n", i, tt.s, tt.stmt, stmt)
			//t.Errorf("%d. %q\n\nstmt mismatch:\n\nexp=%#v\n\ngot=%#v\n\n", i, tt.s, tt.stmt.Edge.Body, stmt.Edge.Body)
		}
	}
}

// Ensure the parser can parse all the Clauses.
func TestParser_Clauses(t *testing.T) {
	var tests = []struct {
		s    string
		stmt ast.Statement
		err  string
	}{
		{
			s:    `MATCH ()`,
			stmt: &ast.MatchStmt{Pattern: &ast.VertexStmt{}},
		},
		// {
		// 	s:    `OPTIONAL MATCH`,
		// 	stmt: &cypher.OptionalMatchStmt{},
		// },
		// {
		// 	s:    `CREATE ()`,
		// 	stmt: &cypher.CreateStmt{Pattern: &cypher.VertexStmt{}},
		// },
		// {
		// 	s:    `DELETE ()`,
		// 	stmt: &cypher.ClauseStmt{Pattern: &cypher.VertexStmt{}, Clause: cypher.DELETE},
		// },
		// {
		// 	s:    `DETACH DELETE ()`,
		// 	stmt: &cypher.ClauseStmt{Pattern: &cypher.VertexStmt{}, Clause: cypher.DETACH_DELETE},
		// },
		// {
		// 	s:    `MERGE ()`,
		// 	stmt: &cypher.ClauseStmt{Pattern: &cypher.VertexStmt{}, Clause: cypher.MERGE},
		// },
		// {
		// 	s:    `REMOVE ()`,
		// 	stmt: &cypher.ClauseStmt{Pattern: &cypher.VertexStmt{}, Clause: cypher.REMOVE},
		// },
		// {
		// 	s:    `RETURN ()`,
		// 	stmt: &cypher.ClauseStmt{Pattern: &cypher.VertexStmt{}, Clause: cypher.RETURN},
		// },
		// {
		// 	s:    `SET ()`,
		// 	stmt: &cypher.ClauseStmt{Pattern: &cypher.VertexStmt{}, Clause: cypher.SET},
		// },
		// {
		// 	s:    `UNION ()`,
		// 	stmt: &cypher.ClauseStmt{Pattern: &cypher.VertexStmt{}, Clause: cypher.UNION},
		// },
		// {
		// 	s:    `UNWIND ()`,
		// 	stmt: &cypher.ClauseStmt{Pattern: &cypher.VertexStmt{}, Clause: cypher.UNWIND},
		// },
		// {
		// 	s:    `WITH ()`,
		// 	stmt: &cypher.ClauseStmt{Pattern: &cypher.VertexStmt{}, Clause: cypher.WITH},
		// },
	}

	for i, tt := range tests {
		stmt, err := parser.NewParser(strings.NewReader(tt.s)).Parse()
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
		stmt ast.Statement
		err  string
	}{
		// {
		// 	s:    `MATCH () WHERE`,
		// 	stmt: &ast.MatchStmt{Pattern: &ast.VertexStmt{}, Next: &ast.WhereStmt{}},
		// },
		// {
		// 	s:    `MATCH () WHERE n.property <> 'value'`,
		// 	stmt: &ast.MatchStmt{Pattern: &ast.VertexStmt{}, Next: &ast.WhereStmt{Predicate: &ast.PredicateStmt{Variable: "n", Property: "property", Operator: ast.NEQ, Value: "value"}}},
		// },
		{
			s:    `MATCH () WHERE n.number >= 1 AND n.number <= 10`,
			stmt: &ast.MatchStmt{Pattern: &ast.VertexStmt{}, Next: &ast.WhereStmt{Predicate: &ast.PredicateStmt{Variable: "n", Property: "number", Operator: ast.GTE, Value: 1, Next: &ast.AndStmt{Predicate: &ast.PredicateStmt{Variable: "n", Property: "number", Operator: ast.LTE, Value: 10}}}}},
		},
		// {
		// 	s:    `MATCH () WHERE 1 <= n.number <= 10`,
		// 	stmt: &ast.MatchStmt{Pattern: &ast.VertexStmt{}, Next: &ast.WhereStmt{Predicate: &ast.PredicateStmt{Variable: "n", Property: "number", Operator: ast.GTE, Value: 1}}},
		// },
		// {
		// 	s:    `MATCH () WHERE n:Person`,
		// 	stmt: &ast.MatchStmt{Pattern: &ast.VertexStmt{}, Next: &ast.WhereStmt{Predicate: &ast.PredicateStmt{Variable: "n", Property: "property", Operator: ast.NEQ, Value: "value"}}},
		// },
		// {
		// 	s:    `MATCH () WHERE variable IS NULL`,
		// 	stmt: &ast.MatchStmt{Pattern: &ast.VertexStmt{}, Next: &ast.WhereStmt{Predicate: &ast.PredicateStmt{Variable: "n", Property: "property", Operator: ast.NEQ, Value: "value"}}},
		// },
		// {
		// 	s:    `MATCH () WHERE n["property"] = $value`,
		// 	stmt: &ast.MatchStmt{Pattern: &ast.VertexStmt{}, Next: &ast.WhereStmt{Predicate: &ast.PredicateStmt{Variable: "n", Property: "property", Operator: ast.NEQ, Value: "value"}}},
		// },
	}

	for i, tt := range tests {
		stmt, err := parser.NewParser(strings.NewReader(tt.s)).Parse()
		if !reflect.DeepEqual(tt.err, errstring(err)) {
			t.Errorf("%d. %q: error mismatch:\n  exp=%s\n  got=%s\n\n", i, tt.s, tt.err, err)
		} else if tt.err == "" && !reflect.DeepEqual(tt.stmt, stmt) {
			t.Errorf("%d. %q\n\nstmt mismatch:\n\nexp=%#v\n\ngot=%#v\n\n", i, tt.s, tt.stmt, stmt)
		}
	}
}