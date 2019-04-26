package cypher

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	graph "github.com/voltable/graph"
	"github.com/voltable/graph/query"
	"github.com/voltable/graph/query/cypher/parser"
	"github.com/voltable/graph/widecolumnstore"
)

func init() {
	query.RegisterQueryEngine(QueryType, query.QueryEngineRegistry{
		NewFunc: newEngine,
	})
}

// RegisterEngine forces the call of init
func RegisterEngine() {
	// Forces the call of init
}

const QueryType = graph.QueryType("cypher")

func newEngine(i widecolumnstore.Storage) (query.QueryEngine, error) {
	e := NewQueryEngine(i)
	return e, nil
}

// NewQueryEngine creates a new QueryEngine
func NewQueryEngine(i widecolumnstore.Storage) *QueryEngine {
	return &QueryEngine{
		Parser:             parser.NewParser(),
		Storage:            i,
		Parts:              NewParts(),
		CypherQueryBuilder: NewQueryBuilderDefault(i),
		//Filter:  NewFilter(),
	}
}

// QueryEngine is a implementation of the Query interface used to pass cypher queries
type QueryEngine struct {
	parser.Parser
	//Filter  CypherFilter
	widecolumnstore.Storage
	Parts
	//Projection Projection
	*CypherQueryBuilder
}

var _ query.QueryEngine = (*QueryEngine)(nil)

// Parse in a cypher query as a string and get back Query that is abstracted from the cypher AST
func (qe QueryEngine) Parse(q string) (*graph.Query, error) {
	stmt, err := qe.Parser.Parse(q)
	if err != nil {
		return nil, errors.Wrap(err, "Parse failed")
	}

	queryPart, err := qe.ToQueryPart(stmt)
	if err != nil {
		return nil, errors.Wrap(err, "ToQueryPart failed")
	}

	plan := NewPlan()
	results := make([]interface{}, 0)
	for _, part := range queryPart {
		frontier := qe.toFrontier(qe.HasPrefix(query.AllNodesKey), part.Variable())

		operator, err := qe.Predicate(part.Patterns)
		if err != nil {
			return nil, errors.Wrap(err, "Predicate failed")
		}

		f, err := plan.SearchPlan(frontier, operator)
		if err != nil {
			return nil, errors.Wrap(err, "QueryEngine Parse")
		}

		results = append(results, Transform(f)...)
		//	f = qe.Filter.Filter(f, part.Predicate())
		//results = append(results, qe.Projection.Transform(f, part.Maps())...)
	}

	query := graph.NewQuery(q, results)

	return query, nil

}

func (qe QueryEngine) toVertices(i query.IteratorFrontier) []interface{} {
	results := make([]interface{}, 0)
	for item, ok := i(); ok; item, ok = i() {
		for _, i := range item.OptimalPath() {
			results = append(results, i.UUID)
		}

	}
	return results
}

func (qe QueryEngine) toFrontier(i widecolumnstore.Iterator, variable string) query.IteratorFrontier {
	return func() (*query.Frontier, bool) {
		kv, ok := i()
		id, err := query.UUID(&kv)
		if err != nil {
			log.Error(errors.Wrap(err, "QueryEngine toFrontier"))
		} else {
			if ok {
				f := query.NewFrontier(&id, variable)
				return &f, true
			}
		}

		return nil, false
	}
}

func Transform(i query.IteratorFrontier) []interface{} {
	results := make([]interface{}, 0)
	for item, ok := i(); ok; item, ok = i() {
		for _, part := range item.OptimalPath() {
			results = append(results, part.UUID)
		}

	}
	return results
}
