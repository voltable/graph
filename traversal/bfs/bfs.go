package bfs

import (
	"github.com/RossMerr/Caudex.Graph/graph/vertices"
	"github.com/RossMerr/Caudex.Graph/storageEngines"
	"github.com/RossMerr/Caudex.Graph/traversal"
	"github.com/oleiade/lane"
)

func init() {
	traversal.RegisterTraversal(TraversalType, traversal.TraversalRegistration{
		NewFunc: newTraversal,
	})
}

const TraversalType = "bfs"

type BFS struct {
	se storageEngines.StorageEngine
}

func newTraversal() (traversal.Traversal, error) {
	bfs := &BFS{}
	return bfs, nil
}

// BFS Breadth-first Search
func (bgs *BFS) Query(root *vertices.Vertex, fn func(*vertices.Vertex) bool) []*vertices.Vertex {
	queue := lane.NewQueue()
	var marked map[string]bool
	marked[root.ID()] = true
	queue.Enqueue(root)
	var results []*vertices.Vertex

	for !queue.Empty() {
		i := queue.Dequeue()

		v, ok := i.(*vertices.Vertex)
		if ok {
			if fn(v) {
				results = append(results, v)
			}

			for _, e := range v.Edges() {
				if !marked[e.ID()] {
					if v, err := bgs.se.Find(e.ID()); err == nil {
						queue.Enqueue(v)
						marked[v.ID()] = true
					}
				}
			}
		}
	}

	return results
}
