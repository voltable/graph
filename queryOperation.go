package graphs

import "github.com/oleiade/lane"

type QueryOperation struct {
	StorageEngine
	rawQuery string
}

func NewQueryOperation(db StorageEngine) *QueryOperation {
	q := QueryOperation{db, ""}
	return &q
}

// DFS Depth-first search
func (q *QueryOperation) DFS(root *Vertex, fn func(*Vertex) bool) *QueryResult {
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
					if v, err := q.Find(e.id); err == nil {
						stack.Push(v)
						marked[v.id] = false
					}
				}
			}
		}
	}

	return NewQueryResult(results)
}

// BFS Breadth-first Search
func (q *QueryOperation) BFS(root *Vertex, fn func(*Vertex) bool) *QueryResult {
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
					if v, err := q.Find(e.id); err == nil {
						queue.Enqueue(v)
						marked[v.id] = true
					}
				}
			}
		}
	}

	return NewQueryResult(results)
}
