package query

import (
	"fmt"
	"sort"

	"github.com/RossMerr/Caudex.Graph/enumerables"
	"github.com/RossMerr/Caudex.Graph/storage"
)

type Plan struct {
	explored        map[string]bool
	storage         storage.Storage
	predicateVertex PredicateVertex
	predicateEdge   PredicateEdge
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

		} else {
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
		}
		sort.Sort(frontier)
		return true
	}

	return false
}

func SearchPlan(storage storage.Storage, iterator enumerables.Iterator, vertexPath *PredicateVertexPath, edgePath *PredicateEdgePath) (iteratorFrontier IteratorFrontier, err error) {
	path := NewVertexPath(iterator, storage, vertexPath.Variable)
	plan := &Plan{explored: make(map[string]bool),
		storage:         storage,
		predicateVertex: vertexPath.PredicateVertex,
		predicateEdge:   edgePath.PredicateEdge,
	}

	in := make(chan *Frontier)
	out := make(chan *Frontier)

	go worker(plan, in, out)

	forEach(in, path.Iterate)

	temp := 0
	for r := range out {
		temp++
		if temp == 19 {
			fmt.Printf("end")
			close(in)
			close(out)
			break
		}
		in <- r
	}

	return nil, nil
}

func forEach(in chan<- *Frontier, i IteratorFrontier) {
	for frontier, ok := i(); ok != Failed; frontier, ok = i() {
		in <- frontier
	}
}

func worker(plan *Plan, in <-chan *Frontier, out chan<- *Frontier) {
	for f := range in {
		if plan.UniformCostSearch(f) {
			out <- f
		}
	}
}
