package query

import (
	"errors"

	"github.com/RossMerr/Caudex.Graph/vertices"
)

type (
	// VariableLength a range for pattern matching
	VariableLength struct {
		Minimum uint
		Maximum uint
	}

	// VertexNext is the Vertex part of the QueryPath sequence
	VertexNext interface {
	}

	// EdgeNext is the Edge part of the QueryPath sequence
	EdgeNext interface {
		Length() *VariableLength
	}

	// Path is a walk in the graph in a alternating sequence of vertices and edges
	Path interface {
		Next() Path
		SetNext(path Path)
	}

	// Root is the start of the path
	Root struct {
		next Path
	}

	// PredicateVertexPath is the Vertex implementation part of the QueryPath sequence
	PredicateVertexPath struct {
		PredicateVertex
		next Path
	}

	// PredicateEdgePath is the Edge implementation part of the QueryPath sequence
	PredicateEdgePath struct {
		PredicateEdge
		next Path

		length VariableLength
	}

	// Iterator is an alias for function to iterate over data.
	Iterator func() (item interface{}, ok bool)

	// PredicateVertex apply the predicate over the vertex
	PredicateVertex func(v *vertices.Vertex) bool

	// PredicateEdge apply the predicate over the edge
	PredicateEdge func(e *vertices.Edge) bool

	// PredicateExpression apply the predicate over the Expr
	PredicateExpression func(e *Path) bool
)

var (
	errPredicateVertexPath = errors.New("Expected PredicateVertexPath")
	errPredicateEdgePath   = errors.New("Expected PredicateEdgePath")
)

// NewPath creates a Path object used to walk the graph
func NewPath() (Path, error) {
	path := &Root{}
	return path, nil
}

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
func (p *Root) Next() Path {
	return p.next
}

// SetNext sets the next Edge in the QueryPath
func (p *Root) SetNext(path Path) {
	if v, ok := path.(*PredicateVertexPath); ok {
		p.next = v
	}
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

// Length returns the range of the lengths for pattern matching
func (p *PredicateEdgePath) Length() *VariableLength {
	return &p.length
}

// SetLength sets the lengths for the pattern matching
func (p *PredicateEdgePath) SetLength(minimum uint, maximum uint) {
	length := VariableLength{Maximum: maximum, Minimum: minimum}
	p.length = length
}
