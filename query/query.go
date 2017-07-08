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

	// QueryEngine is the interface that a queryEngine must implement
	QueryEngine interface {
		// Parser in a string which is your query you want to run, get back a vertexPath that is abstracted from any query language or AST
		Parser(string) (*Query, error)
	}

	// Iterator is an alias for function to iterate over data.
	Iterator func() (item interface{}, ok bool)

	// PredicateVertex apply the predicate over the vertex
	PredicateVertex func(v *vertices.Vertex) bool

	//PredicateEdge apply the predicate over the edge
	PredicateEdge func(*vertices.Edge) bool

	Path struct {
		Vertices []*vertices.Vertex
		Cost     float32
	}

	Frontier []*Path
)

func (f Frontier) Len() int               { return len(f) }
func (f Frontier) Swap(i, j int)          { f[i], f[j] = f[j], f[i] }
func (f Frontier) Less(i, j int) bool     { return f[i].Cost < f[j].Cost }
func (f Frontier) pop() (*Path, Frontier) { return f[0], f[1:] }
func (f Frontier) peek() *Path            { return f[0] }

func NewQuery() *Query {
	q := &Query{Vertices: []PredicateVertex{}, Edges: []PredicateEdge{}}
	return q
}
