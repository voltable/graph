package traversing

import (
	"github.com/RossMerr/Caudex.Graph/graph"
	"github.com/RossMerr/Caudex.Graph/storageEngines"
	"github.com/oleiade/lane"
)

// DFS Depth-first search
func DFS(se storageEngines.StorageEngine, root *graph.Vertex, fn func(*graph.Vertex) bool) []*graph.Vertex {

	stack := lane.NewStack()
	var marked map[string]bool
	stack.Push(root)
	var results []*graph.Vertex

	for !stack.Empty() {
		i := stack.Pop()
		v, ok := i.(*graph.Vertex)
		if ok {
			if fn(v) {
				results = append(results, v)
			}

			if !marked[v.ID()] {
				marked[v.ID()] = true
				for _, e := range v.Edges() {
					if v, err := se.Find(e.ID()); err == nil {
						stack.Push(v)
						marked[v.ID()] = false
					}
				}
			}
		}
	}

	return results
}
