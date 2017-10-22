package query

import (
	"errors"

	"github.com/RossMerr/Caudex.Graph/vertices"
)

type (
	// VertexNext is the Vertex part of the QueryPath sequence
	VertexNext interface {
	}

	// EdgeNext is the Edge part of the QueryPath sequence
	EdgeNext interface {
	}

	// Path is a walk in the graph in a alternating sequence of vertices and edges
	Path interface {
		Next() Path
		SetNext(path Path)
	}

	// PathParts is the separated parts of a walk in the graph
	PathParts []Path

	// PredicateVertexPath is the Vertex implementation part of the QueryPath sequence
	PredicateVertexPath struct {
		PredicateVertex
		next Path

		Variable string
	}

	// PredicateEdgePath is the Edge implementation part of the QueryPath sequence
	PredicateEdgePath struct {
		PredicateEdge
		next Path

		Variable string
	}

	// IteratorFrontier is an alias for function to iterate over Frontier.
	IteratorFrontier func() (item *Frontier, ok Traverse)

	// PredicateVertex apply the predicate over the vertex
	PredicateVertex func(v *vertices.Vertex) (string, Traverse)

	// PredicateEdge apply the predicate over the edge
	PredicateEdge func(e *vertices.Edge, depth uint) (string, Traverse)

	// PredicateExpression apply the predicate over the Expr
	PredicateExpression func(e *Path) bool
)

var (
	errPredicateVertexPath = errors.New("Expected PredicateVertexPath")
	errPredicateEdgePath   = errors.New("Expected PredicateEdgePath")
)

// SetNext sets the next Vertex or Edge in the Path
func SetNext(p Path, path Path) error {
	if _, ok := p.(*PredicateVertexPath); ok {
		if v, ok := path.(*PredicateEdgePath); ok {
			p.SetNext(v)
		} else {
			return errPredicateEdgePath
		}
	} else if _, ok := p.(*PredicateEdgePath); ok {
		if v, ok := path.(*PredicateVertexPath); ok {
			p.SetNext(v)
		} else {
			return errPredicateVertexPath
		}
	}

	return nil
}

// Next returns the next Vertex or Edge in the Path
func Next(p Path) Path {
	return p.Next()
}

// Next returns the next Edge in the QueryPath
func (p *PredicateVertexPath) Next() Path {
	return p.next
}

// SetNext sets the next Edge in the QueryPath
func (p *PredicateVertexPath) SetNext(path Path) {
	if v, ok := path.(*PredicateEdgePath); ok {
		p.next = v
	}
}

// Next returns the next Vertex in the QueryPath
func (p *PredicateEdgePath) Next() Path {
	return p.next
}

// SetNext sets the next Vertex in the QueryPath
func (p *PredicateEdgePath) SetNext(path Path) {
	if v, ok := path.(*PredicateVertexPath); ok {
		p.next = v
	}
}
