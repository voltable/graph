package dfs

import (
	"github.com/RossMerr/Caudex.Graph"
	"github.com/RossMerr/Caudex.Graph/traversal"
	"github.com/RossMerr/Caudex.Graph/vertices"
	"github.com/oleiade/lane"
)

func init() {
	traversal.RegisterTraversal(TraversalType, traversal.TraversalRegistration{
		NewFunc: newTraversal,
	})
}

const TraversalType = "dfs"

type DFS struct {
	g graph.Graph
}

func newTraversal() (traversal.Traversal, error) {
	dfs := &DFS{}
	return dfs, nil
}

// DFS Depth-first search
func (dfs *DFS) Query(root *vertices.Vertex, fn func(*vertices.Vertex) bool) []*vertices.Vertex {
	stack := lane.NewStack()
	var marked map[string]bool
	stack.Push(root)
	var results []*vertices.Vertex

	for !stack.Empty() {
		i := stack.Pop()
		v, ok := i.(*vertices.Vertex)
		if ok {
			if fn(v) {
				results = append(results, v)
			}

			if !marked[v.ID()] {
				marked[v.ID()] = true
				for _, e := range v.Edges() {
					if v, err := dfs.g.Find(e.ID()); err == nil {
						stack.Push(v)
						marked[v.ID()] = false
					}
				}
			}
		}
	}

	return results
}
