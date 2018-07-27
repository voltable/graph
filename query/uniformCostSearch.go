package query

import (
	"sort"
	"sync"

	"github.com/RossMerr/Caudex.Graph/keyvalue"
)

type Traversal interface {
	SearchPlan(iterator IteratorFrontier, predicates []*PredicatePath) (iteratorFrontier IteratorFrontier, err error)
}

type Plan struct {
	wg         *sync.WaitGroup
	predicates []*PredicatePath
	Depth      int
	storage    keyvalue.Storage
}

func NewPlan() *Plan {
	plan := &Plan{
		wg: &sync.WaitGroup{},
	}
	return plan
}

func (t *Plan) uniformCostSearch(frontier *Frontier) bool {
	if frontier.Len() > 0 {
		queue := frontier.Pop()
		depth := len(queue.Parts)
		part := queue.Parts[depth-1]

		if _, ok := frontier.Explored[part.UUID]; !ok {
			frontier.Explored[part.UUID] = true
			pv := t.predicates[depth-1]
			if variable, p := pv.Predicate(part.KeyValue, depth-1); p == Matched {
				queue.Parts[depth-1].Variable = variable
				frontier.AppendKeyValue(queue, part.KeyValue, part.Variable)
				sort.Sort(frontier)
				return depth == t.Depth
			}

		}

		if pe := t.predicates[depth]; pe != nil {
			iterator := t.storage.HasPrefix(keyvalue.RelationshipPrefix(part.UUID))
			for kv, hasEdges := iterator(); hasEdges; kv, hasEdges = iterator() {
				if _, ok := frontier.Explored[kv.UUID()]; !ok {
					// if variable, p := pe.Predicate(kv, depth); p == Visiting || p == Matching {
					// 	if v, err := t.storage.Fetch(e.ID()); err == nil {
					// 		frontier.AppendEdgeAndVertex(queue, e, v, variable, e.Weight)
					// 	}
					// }
				}
			}
		}

		// 		if pe := t.predicateEdge(depth); pe != nil {
		// 			for _, e := range vertex.Edges() {
		// 				if _, ok := frontier.Explored[e.ID()]; !ok {
		// 					if variable, p := pe(e, uint(depth)); p == Visiting || p == Matching {
		// 						if v, err := t.storage.Fetch(e.ID()); err == nil {
		// 							frontier.AppendEdgeAndVertex(queue, e, v, variable, e.Weight)
		// 						}
		// 					}
		// 				}
		// 			}
		// }

		sort.Sort(frontier)
	}
	return false
}

func (t *Plan) SearchPlan(iterator IteratorFrontier, predicates []*PredicatePath) (iteratorFrontier IteratorFrontier, err error) {
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
	if t.uniformCostSearch(f) {
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
