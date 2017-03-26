package cypher_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/RossMerr/Caudex.Graph/query/cypher"
	"github.com/RossMerr/Caudex.Graph/query/cypher/statements"
)

// Ensure the parser can parse the right patterns.
func TestParser_Pattern(t *testing.T) {
	var tests = []struct {
		s    string
		stmt statements.Statement
		err  string
	}{
		{
			s:    `MATCH (you)`,
			stmt: &statements.MatchStatement{Pattern: &statements.VertexStatement{Variable: "you"}},
		},
		{
			s:    `MATCH (:Person)`,
			stmt: &statements.MatchStatement{Pattern: &statements.VertexStatement{Label: "Person"}},
		},
		{
			s:    `MATCH (you:Person)`,
			stmt: &statements.MatchStatement{Pattern: &statements.VertexStatement{Variable: "you", Label: "Person"}},
		},
		{
			s:    `MATCH (you:Person {name:"You"})`,
			stmt: &statements.MatchStatement{Pattern: &statements.VertexStatement{Variable: "you", Label: "Person", Properties: map[string]interface{}{"name": "You"}}},
		},
		{
			s:    `MATCH (you:Person {name:"You",age: 21})`,
			stmt: &statements.MatchStatement{Pattern: &statements.VertexStatement{Variable: "you", Label: "Person", Properties: map[string]interface{}{"name": "You", "age": 21}}},
		},
		{
			s:    `MATCH (you:Person {name:"You",age: 21, happy :true})`,
			stmt: &statements.MatchStatement{Pattern: &statements.VertexStatement{Variable: "you", Label: "Person", Properties: map[string]interface{}{"name": "You", "age": 21, "happy": true}}},
		},
		{
			s:    `MATCH (:Person)--(:Car)`,
			stmt: &statements.MatchStatement{Pattern: &statements.VertexStatement{Label: "Person", Edge: &statements.EdgeStatement{Vertex: &statements.VertexStatement{Label: "Car"}}}},
		},
		{
			s:    `MATCH (:Person)<--(:Car)`,
			stmt: &statements.MatchStatement{Pattern: &statements.VertexStatement{Label: "Person", Edge: &statements.EdgeStatement{Relationship: statements.Outbound, Vertex: &statements.VertexStatement{Label: "Car"}}}},
		},
		{
			s:    `MATCH (:Person)-->(:Car)`,
			stmt: &statements.MatchStatement{Pattern: &statements.VertexStatement{Label: "Person", Edge: &statements.EdgeStatement{Relationship: statements.Inbound, Vertex: &statements.VertexStatement{Label: "Car"}}}},
		},
		{
			s:    `MATCH (:Person)-[]-(:Car)`,
			stmt: &statements.MatchStatement{Pattern: &statements.VertexStatement{Label: "Person", Edge: &statements.EdgeStatement{Body: &statements.EdgeBodyStatement{LengthMinimum: 1, LengthMaximum: 1}, Vertex: &statements.VertexStatement{Label: "Car"}}}},
		},
		{
			s:    `MATCH (:Person)-[*2]-(:Car)`,
			stmt: &statements.MatchStatement{Pattern: &statements.VertexStatement{Label: "Person", Edge: &statements.EdgeStatement{Body: &statements.EdgeBodyStatement{LengthMinimum: 2, LengthMaximum: 2}, Vertex: &statements.VertexStatement{Label: "Car"}}}},
		},
		{
			s:    `MATCH (:Person)-[*..5]-(:Car)`,
			stmt: &statements.MatchStatement{Pattern: &statements.VertexStatement{Label: "Person", Edge: &statements.EdgeStatement{Body: &statements.EdgeBodyStatement{LengthMinimum: 1, LengthMaximum: 5}, Vertex: &statements.VertexStatement{Label: "Car"}}}},
		},
		{
			s:    `MATCH (:Person)-[*2..]-(:Car)`,
			stmt: &statements.MatchStatement{Pattern: &statements.VertexStatement{Label: "Person", Edge: &statements.EdgeStatement{Body: &statements.EdgeBodyStatement{LengthMinimum: 2, LengthMaximum: cypher.MaxUint}, Vertex: &statements.VertexStatement{Label: "Car"}}}},
		},
		{
			s:    `MATCH (:Person)-[*2..5]-(:Car)`,
			stmt: &statements.MatchStatement{Pattern: &statements.VertexStatement{Label: "Person", Edge: &statements.EdgeStatement{Body: &statements.EdgeBodyStatement{LengthMinimum: 2, LengthMaximum: 5}, Vertex: &statements.VertexStatement{Label: "Car"}}}},
		},
		{
			s:    `MATCH (:Person)-[*]-(:Car)`,
			stmt: &statements.MatchStatement{Pattern: &statements.VertexStatement{Label: "Person", Edge: &statements.EdgeStatement{Body: &statements.EdgeBodyStatement{LengthMinimum: 1, LengthMaximum: cypher.MaxUint}, Vertex: &statements.VertexStatement{Label: "Car"}}}},
		},
		{
			s:    `MATCH (:Person)-[:Owns*]-(:Car)`,
			stmt: &statements.MatchStatement{Pattern: &statements.VertexStatement{Label: "Person", Edge: &statements.EdgeStatement{Body: &statements.EdgeBodyStatement{Label: "Owns", LengthMinimum: 1, LengthMaximum: cypher.MaxUint}, Vertex: &statements.VertexStatement{Label: "Car"}}}},
		},
		{
			s:    `MATCH (:Person)-[:Owns*2..5]-(:Car)`,
			stmt: &statements.MatchStatement{Pattern: &statements.VertexStatement{Label: "Person", Edge: &statements.EdgeStatement{Body: &statements.EdgeBodyStatement{Label: "Owns", LengthMinimum: 2, LengthMaximum: 5}, Vertex: &statements.VertexStatement{Label: "Car"}}}},
		},
		{
			s:    `MATCH (:Person)-[purchased:Owns*]-(:Car)`,
			stmt: &statements.MatchStatement{Pattern: &statements.VertexStatement{Label: "Person", Edge: &statements.EdgeStatement{Body: &statements.EdgeBodyStatement{Variable: "purchased", Label: "Owns", LengthMinimum: 1, LengthMaximum: cypher.MaxUint}, Vertex: &statements.VertexStatement{Label: "Car"}}}},
		},
		{
			s:    `MATCH (:Person)-[* {blocked:false}]-(:Car)`,
			stmt: &statements.MatchStatement{Pattern: &statements.VertexStatement{Label: "Person", Edge: &statements.EdgeStatement{Body: &statements.EdgeBodyStatement{LengthMinimum: 1, LengthMaximum: cypher.MaxUint, Properties: map[string]interface{}{"blocked": false}}, Vertex: &statements.VertexStatement{Label: "Car"}}}},
		},
		{
			s:    `MATCH (:Person)-[purchased:Owns*2..5 {blocked:false}]-(:Car)`,
			stmt: &statements.MatchStatement{Pattern: &statements.VertexStatement{Label: "Person", Edge: &statements.EdgeStatement{Body: &statements.EdgeBodyStatement{Variable: "purchased", Label: "Owns", LengthMinimum: 2, LengthMaximum: 5, Properties: map[string]interface{}{"blocked": false}}, Vertex: &statements.VertexStatement{Label: "Car"}}}},
		},
	}

	for i, tt := range tests {
		stmt, err := cypher.NewParser(strings.NewReader(tt.s)).Parse()
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
		stmt statements.Statement
		err  string
	}{
		{
			s:    `MATCH ()`,
			stmt: &statements.MatchStatement{Pattern: &statements.VertexStatement{}},
		},
		// {
		// 	s:    `OPTIONAL MATCH`,
		// 	stmt: &cypher.OptionalMatchStatement{},
		// },
		// {
		// 	s:    `CREATE ()`,
		// 	stmt: &cypher.CreateStatement{Pattern: &cypher.VertexStatement{}},
		// },
		// {
		// 	s:    `DELETE ()`,
		// 	stmt: &cypher.ClauseStatement{Pattern: &cypher.VertexStatement{}, Clause: cypher.DELETE},
		// },
		// {
		// 	s:    `DETACH DELETE ()`,
		// 	stmt: &cypher.ClauseStatement{Pattern: &cypher.VertexStatement{}, Clause: cypher.DETACH_DELETE},
		// },
		// {
		// 	s:    `MERGE ()`,
		// 	stmt: &cypher.ClauseStatement{Pattern: &cypher.VertexStatement{}, Clause: cypher.MERGE},
		// },
		// {
		// 	s:    `REMOVE ()`,
		// 	stmt: &cypher.ClauseStatement{Pattern: &cypher.VertexStatement{}, Clause: cypher.REMOVE},
		// },
		// {
		// 	s:    `RETURN ()`,
		// 	stmt: &cypher.ClauseStatement{Pattern: &cypher.VertexStatement{}, Clause: cypher.RETURN},
		// },
		// {
		// 	s:    `SET ()`,
		// 	stmt: &cypher.ClauseStatement{Pattern: &cypher.VertexStatement{}, Clause: cypher.SET},
		// },
		// {
		// 	s:    `UNION ()`,
		// 	stmt: &cypher.ClauseStatement{Pattern: &cypher.VertexStatement{}, Clause: cypher.UNION},
		// },
		// {
		// 	s:    `UNWIND ()`,
		// 	stmt: &cypher.ClauseStatement{Pattern: &cypher.VertexStatement{}, Clause: cypher.UNWIND},
		// },
		// {
		// 	s:    `WITH ()`,
		// 	stmt: &cypher.ClauseStatement{Pattern: &cypher.VertexStatement{}, Clause: cypher.WITH},
		// },
	}

	for i, tt := range tests {
		stmt, err := cypher.NewParser(strings.NewReader(tt.s)).Parse()
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
		stmt statements.Statement
		err  string
	}{
		// {
		// 	s:    `MATCH () WHERE`,
		// 	stmt: &statements.MatchStatement{Pattern: &statements.VertexStatement{}, Next: &statements.WhereStatement{}},
		// },
		// {
		// 	s:    `MATCH () WHERE n.property <> 'value'`,
		// 	stmt: &statements.MatchStatement{Pattern: &statements.VertexStatement{}, Next: &statements.WhereStatement{Predicate: &statements.PredicateStatement{Variable: "n", Property: "property", Operator: statements.NEQ, Value: "value"}}},
		// },

		{
			s:    `MATCH () WHERE n.number >= 1 AND n.number <= 10`,
			stmt: &statements.MatchStatement{Pattern: &statements.VertexStatement{}, Next: &statements.WhereStatement{Predicate: &statements.PredicateStatement{Variable: "n", Property: "number", Operator: statements.GTE, Value: 1, Next: &statements.AndStatement{Predicate: &statements.PredicateStatement{Variable: "n", Property: "number", Operator: statements.LTE, Value: 10}}}}},
		},
		// {
		// 	s:    `MATCH () WHERE 1 <= n.number <= 10`,
		// 	stmt: &statements.MatchStatement{Pattern: &statements.VertexStatement{}, Next: &statements.WhereStatement{Predicate: &statements.PredicateStatement{Variable: "n", Property: "property", Operator: statements.NEQ, Value: "value"}}},
		// },
		// {
		// 	s:    `MATCH () WHERE n:Person`,
		// 	stmt: &statements.MatchStatement{Pattern: &statements.VertexStatement{}, Next: &statements.WhereStatement{Predicate: &statements.PredicateStatement{Variable: "n", Property: "property", Operator: statements.NEQ, Value: "value"}}},
		// },
		// {
		// 	s:    `MATCH () WHERE variable IS NULL`,
		// 	stmt: &statements.MatchStatement{Pattern: &statements.VertexStatement{}, Next: &statements.WhereStatement{Predicate: &statements.PredicateStatement{Variable: "n", Property: "property", Operator: statements.NEQ, Value: "value"}}},
		// },
		// {
		// 	s:    `MATCH () WHERE n["property"] = $value`,
		// 	stmt: &statements.MatchStatement{Pattern: &statements.VertexStatement{}, Next: &statements.WhereStatement{Predicate: &statements.PredicateStatement{Variable: "n", Property: "property", Operator: statements.NEQ, Value: "value"}}},
		// },
	}

	for i, tt := range tests {
		stmt, err := cypher.NewParser(strings.NewReader(tt.s)).Parse()
		if !reflect.DeepEqual(tt.err, errstring(err)) {
			t.Errorf("%d. %q: error mismatch:\n  exp=%s\n  got=%s\n\n", i, tt.s, tt.err, err)
		} else if tt.err == "" && !reflect.DeepEqual(tt.stmt, stmt) {
			t.Errorf("%d. %q\n\nstmt mismatch:\n\nexp=%#v\n\ngot=%#v\n\n", i, tt.s, tt.stmt, stmt)
		}
	}
}
