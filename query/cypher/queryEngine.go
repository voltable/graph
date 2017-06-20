package cypher

import (
	"strings"

	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/query/cypher/parser"
)

func init() {
	query.RegisterQuery(QueryType, query.QueryRegistration{
		NewFunc: NewQueryEngine,
	})
}

const QueryType = "Cypher"

func NewQueryEngine() (query.Query, error) {
	return &QueryEngine{}, nil
}

type QueryEngine struct {
}

func (qe QueryEngine) Parser(q string) (*query.VertexPath, error) {
	_, err := parser.NewParser(strings.NewReader(q)).Parse()

	return nil, err
}
