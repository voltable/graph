package ast

type ReadStmt interface {
	Stmt
}

type MatchStmt struct {
	Pattern *VertexStmt
	Next    Stmt
}

type OptionalMatchStmt struct {
	Pattern *VertexStmt
	Next    Stmt
}
