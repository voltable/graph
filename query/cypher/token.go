package cypher

import "strings"

type Token int

const (
	ILLEGAL Token = iota
	EOF
	WS

	clausesBag
	MATCH
	RETURN
	UNWIND
	OPTIONALMATCH
	WITH
	UNION
	CREATE
	MERGE
	SET
	DELETE
	DETACHDELETE
	REMOVE
	CALLYIELD
	clausesEnd

	operatorBeg
	ADD // +
	SUB // -
	MUL // *
	DIV // /
	MOD // %
	POW // ^

	// The boolean operators are
	AND // AND
	OR  // OR
	XOR // XOR
	NOT // NOT

	// The comparison operators
	EQ        // =
	NEQ       // <>
	LT        // <
	LTE       // <=
	GT        // >
	GTE       // >=
	ISNULL    // IS NULL
	ISNOTNULL // IS NOT NULL

	// The operators STARTS WITH, ENDS WITH and CONTAINS can be used to search for a string value by its content.
	STARTSWITH // STARTS WITH
	ENDSWITH   // ENDS WI
	CONTAINS   // CONTAINS

	// Regular expression matching
	EQREGEX // =~
	operatorEnd

	LPAREN  // (
	RPAREN  // )
	COMMA   // ,
	COLON   // :
	DOT     // .
	PIPE    // |
	LSQUARE // [
	RSQUARE // ]
)

var keywords map[string]Token

var tokens = [...]string{
	MATCH:         "MATCH",
	RETURN:        "RETURN",
	UNWIND:        "UNWIND",
	OPTIONALMATCH: "OPTIONAL MATCH",
	WITH:          "WITH",
	UNION:         "UNION",
	CREATE:        "CREATE",
	MERGE:         "MERGE",
	SET:           "SET",
	DELETE:        "DELETE",
	DETACHDELETE:  "DETACH DELETE",
	REMOVE:        "REMOVE",
	CALLYIELD:     "CALL …​ YIELD",

	LPAREN:  "(",
	RPAREN:  ")",
	COMMA:   ",",
	COLON:   ":",
	DOT:     ".",
	PIPE:    "|",
	LSQUARE: "[",
	RSQUARE: "]",
}

var eof = rune(0)

func init() {
	keywords = make(map[string]Token)
	for tok := clausesBag + 1; tok < clausesEnd; tok++ {
		keywords[strings.ToLower(tokens[tok])] = tok
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
