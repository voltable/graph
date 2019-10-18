package cypher

import (
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
		Parser:       parser.NewParser(),
		Storage:      i,
		Parts:        NewParts(),
		QueryBuilder: NewQueryBuilder(i),
	}
}

// QueryEngine is a implementation of the Query interface used to pass cypher queries
type QueryEngine struct {
	parser.Parser
	widecolumnstore.Storage
	Parts
	QueryBuilder
}

var _ query.QueryEngine = (*QueryEngine)(nil)

// Parse in a cypher query as a string and get back Query that is abstracted from the cypher AST
func (qe QueryEngine) Parse(q string) (*graph.Query, error) {

	// stmt, err := qe.Parser.Parse(q)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "Parse failed")
	// }

	// iterator, err := qe.QueryBuilder.Build(stmt)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "Query failed")
	// }

	query := &graph.Query{}
	// for kv, ok := iterator(); ok; kv, ok = iterator() {
	// 	query.Results = append(query.Results, kv)
	// }

	return query, nil
}
