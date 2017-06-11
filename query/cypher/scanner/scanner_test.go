package scanner_test

import (
	"strings"
	"testing"

	"github.com/RossMerr/Caudex.Graph/query/cypher/lexer"
	"github.com/RossMerr/Caudex.Graph/query/cypher/scanner"
)

// Ensure the scanner can scan tokens correctly.
func TestScanner_Scan(t *testing.T) {
	var tests = []struct {
		s   string
		tok lexer.Token
		lit string
	}{

		// {s: ``, tok: token.EOF},
		// {s: `MATCH`, tok: token.MATCH},
		// {s: `RETURN`, tok: token.RETURN},
		// {s: `UNWIND`, tok: token.UNWIND},
		// {s: `OPTIONAL`, tok: token.OPTIONAL},
		// {s: `WITH`, tok: token.WITH},
		// {s: `UNION`, tok: token.UNION},
		// {s: `CREATE`, tok: token.CREATE},
		// {s: `MERGE`, tok: token.MERGE},
		// {s: `SET`, tok: token.SET},
		// {s: `DELETE`, tok: token.DELETE},
		// {s: `DETACH`, tok: token.DETACH},
		// {s: `REMOVE`, tok: token.REMOVE},
		// {s: `CALL`, tok: token.CALL},
		// {s: `YIELD`, tok: token.YIELD},
		// {s: `(`, tok: token.LPAREN},
		// {s: `)`, tok: token.RPAREN},
		// {s: `,`, tok: token.COMMA},
		// {s: `:`, tok: token.COLON},
		// {s: `.`, tok: token.DOT},
		// {s: `|`, tok: token.PIPE},
		// {s: `[`, tok: token.LSQUARE},
		// {s: `]`, tok: token.RSQUARE},
		// {s: `{`, tok: token.LCURLY},
		// {s: `}`, tok: token.RCURLY},
		// {s: `"`, tok: token.QUOTATION},

		// {s: `AND`, tok: token.AND},
		// {s: `OR`, tok: token.OR},
		// {s: `XOR`, tok: token.XOR},
		// {s: `NOT`, tok: token.NOT},

		// {s: `<`, tok: token.LT},
		// {s: `>`, tok: token.GT},
		{s: `<>`, tok: lexer.NEQ},
		// {s: `=`, tok: token.EQ},
		// {s: `IS`, tok: token.IS},
		// {s: `NULL`, tok: token.NULL},
	}

	for i, tt := range tests {
		s := scanner.NewScanner(strings.NewReader(tt.s))
		tok, lit := s.Scan()
		if tt.tok != tok {
			t.Errorf("%d. %q token mismatch: exp=%q got=%q <%q>", i, tt.s, tt.tok, tok, lit)
		} else if tt.s != lit {
			t.Errorf("%d. %q literal mismatch: exp=%q got=%q", i, tt.s, tt.lit, lit)
		}
	}
}
