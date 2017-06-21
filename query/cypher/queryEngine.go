package cypher

import (
	"strings"

	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/query/cypher/parser"
)

func init() {
	query.RegisterQuery(queryType, query.QueryRegistration{
		NewFunc: newQueryEngine,
	})
}

const queryType = "Cypher"

func newQueryEngine() (query.Query, error) {
	return &QueryEngine{}, nil
}

// QueryEngine is a implementation of the Query interface used to pass cypher queries
type QueryEngine struct {
}

// Parser in a cypher query as a string and get back VertexPath that is abstracted from the cypher AST
func (qe QueryEngine) Parser(q string) (*query.VertexPath, error) {
	_, err := parser.NewParser(strings.NewReader(q)).Parse()

	return nil, err
}
