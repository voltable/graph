package cypher

import (
	"strings"

	"github.com/RossMerr/Caudex.Graph/enumerables"
	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/query/cypher/parser"
	"github.com/RossMerr/Caudex.Graph/storage"
	"github.com/RossMerr/Caudex.Graph/vertices"
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

func newEngine(i storage.Storage) (query.Engine, error) {
	e := NewEngine(i)
	return e, nil
}

func NewEngine(i storage.Storage) *Engine {
	return &Engine{
		Parser:    parser.NewParser(),
		Traversal: query.NewPlan(i),
		Storage:   i,
		Parts:     NewParts(),
		Filter:    NewFilter(),
	}
}

// Engine is a implementation of the Query interface used to pass cypher queries
type Engine struct {
	Parser     parser.Parser
	Traversal  query.Traversal
	Filter     CypherFilter
	Storage    storage.Storage
	Parts      Parts
	Projection Projection
}

var _ query.Engine = (*Engine)(nil)

// Parse in a cypher query as a string and get back Query that is abstracted from the cypher AST
func (qe Engine) Parse(q string) (*query.Query, error) {
	stmt, err := qe.Parser.Parse(strings.NewReader(q))
	if err != nil {
		return nil, err
	}

	queryPart, err := qe.Parts.ToQueryPart(stmt)
	if err != nil {
		return nil, err
	}

	results := make([]interface{}, 0)
	for _, part := range queryPart {
		frontier := qe.toFrontier(qe.Storage.ForEach(), variableVertex(queryPart))
		f, err := qe.Traversal.SearchPlan(frontier, part.Predicates)
		if err != nil {
			return nil, err
		}
		f = qe.Filter.Filter(f, part.Predicate())
		results = append(results, qe.Projection.Transform(f, part.Maps())...)
	}

	query := query.NewQuery(q, results)

	return query, nil

}

func variableVertex(queryPart []*QueryPart) string {
	if len(queryPart) > 0 {
		if len(queryPart[0].Predicates) > 0 {
			e := queryPart[0].Predicates[0]
			if pv, ok := e.(*query.PredicateVertexPath); ok {
				return pv.Variable
			}
		}
	}
	return ""
}

func (qe Engine) toVertices(i query.IteratorFrontier) []interface{} {
	results := make([]interface{}, 0)
	for item, ok := i(); ok; item, ok = i() {
		for _, i := range item.OptimalPath() {
			results = append(results, i.Object)
		}

	}
	return results
}

func (qe Engine) toFrontier(i enumerables.Iterator, variable string) query.IteratorFrontier {
	return func() (*query.Frontier, bool) {
		item, ok := i()
		if ok {
			if v, is := item.(*vertices.Vertex); is {
				f := query.NewFrontier(v, variable)
				return &f, true
			}
		}

		return nil, false
	}
}
