package graphs

import "github.com/oleiade/lane"

// DFS Depth-first search
func (g *Graph) dfs(root *Vertex, fn func(*Vertex) bool) *QueryResult {

	stack := lane.NewStack()
	var marked map[string]bool
	stack.Push(root)
	var results []*Vertex

	for !stack.Empty() {
		i := stack.Pop()
		v, ok := i.(*Vertex)
		if ok {
			if fn(v) {
				results = append(results, v)
			}

			if !marked[v.id] {
				marked[v.id] = true
				for _, e := range v.Edges() {
					if v, err := g.db.Find(e.id); err == nil {
						stack.Push(v)
						marked[v.id] = false
					}
				}
			}
		}
	}

	return NewQueryResult(results)
}
