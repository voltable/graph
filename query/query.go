package query

import "github.com/RossMerr/Caudex.Graph/vertices"

type (
	// Iterator is an alias for function to iterate over data.
	Iterator func() (item interface{}, ok bool)

	// MatchPattern is used for describing the pattern to search for represents a walk through a the graph and consists of a sequence of alternating nodes and relationships.
	//MatchPattern map[string]Path

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
