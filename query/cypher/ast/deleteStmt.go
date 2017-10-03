package ast

// DeleteStmt used to delete graph elements — nodes, relationships or paths.
type DeleteStmt struct {
	Pattern Patn
	Next    Stmt
}

func (DeleteStmt) patternNode() {}

func (m *DeleteStmt) GetPattern() Patn {
	return m.Pattern
}

func (m *DeleteStmt) GetNext() Stmt {
	return m.Next
}
