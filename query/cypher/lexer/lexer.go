package lexer

import (
	"strings"
)

// Token is the set of lexical tokens
type Token int

const (
	// Special tokens
	ILLEGAL Token = iota
	EOF
	WS

	// Literals
	IDENT // fields, table_name

	clausesBag
	MATCH
	RETURN
	UNWIND
	OPTIONAL
	WITH
	UNION
	CREATE
	MERGE
	SET
	DELETE
	DETACH
	REMOVE
	CALL
	YIELD
	AS
	clausesEnd

	DETACH_DELETE
	OPTIONAL_MATCH

	subClausesBag
	LIMIT
	ON
	ORDER
	SKIP
	WHERE
	subClausesEnd

	BY
	ON_CREATE
	ON_MATCH
	ORDER_BY

	operatorBeg
	ADD // +
	SUB // -
	MUL // *
	DIV // /
	MOD // %
	POW // ^
	operatorEnd
	// The boolean operators are
	booleanBeg
	AND // AND
	OR  // OR
	XOR // XOR
	NOT // NOT
	booleanEnd
	// The comparison operators
	comparisonBeg
	EQ   // =
	LT   // <
	GT   // >
	IS   // IS
	NULL // NULL
	comparisonEnd
	// The operators STARTS WITH, ENDS WITH and CONTAINS can be used to search for a string value by its content.
	STARTSWITH // STARTS WITH
	ENDSWITH   // ENDS WITH
	CONTAINS   // CONTAINS

	// Other characters that are long
	longCharacterBeg
	NEQ     // <>
	LTE     // <=
	GTE     // >=
	EQREGEX // =~
	longCharacterEnd

	LPAREN  // (
	RPAREN  // )
	COMMA   // ,
	COLON   // :
	DOT     // .
	PIPE    // |
	LSQUARE // [
	RSQUARE // ]
	LCURLY  // {
	RCURLY  // }
	TILDE   // ~

	quotationBeg
	QUOTATION       // "
	SINGLEQUOTATION // '
	GRAVE           // `
	quotationEnd
)

var tokens = [...]string{
	MATCH:    "MATCH",
	RETURN:   "RETURN",
	UNWIND:   "UNWIND",
	OPTIONAL: "OPTIONAL",
	WITH:     "WITH",
	UNION:    "UNION",
	CREATE:   "CREATE",
	MERGE:    "MERGE",
	SET:      "SET",

	WHERE: "WHERE",
	LIMIT: "LIMIT",
	ON:    "ON",
	ORDER: "ORDER",
	SKIP:  "SKIP",

	//Delete a node and a relationship.
	DELETE: "DELETE",
	DETACH: "DETACH",
	REMOVE: "REMOVE",
	CALL:   "CALL",
	YIELD:  "YIELD",

	AS: "AS",

	LPAREN:  "(",
	RPAREN:  ")",
	COMMA:   ",",
	COLON:   ":",
	DOT:     ".",
	PIPE:    "|",
	LSQUARE: "[",
	RSQUARE: "]",
	LCURLY:  "{",
	RCURLY:  "}",
	TILDE:   "~",
	EQREGEX: "=~",

	QUOTATION:       "\"",
	SINGLEQUOTATION: "'",
	GRAVE:           "`",

	NEQ: "<>",
	LTE: "<=",
	GTE: ">=",

	AND: "AND",
	OR:  "OR",
	XOR: "XOR",
	NOT: "NOT",
}

var clauses map[string]Token
var subClauses map[string]Token
var comparison map[string]Token
var boolean map[string]Token
var quotations map[string]Token

func init() {
	clauses = make(map[string]Token)
	for tok := clausesBag + 1; tok < clausesEnd; tok++ {
		clauses[strings.ToLower(tokens[tok])] = tok
	}

	subClauses = make(map[string]Token)
	for tok := subClausesBag + 1; tok < subClausesEnd; tok++ {
		subClauses[strings.ToLower(tokens[tok])] = tok
	}

	comparison = make(map[string]Token)
	for tok := comparisonBeg + 1; tok < comparisonEnd; tok++ {
		comparison[strings.ToLower(tokens[tok])] = tok
	}

	boolean = make(map[string]Token)
	for tok := booleanBeg + 1; tok < booleanEnd; tok++ {
		boolean[strings.ToLower(tokens[tok])] = tok
	}

	quotations = make(map[string]Token)
	for tok := quotationBeg + 1; tok < quotationEnd; tok++ {
		quotations[strings.ToLower(tokens[tok])] = tok
	}
}

// String returns the string representation of the token.
func (tok Token) String() string {
	if tok >= 0 && tok < Token(len(tokens)) {
		return tokens[tok]
	}
	return ""
}

// IsClause returns true for clauses tokens.
func (tok Token) IsClause() bool { return tok > clausesBag && tok < clausesEnd }

// IsSubClause returns true for clauses tokens.
func (tok Token) IsSubClause() bool { return tok > subClausesBag && tok < subClausesEnd }

// IsOperator returns true for operator tokens.
func (tok Token) IsOperator() bool { return tok > operatorBeg && tok < operatorEnd }

// IsComparison returns true for comparison tokens.
func (tok Token) IsComparison() bool { return tok > comparisonBeg && tok < comparisonEnd }

// IsQuotation returns true for comparison tokens.
func (tok Token) IsQuotation() bool { return tok > quotationBeg && tok < quotationEnd }

func Clause(ident string) Token {
	if tok, ok := clauses[strings.ToLower(ident)]; ok {
		return tok
	}
	return IDENT
}

func SubClause(ident string) Token {
	if tok, ok := subClauses[strings.ToLower(ident)]; ok {
		return tok
	}
	return IDENT
}

func Boolean(ident string) Token {
	if tok, ok := boolean[strings.ToLower(ident)]; ok {
		return tok
	}
	return IDENT
}

func Comparison(ident string) Token {
	if tok, ok := comparison[strings.ToLower(ident)]; ok {
		return tok
	}
	return IDENT
}

func Quotation(ident string) Token {
	if tok, ok := quotations[strings.ToLower(ident)]; ok {
		return tok
	}
	return IDENT
}
