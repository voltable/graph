package cypher

import "github.com/RossMerr/Caudex.Graph/vertices"

type ReadStatement interface {
	Statement
}

type MatchStatement struct {
	Pattern *VertexStatement
	next	Statement
}

func (m *MatchStatement) Next() (Statement, bool) {
	if m.next != nil {
		return m.next, true
	}

	return nil, false
}

type OptionalMatchStatement struct {
	Pattern *VertexStatement
	next	Statement
}

func (m *OptionalMatchStatement) Next() (Statement, bool) {
	if m.next != nil {
		return m.next, true
	}

	return nil, false
}

// PredicateVertex apply the predicate over the vertex
type PredicateVertex func(v *vertices.Vertex) bool

type WhereStatement struct {
	Where PredicateVertex
	next	Statement
}

func (m *WhereStatement) Next() (Statement, bool) {
	if m.next != nil {
		return m.next, true
	}

	return nil, false
}