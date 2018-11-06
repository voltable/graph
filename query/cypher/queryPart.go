package cypher

import (
	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
)

type Parts interface {
	ToQueryPart(stmt ast.Clauses) ([]*QueryPart, error)
}

// QueryPart is one part of a explicitly separate query parts
type QueryPart struct {
	Return   *ast.ReturnStmt
	Where    *ast.WhereStmt
	Patterns []ast.Patn
}

// Predicate gets the Predicate from the query where statment
func (qp *QueryPart) Predicate() ast.Expr {
	if qp.Where != nil {
		if qp.Where.Predicate != nil {
			return qp.Where.Predicate
		}
	}

	return nil
}

// Maps gets the []*MapProjectionStmt from the query return statment
func (qp *QueryPart) Maps() []*ast.ProjectionMapStmt {

	if qp.Return != nil {
		if qp.Return.Maps != nil {
			return qp.Return.Maps
		}
	}

	return nil
}

func (qp *QueryPart) Variable() string {
	if len(qp.Patterns) > 0 {
		e := qp.Patterns[0]
		return e.V()
	}
	return ""
}

var _ Parts = (*cypherParts)(nil)

type cypherParts struct {
}

func NewParts() Parts {
	return &cypherParts{}
}

// ToQueryPath converts a cypher.Stmt to a QueryPath the queryPath is used to walk the graph
func (qq cypherParts) ToQueryPart(stmt ast.Clauses) ([]*QueryPart, error) {
	arr := make([]*QueryPart, 0)
	qp := QueryPart{Patterns: make([]ast.Patn, 0)}
	arr = append(arr, &qp)
	if pattern, ok := IsPattern(stmt); ok {
		for pattern != nil {
			qp.Patterns = append(qp.Patterns, pattern)
			if edge, ok := pattern.(*ast.EdgePatn); ok && edge != nil {
				pattern = edge.Next()
			} else if vertex, ok := pattern.(*ast.VertexPatn); ok && vertex != nil {
				pattern = vertex.Next()
			} else {
				break
			}
		}
	}

	if where, ok := stmt.GetNext().(*ast.WhereStmt); ok {
		qp.Where = where
		stmt = stmt.GetNext()
	}

	if returns, ok := stmt.GetNext().(*ast.ReturnStmt); ok {
		qp.Return = returns
		stmt = stmt.GetNext()
	}

	return arr, nil
}

func IsPattern(item ast.Stmt) (ast.Patn, bool) {
	if b, ok := item.(*ast.DeleteStmt); ok {
		return b.Pattern, true
	} else if b, ok := item.(*ast.CreateStmt); ok {
		return b.Pattern, true
	} else if b, ok := item.(*ast.OptionalMatchStmt); ok {
		return b.Pattern, true
	} else if b, ok := item.(*ast.MatchStmt); ok {
		return b.Pattern, true
	}
	return nil, false
}
