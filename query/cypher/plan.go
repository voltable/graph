package cypher

import (
	"sync"

	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
	"github.com/RossMerr/Caudex.Graph/query/traversal"
)

type Plan struct {
	wg         *sync.WaitGroup
	builder    QueryBuilder
	storage    query.Storage
	predicates []query.Predicate
}

func NewPlan(builder QueryBuilder, storage query.Storage) *Plan {

	plan := &Plan{
		wg:      &sync.WaitGroup{},
		storage: storage,
		builder: builder,
	}
	return plan
}

func (t *Plan) SearchPlan(iterator query.IteratorFrontier, patterns []ast.Patn) (iteratorFrontier query.IteratorFrontier, err error) {
	t.predicates = t.builder.Predicate(patterns)
	results := make(chan *query.Frontier)

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
	if traversal.UniformCostSearch(t.storage, t.predicates, f) {
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
