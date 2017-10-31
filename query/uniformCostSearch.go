package query

import (
	"sort"
	"sync"

	"github.com/RossMerr/Caudex.Graph/enumerables"
	"github.com/RossMerr/Caudex.Graph/storage"
)

type Plan struct {
	explored        map[string]bool
	storage         storage.Storage
	predicateVertex PredicateVertex
	predicateEdge   PredicateEdge
}

func NewPlan(storage storage.Storage, predicateVertex PredicateVertex, predicateEdge PredicateEdge) *Plan {
	plan := &Plan{explored: make(map[string]bool),
		storage:         storage,
		predicateVertex: predicateVertex,
		predicateEdge:   predicateEdge,
	}
	return plan
}

func (t *Plan) UniformCostSearch(frontier *Frontier) bool {
	if frontier.Len() > 0 {
		vertices, cost := frontier.Pop()

		depth := len(vertices)
		vertex := vertices[depth-1]
		t.explored[vertex.ID()] = true

		if variable, p := t.predicateVertex(vertex.Vertex); p == Matched {
			vertex.Variable = variable
			frontier.Append(vertices, cost, p)
			sort.Sort(frontier)
			return true
		}

		for _, e := range vertex.Edges() {
			if _, ok := t.explored[e.ID()]; !ok {
				if variable, p := t.predicateEdge(e, uint(depth)); p == Visiting || p == Matching {
					if v, err := t.storage.Fetch(e.ID()); err == nil {
						fv := &FrontierVertex{Vertex: v, Variable: variable}
						frontier.Append(append(vertices, fv), cost+e.Weight, p)
					}
				}
			}
		}

		sort.Sort(frontier)
	}

	return false
}

func SearchPlan(storage storage.Storage, iterator enumerables.Iterator, vertexPath *PredicateVertexPath, edgePath *PredicateEdgePath) (iteratorFrontier IteratorFrontier, err error) {
	path := NewVertexPath(iterator, storage, vertexPath.Variable)
	plan := NewPlan(storage, vertexPath.PredicateVertex, edgePath.PredicateEdge)
	wg := &sync.WaitGroup{}

	start, results := forEach(path.Iterate, wg)
	worker(plan, start, results, wg)

	return func() (*Frontier, Traverse) {
		f, ok := <-results
		if ok {
			return f, Matched
		}
		return nil, Failed
	}, nil
}

func worker(plan *Plan, out chan *Frontier, results chan *Frontier, wg *sync.WaitGroup) {
	go func() {
		for f := range out {
			if plan.UniformCostSearch(f) {
				results <- f
			} else if f.Len() > 0 {
				wg.Add(1)
				go func() {
					out <- f
				}()
			}
			wg.Done()
		}
	}()
}

func forEach(i IteratorFrontier, wg *sync.WaitGroup) (chan *Frontier, chan *Frontier) {
	out := make(chan *Frontier)
	results := make(chan *Frontier)
	go func() {
		for frontier, ok := i(); ok != Failed; frontier, ok = i() {
			wg.Add(1)
			out <- frontier
		}
		go func() {
			wg.Wait()
			close(out)
			close(results)
		}()
	}()
	return out, results
}
