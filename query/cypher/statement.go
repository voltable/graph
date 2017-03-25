package cypher


type ClauseStatement struct {
	Statement
	//Clause Token	
	//Pattern *VertexStatement
	//SubClause Token
}

type Statement interface {
	Next() (Statement, bool)
}

