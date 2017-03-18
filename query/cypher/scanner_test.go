package cypher_test

import (
	"strings"
	"testing"

	"github.com/RossMerr/Caudex.Graph/query/cypher"
)

// Ensure the scanner can scan tokens correctly.
func TestScanner_Scan(t *testing.T) {
	var tests = []struct {
		s   string
		tok cypher.Token
		lit string
	}{

		{s: ``, tok: cypher.EOF},
		{s: `MATCH`, tok: cypher.MATCH},
		{s: `RETURN`, tok: cypher.RETURN},
		{s: `UNWIND`, tok: cypher.UNWIND},
		{s: `OPTIONAL`, tok: cypher.OPTIONAL},
		{s: `WITH`, tok: cypher.WITH},
		{s: `UNION`, tok: cypher.UNION},
		{s: `CREATE`, tok: cypher.CREATE},
		{s: `MERGE`, tok: cypher.MERGE},
		{s: `SET`, tok: cypher.SET},
		{s: `DELETE`, tok: cypher.DELETE},
		{s: `DETACH`, tok: cypher.DETACH},
		{s: `REMOVE`, tok: cypher.REMOVE},
		{s: `CALL`, tok: cypher.CALL},
		{s: `YIELD`, tok: cypher.YIELD},
		{s: `(`, tok: cypher.LPAREN},
		{s: `)`, tok: cypher.RPAREN},
		{s: `,`, tok: cypher.COMMA},
		{s: `:`, tok: cypher.COLON},
		{s: `.`, tok: cypher.DOT},
		{s: `|`, tok: cypher.PIPE},
		{s: `[`, tok: cypher.LSQUARE},
		{s: `]`, tok: cypher.RSQUARE},
		{s: `{`, tok: cypher.LCURLY},
		{s: `}`, tok: cypher.RCURLY},
		{s: `"`, tok: cypher.QUOTATION},

		{s: `AND`, tok: cypher.AND},
		{s: `OR`, tok: cypher.OR},
		{s: `XOR`, tok: cypher.XOR},
		{s: `NOT`, tok: cypher.NOT},

		{s: `<`, tok: cypher.LT},
		{s: `>`, tok: cypher.GT},
		{s: `=`, tok: cypher.EQ},
		{s: `IS`, tok: cypher.IS},
		{s: `NULL`, tok: cypher.NULL},
	}

	for i, tt := range tests {
		s := cypher.NewScanner(strings.NewReader(tt.s))
		tok, lit := s.Scan()
		if tt.tok != tok {
			t.Errorf("%d. %q token mismatch: exp=%q got=%q <%q>", i, tt.s, tt.tok, tok, lit)
		} else if tt.s != lit {
			t.Errorf("%d. %q literal mismatch: exp=%q got=%q", i, tt.s, tt.lit, lit)
		}
	}
}
