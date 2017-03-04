package graphs

import "github.com/oleiade/lane"

// BFS Breadth-first Search
func (g *Graph) bfs(root *Vertex, fn func(*Vertex) bool) *QueryResult {
	queue := lane.NewQueue()
	var marked map[string]bool
	marked[root.id] = true
	queue.Enqueue(root)
	var results []*Vertex

	for !queue.Empty() {
		i := queue.Dequeue()

		v, ok := i.(*Vertex)
		if ok {
			if fn(v) {
				results = append(results, v)
			}

			for _, e := range v.Edges() {
				if !marked[e.id] {
					if v, err := g.db.Find(e.id); err == nil {
						queue.Enqueue(v)
						marked[v.id] = true
					}
				}
			}
		}
	}

	return NewQueryResult(results)
}
