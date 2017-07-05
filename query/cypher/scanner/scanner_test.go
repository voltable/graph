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

		{s: ``, tok: lexer.EOF},
		{s: `MATCH`, tok: lexer.MATCH},
		{s: `RETURN`, tok: lexer.RETURN},
		{s: `UNWIND`, tok: lexer.UNWIND},
		{s: `OPTIONAL`, tok: lexer.OPTIONAL},
		{s: `WITH`, tok: lexer.WITH},
		{s: `UNION`, tok: lexer.UNION},
		{s: `CREATE`, tok: lexer.CREATE},
		{s: `MERGE`, tok: lexer.MERGE},
		{s: `SET`, tok: lexer.SET},
		{s: `DELETE`, tok: lexer.DELETE},
		{s: `DETACH`, tok: lexer.DETACH},
		{s: `REMOVE`, tok: lexer.REMOVE},
		{s: `CALL`, tok: lexer.CALL},
		{s: `YIELD`, tok: lexer.YIELD},
		{s: `(`, tok: lexer.LPAREN},
		{s: `)`, tok: lexer.RPAREN},
		{s: `,`, tok: lexer.COMMA},
		{s: `:`, tok: lexer.COLON},
		{s: `.`, tok: lexer.DOT},
		{s: `|`, tok: lexer.PIPE},
		{s: `[`, tok: lexer.LSQUARE},
		{s: `]`, tok: lexer.RSQUARE},
		{s: `{`, tok: lexer.LCURLY},
		{s: `}`, tok: lexer.RCURLY},
		{s: `"`, tok: lexer.QUOTATION},

		{s: `AND`, tok: lexer.AND},
		{s: `OR`, tok: lexer.OR},
		{s: `XOR`, tok: lexer.XOR},
		{s: `NOT`, tok: lexer.NOT},

		{s: `<`, tok: lexer.LT},
		{s: `>`, tok: lexer.GT},
		{s: `<>`, tok: lexer.NEQ},
		{s: `=`, tok: lexer.EQ},
		{s: `IS`, tok: lexer.IS},
		{s: `NULL`, tok: lexer.NULL},
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
