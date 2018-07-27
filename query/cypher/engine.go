package cypher

import (
	"github.com/RossMerr/Caudex.Graph"
	"github.com/RossMerr/Caudex.Graph/keyvalue"
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

func newEngine(i keyvalue.Storage) (query.Engine, error) {
	e := NewEngine(i)
	return e, nil
}

func NewEngine(i keyvalue.Storage) *Engine {
	return &Engine{
		Parser:  parser.NewParser(),
		Storage: i,
		Parts:   NewParts(),
		Filter:  NewFilter(),
	}
}

// Engine is a implementation of the Query interface used to pass cypher queries
type Engine struct {
	Parser     parser.Parser
	Filter     CypherFilter
	Storage    keyvalue.Storage
	Parts      Parts
	Projection Projection
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
	plan := query.NewPlan()
	results := make([]interface{}, 0)
	for _, part := range queryPart {
		// todo need to build up the hasprefix
		// qe.Storage.HasPrefix([]byte(""))
		frontier := qe.toFrontier(qe.Storage.ForEach(), part, variableVertex(queryPart))
		f, err := plan.SearchPlan(frontier, part.Predicates)
		if err != nil {
			return nil, err
		}
		f = qe.Filter.Filter(f, part.Predicate())
		results = append(results, qe.Projection.Transform(f, part.Maps())...)
	}

	query := graph.NewQuery(q, results)

	return query, nil

}

func variableVertex(queryPart []*QueryPart) string {
	if len(queryPart) > 0 {
		if len(queryPart[0].Predicates) > 0 {
			e := queryPart[0].Predicates[0]
			return e.Variable
		}
	}
	return ""
}

func (qe Engine) toVertices(i query.IteratorFrontier) []interface{} {
	results := make([]interface{}, 0)
	for item, ok := i(); ok; item, ok = i() {
		for _, i := range item.OptimalPath() {
			results = append(results, i.KeyValue)
		}

	}
	return results
}

func (qe Engine) toFrontier(i keyvalue.Iterator, part *QueryPart, variable string) query.IteratorFrontier {
	return func() (*query.Frontier, bool) {
		kv, ok := i()
		if ok {
			f := query.NewFrontier(kv, variable)
			return &f, true
		}

		return nil, false
	}
}
