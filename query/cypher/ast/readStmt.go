package ast

type Boolean int

const (
	// Boolean operators
	AND   Boolean = iota
	OR    Boolean = iota
	NOT   Boolean = iota
	XOR   Boolean = iota
	ORNOT Boolean = iota
)

type Comparison int

const (
	// EQ =
	EQ Comparison = iota
	// NEQ <>
	NEQ Comparison = iota
	// LT <
	LT Comparison = iota
	// LTE <=
	LTE Comparison = iota
	// GT >
	GT Comparison = iota
	// GTE >=
	GTE  Comparison = iota
	IS   Comparison = iota
	NULL Comparison = iota
)

type ReadStmt interface {
	Statement
}

type MatchStmt struct {
	Pattern *VertexStmt
	Next    Statement
}

type OptionalMatchStmt struct {
	Pattern *VertexStmt
	Next    Statement
}

type PrecedenceStmt struct {
	Predicate *PredicateStmt
}

// PredicateStatement
// n.name = 'Peter'
// n is Variable
// name is Property
// = is Operator
// Peter is Value
type PredicateStmt struct {
	Variable string
	Property string
	Operator Comparison
	Value    interface{}
	Next     BooleanStmt
}

type WhereStmt struct {
	Predicate *PredicateStmt
}
