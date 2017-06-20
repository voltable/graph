package query

import "github.com/RossMerr/Caudex.Graph/vertices"

type (
	Query interface {
		Parser(string) (*VertexPath, error)
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
