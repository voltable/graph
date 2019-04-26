package cypher

import (
	"sync"

	"github.com/voltable/graph/query"
	"github.com/voltable/graph/query/cypher/traversal"
	"github.com/voltable/graph/widecolumnstore"
)

type Plan struct {
	wg       *sync.WaitGroup
	operator widecolumnstore.Operator
	engine   query.Graph
}

func NewPlan() *Plan {

	plan := &Plan{
		wg: &sync.WaitGroup{},
	}
	return plan
}

func (t *Plan) SearchPlan(iterator query.IteratorFrontier, operator widecolumnstore.Operator) (query.IteratorFrontier, error) {
	results := make(chan *query.Frontier)
	t.operator = operator

	t.forEach(iterator, results)

	go func() {
		t.wg.Wait()
		close(results)
	}()

	return func() (*query.Frontier, bool) {
		f, opened := <-results
		if opened {
			return f, true
		}
		return nil, false
	}, nil
}

func (t *Plan) worker(f *query.Frontier, results chan *query.Frontier) {
	if traversal.UniformCostSearch(t.engine, t.operator, f) {
		results <- f
		t.wg.Done()
	} else if f.Len() > 0 {
		t.worker(f, results)
	} else {
		t.wg.Done()
	}
}

func (t *Plan) forEach(i query.IteratorFrontier, results chan *query.Frontier) {
	for item, ok := i(); ok; item, ok = i() {
		t.wg.Add(1)
		go t.worker(item, results)
	}
}
