package cypher

import "strings"

type Token int

const (
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
	clausesEnd

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
	NEQ  // <>
	LT   // <
	LTE  // <=
	GT   // >
	GTE  // >=
	IS   // IS
	NULL // NULL
	comparisonEnd
	// The operators STARTS WITH, ENDS WITH and CONTAINS can be used to search for a string value by its content.
	STARTSWITH // STARTS WITH
	ENDSWITH   // ENDS WITH
	CONTAINS   // CONTAINS

	// Regular expression matching
	EQREGEX // =~

	LPAREN    // (
	RPAREN    // )
	COMMA     // ,
	COLON     // :
	DOT       // .
	PIPE      // |
	LSQUARE   // [
	RSQUARE   // ]
	LCURLY    // {
	RCURLY    // }
	QUOTATION // "
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
	DELETE:   "DELETE",
	DETACH:   "DETACH",
	REMOVE:   "REMOVE",
	CALL:     "CALL",
	YIELD:    "YIELD",

	LPAREN:    "(",
	RPAREN:    ")",
	COMMA:     ",",
	COLON:     ":",
	DOT:       ".",
	PIPE:      "|",
	LSQUARE:   "[",
	RSQUARE:   "]",
	LCURLY:    "{",
	RCURLY:    "}",
	QUOTATION: "\"",

	AND: "AND",
	OR:  "OR",
	XOR: "XOR",
	NOT: "NOT",
}

var keywords map[string]Token
var comparison map[string]Token
var boolean map[string]Token

var eof = rune(0)

func init() {
	keywords = make(map[string]Token)
	for tok := clausesBag + 1; tok < clausesEnd; tok++ {
		keywords[strings.ToLower(tokens[tok])] = tok
	}

	comparison = make(map[string]Token)
	for tok := comparisonBeg + 1; tok < comparisonEnd; tok++ {
		comparison[strings.ToLower(tokens[tok])] = tok
	}

	boolean = make(map[string]Token)
	for tok := booleanBeg + 1; tok < booleanEnd; tok++ {
		boolean[strings.ToLower(tokens[tok])] = tok
	}
}

// String returns the string representation of the token.
func (tok Token) String() string {
	if tok >= 0 && tok < Token(len(tokens)) {
		return tokens[tok]
	}
	return ""
}

// isOperator returns true for operator tokens.
func (tok Token) isOperator() bool { return tok > operatorBeg && tok < operatorEnd }

func (tok Token) isComparison() bool { return tok > comparisonBeg && tok < comparisonEnd }

func Keyword(ident string) Token {
	if tok, ok := keywords[strings.ToLower(ident)]; ok {
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
