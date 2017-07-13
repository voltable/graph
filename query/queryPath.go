package query

import (
	"github.com/RossMerr/Caudex.Graph/vertices"
)

type (
	// QueryPath is a walk in the graph in a alternating sequence of vertices and edges
	QueryPath struct {
		next VertexNext
	}

	// VariableLength a range for pattern matching
	VariableLength struct {
		Minimum uint
		Maximum uint
	}

	// VertexNext is the Vertex part of the QueryPath sequence
	VertexNext interface {
		Next() EdgeNext
	}

	// EdgeNext is the Edge part of the QueryPath sequence
	EdgeNext interface {
		Next() VertexNext
		Length() *VariableLength
	}

	Path interface {
		path()
	}

	// PredicateVertexPath is the Vertex implementation part of the QueryPath sequence
	PredicateVertexPath struct {
		PredicateVertex
		next EdgeNext
	}

	// PredicateEdgePath is the Edge implementation part of the QueryPath sequence
	PredicateEdgePath struct {
		PredicateEdge
		next   VertexNext
		length VariableLength
	}

	// Iterator is an alias for function to iterate over data.
	Iterator func() (item interface{}, ok bool)

	// PredicateVertex apply the predicate over the vertex
	PredicateVertex func(v *vertices.Vertex) bool

	//PredicateEdge apply the predicate over the edge
	PredicateEdge func(*vertices.Edge) bool
)

func NewQueryPath() *QueryPath {
	q := &QueryPath{}
	return q
}

// Next returns the next Vertex in the QueryPath
func (p *QueryPath) Next() VertexNext {
	return p.next
}

// SetNext sets the next Edge in the QueryPath
func (p *QueryPath) SetNext(path Path) {
	if v, ok := path.(*PredicateVertexPath); ok {
		p.next = v
	}
}

// Next returns the next Edge in the QueryPath
func (p *PredicateVertexPath) Next() EdgeNext {
	return p.next
}

// SetNext sets the next Edge in the QueryPath
func (p *PredicateVertexPath) SetNext(path Path) {
	if v, ok := path.(*PredicateEdgePath); ok {
		p.next = v
	}
}

// Next returns the next Vertex in the QueryPath
func (p *PredicateEdgePath) Next() VertexNext {
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
func (p *PredicateEdgePath) SetLength(l VariableLength) {
	p.length = l
}

func (p *PredicateEdgePath) path()   {}
func (p *PredicateVertexPath) path() {}
