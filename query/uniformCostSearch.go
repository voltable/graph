package query

import (
	"fmt"
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

	isResultsOpen bool
	mux           sync.Mutex
	Results       chan *Frontier
}

func NewPlan(storage storage.Storage, predicateVertex PredicateVertex, predicateEdge PredicateEdge) *Plan {
	plan := &Plan{explored: make(map[string]bool),
		storage:         storage,
		predicateVertex: predicateVertex,
		predicateEdge:   predicateEdge,
		isResultsOpen:   true,
		Results:         make(chan *Frontier),
	}
	return plan
}

func (p *Plan) CloseResults() {
	p.mux.Lock()
	if p.isResultsOpen {
		close(p.Results)
		p.isResultsOpen = false
	}
	p.mux.Unlock()
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

	start := forEach(path.Iterate)
	out := worker(plan, start)

	go func() {
		in := out
		out = worker(plan, in)
	}()

	for f := range plan.Results {
		fmt.Printf("results %+v\n", f)
	}

	return nil, nil
}

func forEach(i IteratorFrontier) chan *Frontier {
	out := make(chan *Frontier)
	go func() {
		for frontier, ok := i(); ok != Failed; frontier, ok = i() {
			out <- frontier
		}
		close(out)
	}()
	return out
}

func worker(plan *Plan, in <-chan *Frontier) <-chan *Frontier {
	out := make(chan *Frontier)
	go func() {
		canClose := true
		for f := range in {
			if plan.UniformCostSearch(f) {
				plan.Results <- f
			} else if f.Len() > 0 {
				canClose = false
				out <- f
			}
		}
		if canClose {
			plan.CloseResults()
		}
		close(out)
	}()
	return out
}
