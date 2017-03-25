package cypher_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/RossMerr/Caudex.Graph/query/cypher"
)

// Ensure the parser can parse the right patterns.
// func TestParser_Pattern(t *testing.T) {
// 	var tests = []struct {
// 		s    string
// 		stmt *cypher.ClauseStatement
// 		err  string
// 	}{
// 		{
// 			s:    `MATCH (you)`,
// 			stmt: &cypher.ClauseStatement{Pattern: &cypher.VertexStatement{Variable: "you"}, Clause: cypher.MATCH},
// 		},
// 		{
// 			s:    `MATCH (:Person)`,
// 			stmt: &cypher.ClauseStatement{Pattern: &cypher.VertexStatement{Label: "Person"}, Clause: cypher.MATCH},
// 		},
// 		{
// 			s:    `MATCH (you:Person)`,
// 			stmt: &cypher.ClauseStatement{Pattern: &cypher.VertexStatement{Variable: "you", Label: "Person"}, Clause: cypher.MATCH},
// 		},
// 		{
// 			s:    `MATCH (you:Person {name:"You"})`,
// 			stmt: &cypher.ClauseStatement{Pattern: &cypher.VertexStatement{Variable: "you", Label: "Person", Properties: map[string]interface{}{"name": "You"}}, Clause: cypher.MATCH},
// 		},
// 		{
// 			s:    `MATCH (you:Person {name:"You",age: 21})`,
// 			stmt: &cypher.ClauseStatement{Pattern: &cypher.VertexStatement{Variable: "you", Label: "Person", Properties: map[string]interface{}{"name": "You", "age": 21}}, Clause: cypher.MATCH},
// 		},
// 		{
// 			s:    `MATCH (you:Person {name:"You",age: 21, happy :true})`,
// 			stmt: &cypher.ClauseStatement{Pattern: &cypher.VertexStatement{Variable: "you", Label: "Person", Properties: map[string]interface{}{"name": "You", "age": 21, "happy": true}}, Clause: cypher.MATCH},
// 		},
// 		{
// 			s:    `MATCH (:Person)--(:Car)`,
// 			stmt: &cypher.ClauseStatement{Pattern: &cypher.VertexStatement{Label: "Person", Edge: &cypher.EdgeStatement{Vertex: &cypher.VertexStatement{Label: "Car"}}}, Clause: cypher.MATCH},
// 		},
// 		{
// 			s:    `MATCH (:Person)<--(:Car)`,
// 			stmt: &cypher.ClauseStatement{Pattern: &cypher.VertexStatement{Label: "Person", Edge: &cypher.EdgeStatement{Relationship: cypher.Outbound, Vertex: &cypher.VertexStatement{Label: "Car"}}}, Clause: cypher.MATCH},
// 		},
// 		{
// 			s:    `MATCH (:Person)-->(:Car)`,
// 			stmt: &cypher.ClauseStatement{Pattern: &cypher.VertexStatement{Label: "Person", Edge: &cypher.EdgeStatement{Relationship: cypher.Inbound, Vertex: &cypher.VertexStatement{Label: "Car"}}}, Clause: cypher.MATCH},
// 		},
// 		{
// 			s:    `MATCH (:Person)-[]-(:Car)`,
// 			stmt: &cypher.ClauseStatement{Pattern: &cypher.VertexStatement{Label: "Person", Edge: &cypher.EdgeStatement{Body: &cypher.EdgeBodyStatement{LengthMinimum: 1, LengthMaximum: 1}, Vertex: &cypher.VertexStatement{Label: "Car"}}}, Clause: cypher.MATCH},
// 		},
// 		{
// 			s:    `MATCH (:Person)-[*2]-(:Car)`,
// 			stmt: &cypher.ClauseStatement{Pattern: &cypher.VertexStatement{Label: "Person", Edge: &cypher.EdgeStatement{Body: &cypher.EdgeBodyStatement{LengthMinimum: 2, LengthMaximum: 2}, Vertex: &cypher.VertexStatement{Label: "Car"}}}, Clause: cypher.MATCH},
// 		},
// 		{
// 			s:    `MATCH (:Person)-[*..5]-(:Car)`,
// 			stmt: &cypher.ClauseStatement{Pattern: &cypher.VertexStatement{Label: "Person", Edge: &cypher.EdgeStatement{Body: &cypher.EdgeBodyStatement{LengthMinimum: 1, LengthMaximum: 5}, Vertex: &cypher.VertexStatement{Label: "Car"}}}, Clause: cypher.MATCH},
// 		},
// 		{
// 			s:    `MATCH (:Person)-[*2..]-(:Car)`,
// 			stmt: &cypher.ClauseStatement{Pattern: &cypher.VertexStatement{Label: "Person", Edge: &cypher.EdgeStatement{Body: &cypher.EdgeBodyStatement{LengthMinimum: 2, LengthMaximum: cypher.MaxUint}, Vertex: &cypher.VertexStatement{Label: "Car"}}}, Clause: cypher.MATCH},
// 		},
// 		{
// 			s:    `MATCH (:Person)-[*2..5]-(:Car)`,
// 			stmt: &cypher.ClauseStatement{Pattern: &cypher.VertexStatement{Label: "Person", Edge: &cypher.EdgeStatement{Body: &cypher.EdgeBodyStatement{LengthMinimum: 2, LengthMaximum: 5}, Vertex: &cypher.VertexStatement{Label: "Car"}}}, Clause: cypher.MATCH},
// 		},
// 		{
// 			s:    `MATCH (:Person)-[*]-(:Car)`,
// 			stmt: &cypher.ClauseStatement{Pattern: &cypher.VertexStatement{Label: "Person", Edge: &cypher.EdgeStatement{Body: &cypher.EdgeBodyStatement{LengthMinimum: 1, LengthMaximum: cypher.MaxUint}, Vertex: &cypher.VertexStatement{Label: "Car"}}}, Clause: cypher.MATCH},
// 		},
// 		{
// 			s:    `MATCH (:Person)-[:Owns*]-(:Car)`,
// 			stmt: &cypher.ClauseStatement{Pattern: &cypher.VertexStatement{Label: "Person", Edge: &cypher.EdgeStatement{Body: &cypher.EdgeBodyStatement{Label: "Owns", LengthMinimum: 1, LengthMaximum: cypher.MaxUint}, Vertex: &cypher.VertexStatement{Label: "Car"}}}, Clause: cypher.MATCH},
// 		},
// 		{
// 			s:    `MATCH (:Person)-[:Owns*2..5]-(:Car)`,
// 			stmt: &cypher.ClauseStatement{Pattern: &cypher.VertexStatement{Label: "Person", Edge: &cypher.EdgeStatement{Body: &cypher.EdgeBodyStatement{Label: "Owns", LengthMinimum: 2, LengthMaximum: 5}, Vertex: &cypher.VertexStatement{Label: "Car"}}}, Clause: cypher.MATCH},
// 		},
// 		{
// 			s:    `MATCH (:Person)-[purchased:Owns*]-(:Car)`,
// 			stmt: &cypher.ClauseStatement{Pattern: &cypher.VertexStatement{Label: "Person", Edge: &cypher.EdgeStatement{Body: &cypher.EdgeBodyStatement{Variable: "purchased", Label: "Owns", LengthMinimum: 1, LengthMaximum: cypher.MaxUint}, Vertex: &cypher.VertexStatement{Label: "Car"}}}, Clause: cypher.MATCH},
// 		},
// 		{
// 			s:    `MATCH (:Person)-[* {blocked:false}]-(:Car)`,
// 			stmt: &cypher.ClauseStatement{Pattern: &cypher.VertexStatement{Label: "Person", Edge: &cypher.EdgeStatement{Body: &cypher.EdgeBodyStatement{LengthMinimum: 1, LengthMaximum: cypher.MaxUint, Properties: map[string]interface{}{"blocked": false}}, Vertex: &cypher.VertexStatement{Label: "Car"}}}, Clause: cypher.MATCH},
// 		},
// 		{
// 			s:    `MATCH (:Person)-[purchased:Owns*2..5 {blocked:false}]-(:Car)`,
// 			stmt: &cypher.ClauseStatement{Pattern: &cypher.VertexStatement{Label: "Person", Edge: &cypher.EdgeStatement{Body: &cypher.EdgeBodyStatement{Variable: "purchased", Label: "Owns", LengthMinimum: 2, LengthMaximum: 5, Properties: map[string]interface{}{"blocked": false}}, Vertex: &cypher.VertexStatement{Label: "Car"}}}, Clause: cypher.MATCH},
// 		},
// 	}

// 	for i, tt := range tests {
// 		stmt, err := cypher.NewParser(strings.NewReader(tt.s)).Parse()
// 		if !reflect.DeepEqual(tt.err, errstring(err)) {
// 			t.Errorf("%d. %q: error mismatch:\n  exp=%s\n  got=%s\n\n", i, tt.s, tt.err, err)
// 		} else if tt.err == "" && !reflect.DeepEqual(tt.stmt, stmt) {
// 			t.Errorf("%d. %q\n\nstmt mismatch:\n\nexp=%#v\n\ngot=%#v\n\n", i, tt.s, tt.stmt, stmt)
// 			//t.Errorf("%d. %q\n\nstmt mismatch:\n\nexp=%#v\n\ngot=%#v\n\n", i, tt.s, tt.stmt.Edge.Body, stmt.Edge.Body)
// 		}
// 	}
// }

// Ensure the parser can parse all the Clauses.
func TestParser_Clauses(t *testing.T) {
	var tests = []struct {
		s    string
		stmt cypher.Statement
		err  string
	}{
		{
			s:    `MATCH ()`,
			stmt: &cypher.MatchStatement{Pattern: &cypher.VertexStatement{}},
		},

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
		// 	s:    `MATCH ()`,
		// 	stmt: &cypher.ClauseStatement{Pattern: &cypher.VertexStatement{}, Clause: cypher.MATCH},
		// },
		// {
		// 	s:    `MERGE ()`,
		// 	stmt: &cypher.ClauseStatement{Pattern: &cypher.VertexStatement{}, Clause: cypher.MERGE},
		// },
		// {
		// 	s:    `OPTIONAL MATCH ()`,
		// 	stmt: &cypher.ClauseStatement{Pattern: &cypher.VertexStatement{}, Clause: cypher.OPTIONAL_MATCH},
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
