package cypher_test

import (
	"reflect"
	"strings"
	"testing"

	"fmt"

	"github.com/RossMerr/Caudex.Graph/query/cypher"
)

// Ensure the parser can parse strings into Statement ASTs.
func TestParser_ParseStatement(t *testing.T) {
	var tests = []struct {
		s    string
		stmt *cypher.MatchVertexStatement
		err  string
	}{
		{
			s:    `MATCH (you:Person {name:"You"})`,
			stmt: &cypher.MatchVertexStatement{Variable: "you", Label: "Person", Properties: map[string]interface{}{"name": "You"}},
		},
		{
			s:    `MATCH (you:Person {name:"You",age: 21})`,
			stmt: &cypher.MatchVertexStatement{Variable: "you", Label: "Person", Properties: map[string]interface{}{"name": "You", "age": 21}},
		},
	}

	for i, tt := range tests {
		stmt, err := cypher.NewParser(strings.NewReader(tt.s)).Parse()
		if !reflect.DeepEqual(tt.err, errstring(err)) {
			t.Errorf("%d. %q: error mismatch:\n  exp=%s\n  got=%s\n\n", i, tt.s, tt.err, err)
		} else if tt.err == "" && !reflect.DeepEqual(tt.stmt, stmt) {
			fmt.Printf("%d", stmt)
			fmt.Printf("%d", tt.stmt)
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
