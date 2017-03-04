package traversing

import (
	"github.com/RossMerr/Caudex.Graph/graph"
	"github.com/RossMerr/Caudex.Graph/storageEngines"
	"github.com/oleiade/lane"
)

// BFS Breadth-first Search
func BFS(se storageEngines.StorageEngine, root *graph.Vertex, fn func(*graph.Vertex) bool) []*graph.Vertex {
	queue := lane.NewQueue()
	var marked map[string]bool
	marked[root.ID()] = true
	queue.Enqueue(root)
	var results []*graph.Vertex

	for !queue.Empty() {
		i := queue.Dequeue()

		v, ok := i.(*graph.Vertex)
		if ok {
			if fn(v) {
				results = append(results, v)
			}

			for _, e := range v.Edges() {
				if !marked[e.ID()] {
					if v, err := se.Find(e.ID()); err == nil {
						queue.Enqueue(v)
						marked[v.ID()] = true
					}
				}
			}
		}
	}

	return results
}
