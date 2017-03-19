package cypher_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/RossMerr/Caudex.Graph/query/cypher"
)

// Ensure the parser can parse strings into Statement ASTs.
func TestParser_ParseStatement(t *testing.T) {
	var tests = []struct {
		s    string
		stmt *cypher.VertexStatement
		err  string
	}{
		{
			s:    `MATCH ()`,
			stmt: &cypher.VertexStatement{},
		},
		{
			s:    `MATCH (you)`,
			stmt: &cypher.VertexStatement{Variable: "you"},
		},
		{
			s:    `MATCH (:Person)`,
			stmt: &cypher.VertexStatement{Label: "Person"},
		},
		{
			s:    `MATCH (you:Person)`,
			stmt: &cypher.VertexStatement{Variable: "you", Label: "Person"},
		},
		{
			s:    `MATCH (you:Person {name:"You"})`,
			stmt: &cypher.VertexStatement{Variable: "you", Label: "Person", Properties: map[string]interface{}{"name": "You"}},
		},
		{
			s:    `MATCH (you:Person {name:"You",age: 21})`,
			stmt: &cypher.VertexStatement{Variable: "you", Label: "Person", Properties: map[string]interface{}{"name": "You", "age": 21}},
		},
		{
			s:    `MATCH (you:Person {name:"You",age: 21, happy :true})`,
			stmt: &cypher.VertexStatement{Variable: "you", Label: "Person", Properties: map[string]interface{}{"name": "You", "age": 21, "happy": true}},
		},

		{
			s:    `MATCH (:Person)--(:Car)`,
			stmt: &cypher.VertexStatement{Label: "Person", Edge: &cypher.EdgeStatement{Vertex: &cypher.VertexStatement{Label: "Car"}}},
		},
		{
			s:    `MATCH (:Person)<--(:Car)`,
			stmt: &cypher.VertexStatement{Label: "Person", Edge: &cypher.EdgeStatement{Relationship: cypher.Outbound, Vertex: &cypher.VertexStatement{Label: "Car"}}},
		},
		{
			s:    `MATCH (:Person)-->(:Car)`,
			stmt: &cypher.VertexStatement{Label: "Person", Edge: &cypher.EdgeStatement{Relationship: cypher.Inbound, Vertex: &cypher.VertexStatement{Label: "Car"}}},
		},
		{
			s:    `MATCH (:Person)-[]-(:Car)`,
			stmt: &cypher.VertexStatement{Label: "Person", Edge: &cypher.EdgeStatement{Body: &cypher.EdgeBodyStatement{LengthMinimum: 1, LengthMaximum: 1}, Vertex: &cypher.VertexStatement{Label: "Car"}}},
		},
		{
			s:    `MATCH (:Person)-[*2]-(:Car)`,
			stmt: &cypher.VertexStatement{Label: "Person", Edge: &cypher.EdgeStatement{Body: &cypher.EdgeBodyStatement{LengthMinimum: 2, LengthMaximum: 2}, Vertex: &cypher.VertexStatement{Label: "Car"}}},
		},
		{
			s:    `MATCH (:Person)-[*..5]-(:Car)`,
			stmt: &cypher.VertexStatement{Label: "Person", Edge: &cypher.EdgeStatement{Body: &cypher.EdgeBodyStatement{LengthMinimum: 1, LengthMaximum: 5}, Vertex: &cypher.VertexStatement{Label: "Car"}}},
		},
		{
			s:    `MATCH (:Person)-[*2..]-(:Car)`,
			stmt: &cypher.VertexStatement{Label: "Person", Edge: &cypher.EdgeStatement{Body: &cypher.EdgeBodyStatement{LengthMinimum: 2, LengthMaximum: cypher.MaxUint}, Vertex: &cypher.VertexStatement{Label: "Car"}}},
		},
		{
			s:    `MATCH (:Person)-[*2..5]-(:Car)`,
			stmt: &cypher.VertexStatement{Label: "Person", Edge: &cypher.EdgeStatement{Body: &cypher.EdgeBodyStatement{LengthMinimum: 2, LengthMaximum: 5}, Vertex: &cypher.VertexStatement{Label: "Car"}}},
		},
		{
			s:    `MATCH (:Person)-[*]-(:Car)`,
			stmt: &cypher.VertexStatement{Label: "Person", Edge: &cypher.EdgeStatement{Body: &cypher.EdgeBodyStatement{LengthMinimum: 1, LengthMaximum: cypher.MaxUint}, Vertex: &cypher.VertexStatement{Label: "Car"}}},
		},
		{
			s:    `MATCH (:Person)-[:Owns*]-(:Car)`,
			stmt: &cypher.VertexStatement{Label: "Person", Edge: &cypher.EdgeStatement{Body: &cypher.EdgeBodyStatement{Label: "Owns", LengthMinimum: 1, LengthMaximum: cypher.MaxUint}, Vertex: &cypher.VertexStatement{Label: "Car"}}},
		},
		{
			s:    `MATCH (:Person)-[:Owns*2..5]-(:Car)`,
			stmt: &cypher.VertexStatement{Label: "Person", Edge: &cypher.EdgeStatement{Body: &cypher.EdgeBodyStatement{Label: "Owns", LengthMinimum: 2, LengthMaximum: 5}, Vertex: &cypher.VertexStatement{Label: "Car"}}},
		},
		{
			s:    `MATCH (:Person)-[purchased:Owns*]-(:Car)`,
			stmt: &cypher.VertexStatement{Label: "Person", Edge: &cypher.EdgeStatement{Body: &cypher.EdgeBodyStatement{Variable: "purchased", Label: "Owns", LengthMinimum: 1, LengthMaximum: cypher.MaxUint}, Vertex: &cypher.VertexStatement{Label: "Car"}}},
		},
		{
			s:    `MATCH (:Person)-[* {blocked:false}]-(:Car)`,
			stmt: &cypher.VertexStatement{Label: "Person", Edge: &cypher.EdgeStatement{Body: &cypher.EdgeBodyStatement{LengthMinimum: 1, LengthMaximum: cypher.MaxUint, Properties: map[string]interface{}{"blocked": false}}, Vertex: &cypher.VertexStatement{Label: "Car"}}},
		},
		{
			s:    `MATCH (:Person)-[purchased:Owns*2..5 {blocked:false}]-(:Car)`,
			stmt: &cypher.VertexStatement{Label: "Person", Edge: &cypher.EdgeStatement{Body: &cypher.EdgeBodyStatement{Variable: "purchased", Label: "Owns", LengthMinimum: 2, LengthMaximum: 5, Properties: map[string]interface{}{"blocked": false}}, Vertex: &cypher.VertexStatement{Label: "Car"}}},
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

// errstring returns the string representation of an error.
func errstring(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}
