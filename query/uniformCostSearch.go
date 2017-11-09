package query

import (
	"fmt"
	"sort"
	"sync"

	"github.com/RossMerr/Caudex.Graph/enumerables"
	"github.com/RossMerr/Caudex.Graph/storage"
	"github.com/RossMerr/Caudex.Graph/vertices"
)

type Plan struct {
	explored   map[string]bool
	storage    storage.Storage
	wg         *sync.WaitGroup
	predicates []interface{}
}

func NewPlan(storage storage.Storage) *Plan {
	plan := &Plan{
		explored: make(map[string]bool),
		storage:  storage,
		wg:       &sync.WaitGroup{},
	}
	return plan
}

func (t *Plan) variableVertex() string {
	e := t.predicates[0]
	if pv, ok := e.(*PredicateVertexPath); ok {
		return pv.Variable
	}
	return ""
}

func (t *Plan) predicateVertex(i int) PredicateVertex {
	if i < len(t.predicates) {
		e := t.predicates[i]
		if pv, ok := e.(*PredicateVertexPath); ok {
			return pv.PredicateVertex
		}
	}
	return nil
}

func (t *Plan) predicateEdge(i int) PredicateEdge {
	if i < len(t.predicates) {
		e := t.predicates[i]
		if pv, ok := e.(*PredicateEdgePath); ok {
			return pv.PredicateEdge
		}
	}
	return nil
}

func (t *Plan) UniformCostSearch(frontier *Frontier) bool {
	if frontier.Len() > 0 {
		vertices, cost := frontier.Pop()

		depth := len(vertices)
		vertex := vertices[depth-1]
		t.explored[vertex.ID()] = true

		if pv := t.predicateVertex(depth - 1); pv != nil {
			if variable, p := pv(vertex.Vertex); p == Matched {
				vertex.Variable = variable
				frontier.Append(vertices, cost, p)
				sort.Sort(frontier)
				return true
			}
		}

		for _, e := range vertex.Edges() {
			if _, ok := t.explored[e.ID()]; !ok {
				if pe := t.predicateEdge(depth); pe != nil {
					if variable, p := pe(e, uint(depth)); p == Visiting || p == Matching {
						if v, err := t.storage.Fetch(e.ID()); err == nil {
							fv := &FrontierVertex{Vertex: v, Variable: variable}
							frontier.Append(append(vertices, fv), cost+e.Weight, p)
						}
					}
				}
			}
		}

		sort.Sort(frontier)

	}

	return false
}

var count = 0

func (t *Plan) SearchPlan(iterator enumerables.Iterator, predicates []interface{}) (iteratorFrontier IteratorFrontier, err error) {
	t.predicates = predicates

	start, results := t.forEach(toFontier(iterator, t.variableVertex()))
	t.worker(start, results)

	return func() (*Frontier, Traverse) {
		if count >= 1 {
			fmt.Printf("%+v", results)
		}
		count++
		f, ok := <-results
		if ok {
			return f, Matched
		}
		return nil, Failed
	}, nil
}

func (t *Plan) worker(out chan *Frontier, results chan *Frontier) {
	go func() {
		for f := range out {
			if t.UniformCostSearch(f) {
				results <- f
			} else if f.Len() > 0 {
				t.wg.Add(1)
				go func() {
					out <- f
				}()
			}
			t.wg.Done()
		}
	}()
}

func (t *Plan) forEach(i IteratorFrontier) (chan *Frontier, chan *Frontier) {
	out := make(chan *Frontier)
	results := make(chan *Frontier)
	go func() {
		for frontier, ok := i(); ok != Failed; frontier, ok = i() {
			t.wg.Add(1)
			out <- frontier
		}
		go func() {
			t.wg.Wait()
			close(out)
			close(results)
		}()
	}()
	return out, results
}

func toFontier(i enumerables.Iterator, variable string) IteratorFrontier {
	return func() (*Frontier, Traverse) {
		for item, ok := i(); ok; item, ok = i() {
			if v, is := item.(*vertices.Vertex); is {
				f := NewFrontier(v, variable)
				return &f, Visiting
			}
		}
		return nil, Failed
	}
}
