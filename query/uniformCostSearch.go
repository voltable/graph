package query

import (
	"sort"
	"sync"

	"github.com/RossMerr/Caudex.Graph"
)

type Traversal interface {
	SearchPlan(iterator IteratorFrontier, predicates []interface{}) (iteratorFrontier IteratorFrontier, err error)
}

type Plan struct {
	storage    graph.Storage
	wg         *sync.WaitGroup
	predicates []interface{}
	Depth      int
}

func NewPlan(storage graph.Storage) *Plan {
	plan := &Plan{
		storage: storage,
		wg:      &sync.WaitGroup{},
	}
	return plan
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
		queue := frontier.Pop()
		depth := len(queue.Parts)
		vertex := queue.Parts[depth-1].Object.(*graph.Vertex)
		if _, ok := frontier.Explored[vertex.ID()]; !ok {
			frontier.Explored[vertex.ID()] = true
			if pv := t.predicateVertex(depth - 1); pv != nil {
				if variable, p := pv(vertex); p == Matched {
					queue.Parts[depth-1].Variable = variable
					frontier.AppendQueue(queue)
					sort.Sort(frontier)
					return depth == t.Depth
				}
			}
		}
		if pe := t.predicateEdge(depth); pe != nil {
			for _, e := range vertex.Edges() {
				if _, ok := frontier.Explored[e.ID()]; !ok {
					if variable, p := pe(e, uint(depth)); p == Visiting || p == Matching {
						if v, err := t.storage.Fetch(e.ID()); err == nil {
							frontier.AppendEdgeAndVertex(queue, e, v, variable, e.Weight)
						}
					}
				}
			}
		}
		sort.Sort(frontier)
	}
	return false
}

func (t *Plan) SearchPlan(iterator IteratorFrontier, predicates []interface{}) (iteratorFrontier IteratorFrontier, err error) {
	t.predicates = predicates
	t.Depth = len(predicates)

	results := make(chan *Frontier)

	t.forEach(iterator, results)

	go func() {
		t.wg.Wait()
		close(results)
	}()

	return func() (*Frontier, bool) {
		f, opened := <-results
		if opened {
			return f, true
		}
		return nil, false
	}, nil
}

func (t *Plan) worker(f *Frontier, results chan *Frontier) {
	if t.UniformCostSearch(f) {
		results <- f
		t.wg.Done()
	} else if f.Len() > 0 {
		t.worker(f, results)
	} else {
		t.wg.Done()
	}
}

func (t *Plan) forEach(i IteratorFrontier, results chan *Frontier) {
	for item, ok := i(); ok; item, ok = i() {
		t.wg.Add(1)
		go t.worker(item, results)
	}
}
