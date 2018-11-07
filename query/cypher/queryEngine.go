package cypher

import (
	graph "github.com/RossMerr/Caudex.Graph"
	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/query/cypher/parser"
	"github.com/RossMerr/Caudex.Graph/widecolumnstore"
	"github.com/pkg/errors"
)

func init() {
	query.RegisterQueryEngine(queryType, query.QueryEngineRegistry{
		NewFunc: newEngine,
	})
}

// RegisterEngine forces the call of init
func RegisterEngine() {
	// Forces the call of init
}

const queryType = "cypher"

func newEngine(i widecolumnstore.Storage) (query.QueryEngine, error) {
	e := NewQueryEngine(i)
	return e, nil
}

// NewQueryEngine creates a new QueryEngine
func NewQueryEngine(i widecolumnstore.Storage) *QueryEngine {
	return &QueryEngine{
		Parser:  parser.NewParser(),
		Storage: i,
		Parts:   NewParts(),
		Builder: NewQueryBuilderDefault(i),
		//Filter:  NewFilter(),
	}
}

// QueryEngine is a implementation of the Query interface used to pass cypher queries
type QueryEngine struct {
	Parser parser.Parser
	//Filter  CypherFilter
	Storage widecolumnstore.Storage
	Parts   Parts
	//Projection Projection
	Builder *QueryBuilder
}

var _ query.QueryEngine = (*QueryEngine)(nil)

// Parse in a cypher query as a string and get back Query that is abstracted from the cypher AST
func (qe QueryEngine) Parse(q string) (*graph.Query, error) {
	stmt, err := qe.Parser.Parse(q)
	if err != nil {
		return nil, err
	}

	queryPart, err := qe.Parts.ToQueryPart(stmt)
	if err != nil {
		return nil, err
	}
	plan := NewPlan(qe.Builder)
	results := make([]interface{}, 0)
	for _, part := range queryPart {
		prefix := widecolumnstore.NewKey(query.TID, &widecolumnstore.Column{}).Marshal()
		frontier := qe.toFrontier(qe.Storage.HasPrefix(prefix), part.Variable())
		f, err := plan.SearchPlan(frontier, part.Patterns)
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
		id := query.UUID(&kv)
		if ok {
			f := query.NewFrontier(id, variable)
			return &f, true
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
