package cypher

import (
	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"

	"github.com/RossMerr/Caudex.Graph/query"
)

// ToQueryPath converts a cypher.Stmt to a QueryPath to keep it all abstracted
func ToQueryPath(stmt ast.Stmt) (*query.QueryPath, error) {
	q := query.NewQueryPath()
	var next func(query.Path)
	next = q.SetNext
	if b, ok := stmt.(*ast.CreateStmt); ok {
		pattern := b.Pattern
		for pattern != nil {
			if v, ok := b.Pattern.(*ast.VertexPatn); ok {
				pvp := query.PredicateVertexPath{PredicateVertex: v.ToPredicateVertex()}
				next(&pvp)
				next = pvp.SetNext
			} else if e, ok := b.Pattern.(*ast.EdgePatn); ok {
				pvp := query.PredicateEdgePath{PredicateEdge: e.ToPredicateEdge()}
				next(&pvp)
				next = pvp.SetNext
			}
		}
	}

	return q, nil
}
