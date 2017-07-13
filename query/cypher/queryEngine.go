package cypher

import (
	"strings"

	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/query/cypher/parser"
)

func init() {
	query.RegisterQueryEngine(queryType, query.QueryEngineRegistration{
		NewFunc: newQueryEngine,
	})
}

const queryType = "Cypher"

func newQueryEngine() (query.QueryEngine, error) {
	return &QueryEngine{}, nil
}

// QueryEngine is a implementation of the Query interface used to pass cypher queries
type QueryEngine struct {
}

// Parser in a cypher query as a string and get back Query that is abstracted from the cypher AST
func (qe QueryEngine) Parser(q string) (*query.QueryPath, error) {
	stmt, err := parser.NewParser(strings.NewReader(q)).Parse()
	if err != nil {
		return nil, err
	}
	return ToQueryPath(stmt)
}
