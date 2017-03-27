package ast

type ClauseStmt struct {
	Statement
	//Clause Token
	//Pattern *VertexStatement
	//SubClause Token
}

type Statement interface {
	//Next() (Statement, bool)
}
