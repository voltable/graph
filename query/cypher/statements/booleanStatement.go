package statements

type BooleanStatement interface {
}
type AndStatement struct {
	Precedence *PrecedenceStatement
	Predicate  *PredicateStatement
}

type OrStatement struct {
	Precedence *PrecedenceStatement
	Predicate  *PredicateStatement
}

type NotStatement struct {
	Precedence *PrecedenceStatement
	Predicate  *PredicateStatement
}

type XorStatement struct {
	Precedence *PrecedenceStatement
	Predicate  *PredicateStatement
}

type OrNotStatement struct {
	Precedence *PrecedenceStatement
	Predicate  *PredicateStatement
}
