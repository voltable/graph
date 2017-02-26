package graphs

import (
	"errors"

	"github.com/oleiade/lane"
)

var (
	errNotFound = errors.New("Not found")
)

type QueryOperation struct {
	StorageEngine
	rawQuery string
}

func NewQueryOperation(db StorageEngine) *QueryOperation {
	q := QueryOperation{db, ""}
	return &q
}

func (q *QueryOperation) Query(s string) ([]Vertex, error) {

	return nil, nil
}

// DFS Depth-first search
func (q *QueryOperation) DFS(root *Vertex, fn func(*Vertex) bool) (*Vertex, error) {
	stack := lane.NewStack()
	var marked map[string]bool
	stack.Push(root)

	for !stack.Empty() {
		i := stack.Pop()
		v, ok := i.(*Vertex)
		if ok {
			if fn(v) {
				return v, nil
			}

			if !marked[v.ID] {
				marked[v.ID] = true
				for _, e := range v.Edges() {
					if v, err := q.Find(e.id); err == nil {
						stack.Push(v)
						marked[v.ID] = false
					}
				}
			}
		}
	}

	return nil, errNotFound
}

// BFS Breadth-first Search
func (q *QueryOperation) BFS(root *Vertex, fn func(*Vertex) bool) (*Vertex, error) {
	queue := lane.NewQueue()
	var marked map[string]bool
	marked[root.ID] = true
	queue.Enqueue(root)

	for !queue.Empty() {
		i := queue.Dequeue()

		v, ok := i.(*Vertex)
		if ok {
			if fn(v) {
				return v, nil
			}

			for _, e := range v.Edges() {
				if !marked[e.id] {
					if v, err := q.Find(e.id); err == nil {
						queue.Enqueue(v)
						marked[v.ID] = true
					}
				}
			}
		}
	}

	return nil, errNotFound
}
