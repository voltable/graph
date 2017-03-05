package bfs

import (
	graph "github.com/RossMerr/Caudex.Graph"
	"github.com/RossMerr/Caudex.Graph/traversal"
	"github.com/RossMerr/Caudex.Graph/vertices"
	"github.com/oleiade/lane"
)

func init() {
	traversal.RegisterTraversal(TraversalType, traversal.TraversalRegistration{
		NewFunc: newTraversal,
	})
}

const TraversalType = "bfs"

type BFS struct {
	g graph.Graph
}

func newTraversal() (traversal.Traversal, error) {
	bfs := &BFS{}
	return bfs, nil
}

// BFS Breadth-first Search
func (bgs *BFS) Query(root *vertices.Vertex, predicate func(*vertices.Vertex) bool) []*vertices.Vertex {
	queue := lane.NewQueue()
	var marked map[string]bool
	marked[root.ID()] = true
	queue.Enqueue(root)
	var results []*vertices.Vertex
Loop:
	for !queue.Empty() {
		i := queue.Dequeue()

		v, ok := i.(*vertices.Vertex)
		if ok {
			if predicate(v) {
				results = append(results, v)
				continue Loop
			}

			for _, e := range v.Edges() {
				if !marked[e.ID()] {
					if v, err := bgs.g.Find(e.ID()); err == nil {
						queue.Enqueue(v)
						marked[v.ID()] = true
					}
				}
			}
		}
	}

	return results
}
