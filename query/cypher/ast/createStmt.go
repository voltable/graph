package ast

// CreateStmt used to create nodes and relationships.
type CreateStmt struct {
	Pattern *Patn
	Next    Clauses
}

func (CreateStmt) patternNode() {}

func (m *CreateStmt) GetPattern() *Patn {
	return m.Pattern
}

func (m *CreateStmt) GetNext() Clauses {
	return m.Next
}
