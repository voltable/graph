package cypher

import (
	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"

	"github.com/RossMerr/Caudex.Graph/query"
)

// ToQuery converts a cypher.Stmt to a Query to keep it all abstracted
func ToQuery(stmt ast.Stmt) (*query.Query, error) {

	q := query.NewQuery()
	var pv []query.PredicateVertex
	var pe []query.PredicateEdge

	if b, ok := stmt.(*ast.CreateStmt); ok {
		pattern := b.Pattern
		for pattern != nil {
			if v, ok := b.Pattern.(*ast.VertexPatn); ok {
				pv = append(pv, v.ToPredicateVertex())
				pattern = v.Edge
			} else if e, ok := b.Pattern.(*ast.EdgePatn); ok {
				pe = append(pe, e.ToPredicateEdge())
				pattern = e.Vertex
			}
		}
	}

	q.Edges = pe
	q.Vertices = pv
	return q, nil
}
