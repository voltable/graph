package ast

type BooleanStmt interface {
}
type AndStmt struct {
	Precedence *PrecedenceStmt
	Predicate  *PredicateStmt
}

type OrStmt struct {
	Precedence *PrecedenceStmt
	Predicate  *PredicateStmt
}

type NotStmt struct {
	Precedence *PrecedenceStmt
	Predicate  *PredicateStmt
}

type XorStmt struct {
	Precedence *PrecedenceStmt
	Predicate  *PredicateStmt
}

type OrNotStmt struct {
	Precedence *PrecedenceStmt
	Predicate  *PredicateStmt
}
