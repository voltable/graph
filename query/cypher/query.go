package cypher

import (
	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
)

// Query is the internal Query interface for use in cypher
type Query struct {
	stmt    ast.Stmt
	path    query.Path
	results []interface{}
}

var _ query.Query = (*Query)(nil)
var _ query.QueryInternal = (*Query)(nil)

// NewQuery create the query struct with the Path and Stmt, both Path and Stmt are latter used for travering and filtering
func NewQuery(stmt ast.Stmt, path query.Path) *Query {
	return &Query{stmt: stmt, path: path}
}

// Results returns the result array
func (q *Query) Results() []interface{} {
	return q.results
}

// SetResults is used by the internal query interface to update results
func (q *Query) SetResults(results []interface{}) {
	q.results = results
}

// Path return the query path
func (q *Query) Path() query.Path {
	return q.path
}
