package cypher

import (
	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
	"github.com/RossMerr/Caudex.Graph/query/cypher/ir"
)

type Parts interface {
	ToQueryPart(stmt ast.Clauses) ([]*QueryPart, error)
}

// QueryPart is one part of a explicitly separate query parts
type QueryPart struct {
	Return     *ast.ReturnStmt
	Where      *ast.WhereStmt
	Predicates []interface{}
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
func (qp *QueryPart) Maps() []*ast.MapProjectionStmt {

	if qp.Return != nil {
		if qp.Return.Maps != nil {
			return qp.Return.Maps
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
	qp := QueryPart{Predicates: make([]interface{}, 0)}
	arr = append(arr, &qp)
	if pattern, ok := IsPattern(stmt); ok {
		for pattern != nil {
			if v, ok := pattern.(*ir.VertexPatn); ok && v != nil {
				qp.Predicates = append(qp.Predicates, v.ToPredicateVertexPath())
				pattern = v.Edge
			} else if e, ok := pattern.(*ir.EdgePatn); ok && e != nil {
				qp.Predicates = append(qp.Predicates, e.ToPredicateEdgePath())
				pattern = e.Vertex
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

func IsPattern(item ast.Stmt) (ir.Patn, bool) {
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
