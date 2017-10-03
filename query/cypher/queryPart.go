package cypher

import (
	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
)

type Parts interface {
	ToQueryPart(stmt ast.Clauses) ([]*QueryPart, error)
}

// QueryPart is one part of a explicitly separate query parts
type QueryPart struct {
	Path  query.Path
	Where *ast.WhereStmt
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

var _ Parts = (*cypherParts)(nil)

type cypherParts struct {
}

func NewParts() Parts {
	return &cypherParts{}
}

// ToQueryPath converts a cypher.Stmt to a QueryPath the queryPath is used to walk the graph
func (qq cypherParts) ToQueryPart(stmt ast.Clauses) ([]*QueryPart, error) {
	arr := make([]*QueryPart, 0)
	q, _ := NewPath()
	qp := QueryPart{Path: q}
	arr = append(arr, &qp)
	var next func(query.Path)
	next = q.SetNext
	if pattern, ok := IsPattern(stmt); ok {
		for pattern != nil {
			if v, ok := pattern.(*ast.VertexPatn); ok && v != nil {
				pvp := query.PredicateVertexPath{PredicateVertex: v.ToPredicateVertex(), Variable: v.Variable}
				next(&pvp)
				next = pvp.SetNext
				pattern = v.Edge
			} else if e, ok := pattern.(*ast.EdgePatn); ok && e != nil {
				pvp := query.PredicateEdgePath{PredicateEdge: e.ToPredicateEdge(), Variable: e.Variable}
				if e.Body != nil {
					pvp.SetLength(e.Body.LengthMinimum, e.Body.LengthMaximum)
				}
				next(&pvp)
				next = pvp.SetNext
				pattern = e.Vertex
			} else {
				break
			}
		}

	}

	if where, ok := stmt.GetNext().(*ast.WhereStmt); ok {
		qp.Where = where
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
