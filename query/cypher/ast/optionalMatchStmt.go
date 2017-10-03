package ast

// OptionalMatchStmt used to search for the pattern described in it, can match on nil
type OptionalMatchStmt struct {
	Pattern Patn
	Next    Stmt
}

func (OptionalMatchStmt) patternNode() {}

func (m *OptionalMatchStmt) GetPattern() Patn {
	return m.Pattern
}

func (m *OptionalMatchStmt) GetNext() Stmt {
	return m.Next
}
