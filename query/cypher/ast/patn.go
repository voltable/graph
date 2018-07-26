package ast

// Patn all pattern nodes implement the Patn interface.
type Patn interface {
	patnNode()
	Next() Patn
}

type EdgePatn struct {
	Variable     string
	Relationship Digraph
	Body         *EdgeBodyStmt

	Vertex *VertexPatn
}

func (*EdgePatn) patnNode() {}
func (s *EdgePatn) Next() Patn {
	return s.Vertex
}

type EdgeBodyStmt struct {
	Variable      string
	Properties    map[string]interface{}
	Type          string
	LengthMinimum uint
	LengthMaximum uint
}

type VertexPatn struct {
	Variable   string
	Properties map[string]interface{}
	Label      string

	Edge *EdgePatn
}

func (*VertexPatn) patnNode() {}
func (s *VertexPatn) Next() Patn {
	return s.Edge
}
