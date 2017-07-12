package query

import (
	"github.com/RossMerr/Caudex.Graph/vertices"
)

type (

	// Query used to hold predicates that makes up a query path
	Query struct {
		Vertices []PredicateVertex
		Edges    []PredicateEdge
	}

	// QueryPath is a walk in the graph in a alternating sequence of vertices and edges
	QueryPath struct {
		next PredicateVertexPath
	}

	PatternLength struct {
		LengthMinimum uint
		LengthMaximum uint
	}

	VertexNext interface {
		Next() EdgeNext
		Length() PatternLength
	}

	EdgeNext interface {
		Next() VertexNext
		Length() PatternLength
	}

	PredicateVertexPath struct {
		PredicateVertex
		next   EdgeNext
		length PatternLength
	}

	PredicateEdgePath struct {
		PredicateEdge
		next   VertexNext
		length PatternLength
	}

	// Iterator is an alias for function to iterate over data.
	Iterator func() (item interface{}, ok bool)

	// PredicateVertex apply the predicate over the vertex
	PredicateVertex func(v *vertices.Vertex) bool

	//PredicateEdge apply the predicate over the edge
	PredicateEdge func(*vertices.Edge) bool
)

func NewQuery() *Query {
	q := &Query{Vertices: []PredicateVertex{}, Edges: []PredicateEdge{}}
	return q
}

func (p PredicateVertexPath) Next() EdgeNext {
	return p.next
}

func (p PredicateEdgePath) Next() VertexNext {
	return p.next
}

func (p PredicateVertexPath) Length() PatternLength {
	return p.length
}

func (p PredicateEdgePath) Length() PatternLength {
	return p.length
}
