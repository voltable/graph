package cypher

import (
	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
)

type Parts interface {
	ToQueryPart(stmt ast.Stmt) ([]*QueryPart, error)
}

// QueryPart is one part of a explicitly separate query parts
type QueryPart struct {
	Path  query.Path
	Where *ast.WhereStmt
}

var _ Parts = (*cypherParts)(nil)

type cypherParts struct {
}

func NewParts() Parts {
	return &cypherParts{}
}

// ToQueryPath converts a cypher.Stmt to a QueryPath the queryPath is used to walk the graph
func (qq cypherParts) ToQueryPart(stmt ast.Stmt) ([]*QueryPart, error) {
	arr := make([]*QueryPart, 0)
	q, _ := NewPath()
	qp := QueryPart{Path: q}
	arr = append(arr, &qp)
	var next func(query.Path)
	next = q.SetNext
	if pattern, ok := IsPattern(stmt); ok {
		for pattern != nil {
			if v, ok := pattern.(*ast.VertexPatn); ok && v != nil {
				pvp := query.PredicateVertexPath{PredicateVertex: v.ToPredicateVertex()}
				next(&pvp)
				next = pvp.SetNext
				pattern = v.Edge

			} else if e, ok := pattern.(*ast.EdgePatn); ok && e != nil {
				pvp := query.PredicateEdgePath{PredicateEdge: e.ToPredicateEdge()}
				if e.Body != nil {
					pvp.SetLength(e.Body.LengthMinimum, e.Body.LengthMaximum)
				}
				next(&pvp)
				next = pvp.SetNext
				pattern = e.Vertex
				// don't like making the WhereStmt a pattern
			} else if w, ok := pattern.(*ast.WhereStmt); ok && w != nil {
				//todo this might not be right
				qp.Where = w
				break
			} else {
				break
			}
		}
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
