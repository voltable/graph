package statements

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

type ReadStatement interface {
	Statement
}

type MatchStatement struct {
	Pattern *VertexStatement
	Next    Statement
}

type OptionalMatchStatement struct {
	Pattern *VertexStatement
	Next    Statement
}

type PrecedenceStatement struct {
	Predicate *PredicateStatement
}

// PredicateStatement
// n.name = 'Peter'
// n is Variable
// name is Property
// = is Operator
// Peter is Value
type PredicateStatement struct {
	Variable string
	Property string
	Operator Comparison
	Value    interface{}
	Next     BooleanStatement
}

type WhereStatement struct {
	Predicate *PredicateStatement
}
