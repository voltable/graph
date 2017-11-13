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
	storage    storage.Storage
	wg         *sync.WaitGroup
	predicates []interface{}
	Depth      int
}

func NewPlan(storage storage.Storage) *Plan {
	plan := &Plan{
		storage: storage,
		wg:      &sync.WaitGroup{},
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

var count int = 0

func (t *Plan) UniformCostSearch(frontier *Frontier) bool {
	if frontier.Len() > 0 {
		count++
		vertices, cost := frontier.Pop()

		depth := len(vertices)
		vertex := vertices[depth-1]
		predicateDepth := depth + (depth - 1)

		fmt.Printf("count %+v\n", count)

		fmt.Printf("depth %+v\n", depth)

		fmt.Printf("vertex %+v\n", vertex.ID())
		if _, ok := frontier.Explored[vertex.Vertex.ID()]; !ok {
			if pv := t.predicateVertex(predicateDepth - 1); pv != nil {
				if variable, p := pv(vertex.Vertex); p == Matched {
					vertex.Variable = variable
					frontier.Append(vertices, cost, p)
					sort.Sort(frontier)
					frontier.Explored[vertex.ID()] = true
					fmt.Printf("result: %+v\n", predicateDepth == t.Depth)
					return predicateDepth == t.Depth
				}
			}
		}

		if pe := t.predicateEdge(predicateDepth); pe != nil {
			for _, e := range vertex.Edges() {
				if _, ok := frontier.Explored[e.ID()]; !ok {
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
	fmt.Printf("result: false\n")
	return false
}

func (t *Plan) SearchPlan(iterator enumerables.Iterator, predicates []interface{}) (iteratorFrontier IteratorFrontier, err error) {
	t.predicates = predicates
	t.Depth = len(predicates)

	start, results := t.forEach(toFontier(iterator, t.variableVertex()))
	t.worker(start, results)

	return func() (*Frontier, Traverse) {
		f, opened := <-results
		if opened {
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
				t.wg.Done()
			} else if f.Len() > 0 {
				go func() {
					out <- f
				}()
			} else {
				t.wg.Done()
			}
		}

	}()
}

func (t *Plan) forEach(i IteratorFrontier) (chan *Frontier, chan *Frontier) {
	out := make(chan *Frontier)
	results := make(chan *Frontier)
	go func() {
		defer close(out)
		defer close(results)

		for frontier, ok := i(); ok != Failed; frontier, ok = i() {
			t.wg.Add(1)
			out <- frontier
		}

		t.wg.Wait()

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
