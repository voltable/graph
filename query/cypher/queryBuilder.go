package cypher

import (
	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"

	"github.com/RossMerr/Caudex.Graph/query"
)

// ToQueryPath converts a cypher.Stmt to a QueryPath the queryPath is used to walk the graph
func ToQueryPath(stmt ast.Stmt, toPredicateVertex func(*ast.VertexPatn) query.PredicateVertex,
	toPredicateEdge func(patn *ast.EdgePatn) query.PredicateEdge) (*query.QueryPath, error) {
	q := query.NewQueryPath()
	var next func(query.Path)
	next = q.SetNext
	if pattern, ok := isPattern(stmt); ok {
		for pattern != nil {
			if v, ok := pattern.(*ast.VertexPatn); ok && v != nil {
				pvp := query.PredicateVertexPath{PredicateVertex: toPredicateVertex(v)}
				next(&pvp)
				next = pvp.SetNext
				pattern = v.Edge

			} else if e, ok := pattern.(*ast.EdgePatn); ok && e != nil {
				pvp := query.PredicateEdgePath{PredicateEdge: toPredicateEdge(e)}
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

	next = nil
	return q, nil
}

func isPattern(item ast.Stmt) (ast.Patn, bool) {
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
