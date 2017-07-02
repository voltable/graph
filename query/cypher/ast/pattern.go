package ast

import (
	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/vertices"
)

// Digraph represents the directed or undirected relationship on a Edge
// a character consisting of two joined letters; a ligature.
type Digraph int

const (
	// Undirected graphs have edges that do not have a direction. The edges indicate a two-way relationship, in that each edge can be traversed in both directions.
	Undirected Digraph = iota
	Inbound    Digraph = iota
	Outbound   Digraph = iota
)

// Patn all pattern nodes implement the Patn interface.
type Patn interface {
	patnNode()
}

type VertexPatn struct {
	Variable   string
	Properties map[string]interface{}
	Label      string

	Edge *EdgePatn
}

type EdgePatn struct {
	Relationship Digraph
	Body         *EdgeBodyStmt

	Vertex *VertexPatn
}

type EdgeBodyStmt struct {
	Variable      string
	Properties    map[string]interface{}
	Type          string
	LengthMinimum uint
	LengthMaximum uint
}

func (*VertexPatn) patnNode() {}
func (*EdgePatn) patnNode()   {}

// ToPredicateVertex creates a PredicateVertex out of the VertexPatn
func (patn *VertexPatn) ToPredicateVertex() query.PredicateVertex {
	return func(v *vertices.Vertex) bool {
		if patn.Label != v.Label() {
			return false
		}

		if len(patn.Properties) != v.PropertiesCount() {
			return false
		}

		for key, value := range patn.Properties {
			if v.Property(key) != value {
				return false
			}
		}

		return true
	}
}

// ToPredicateEdge creates a PredicateEdge out of the EdgePatn
func (patn *EdgePatn) ToPredicateEdge() query.PredicateEdge {
	return func(v *vertices.Edge) bool {
		if patn.Body.Type != v.RelationshipType() {
			return false
		}

		if len(patn.Body.Properties) != v.PropertiesCount() {
			return false
		}

		for key, value := range patn.Body.Properties {
			if v.Property(key) != value {
				return false
			}
		}

		return true
	}
}
