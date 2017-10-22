package cypher

import (
	"strings"

	"github.com/RossMerr/Caudex.Graph/enumerables"
	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/query/cypher/parser"
	"github.com/RossMerr/Caudex.Graph/storage"
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
		Traversal: query.NewTraversal(i),
		Storage:   i,
		Parts:     NewParts(),
		Filter:    NewFilter(),
	}
}

// Engine is a implementation of the Query interface used to pass cypher queries
type Engine struct {
	Parser    parser.Parser
	Traversal CypherTraversal
	Filter    CypherFilter
	Storage   storage.Storage
	Parts     Parts
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

	forEach := qe.Storage.ForEach()
	for _, part := range queryPart {
		f, err := qe.Traversal.Travers(forEach, part.Path)
		if err != nil {
			return nil, err
		}
		forEach = qe.Filter.Filter(f, part.Predicate())
	}

	results := qe.toVertices(forEach)
	query := query.NewQuery(q, results)

	return query, nil

}

func (qe Engine) toVertices(i enumerables.Iterator) []interface{} {
	results := make([]interface{}, 0)
	for item, ok := i(); ok; item, ok = i() {
		results = append(results, item)
	}
	return results
}
