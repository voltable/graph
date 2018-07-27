package cypher

import (
	"github.com/RossMerr/Caudex.Graph"
	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/query/cypher/parser"
)

func init() {
	query.RegisterEngine(queryType, query.EngineRegistration{
		NewFunc: newEngine,
	})
}

// RegisterEngine forces the call of init
func RegisterEngine() {
	// Forces the call of init
}

const queryType = "Cypher"

func newEngine(i query.Storage) (query.Engine, error) {
	e := NewEngine(i)
	return e, nil
}

func NewEngine(i query.Storage) *Engine {
	builder, _ := FristQueryBuilder(i)
	return &Engine{
		Parser:  parser.NewParser(),
		Storage: i,
		Parts:   NewParts(),
		Builder: builder,
		//Filter:  NewFilter(),
	}
}

// Engine is a implementation of the Query interface used to pass cypher queries
type Engine struct {
	Parser parser.Parser
	//Filter  CypherFilter
	Storage query.Storage
	Parts   Parts
	//Projection Projection
	Builder QueryBuilder
}

var _ query.Engine = (*Engine)(nil)

// Parse in a cypher query as a string and get back Query that is abstracted from the cypher AST
func (qe Engine) Parse(q string) (*graph.Query, error) {
	stmt, err := qe.Parser.Parse(q)
	if err != nil {
		return nil, err
	}

	queryPart, err := qe.Parts.ToQueryPart(stmt)
	if err != nil {
		return nil, err
	}
	plan := NewPlan(qe.Builder, qe.Storage)
	results := make([]interface{}, 0)
	for _, part := range queryPart {
		frontier := qe.toFrontier(qe.Storage.ForEach(), part, variableVertex(queryPart))
		f, err := plan.SearchPlan(frontier, part.Patterns)
		if err != nil {
			return nil, err
		}
		results = append(results, Transform(f)...)
		//f = qe.Filter.Filter(f, part.Predicate())
		//results = append(results, qe.Projection.Transform(f, part.Maps())...)
	}

	query := graph.NewQuery(q, results)

	return query, nil

}

func variableVertex(queryPart []*QueryPart) string {
	if len(queryPart) > 0 {
		if len(queryPart[0].Patterns) > 0 {
			e := queryPart[0].Patterns[0]
			return e.V()
		}
	}
	return ""
}

func (qe Engine) toVertices(i query.IteratorFrontier) []interface{} {
	results := make([]interface{}, 0)
	for item, ok := i(); ok; item, ok = i() {
		for _, i := range item.OptimalPath() {
			results = append(results, i.UUID)
		}

	}
	return results
}

func (qe Engine) toFrontier(i query.IteratorUUID, part *QueryPart, variable string) query.IteratorFrontier {
	return func() (*query.Frontier, bool) {
		id, ok := i()
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
