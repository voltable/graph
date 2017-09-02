package cypher

import (
	"strings"

	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
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

func newEngine() (query.Engine, error) {
	return &Engine{}, nil
}

// Engine is a implementation of the Query interface used to pass cypher queries
type Engine struct {
}

// Parser in a cypher query as a string and get back Query that is abstracted from the cypher AST
func (qe Engine) Parser(q string) (query.QueryInternal, error) {
	stmt, err := parser.NewParser(strings.NewReader(q)).Parse()
	if err != nil {
		return nil, err
	}
	path, err := ToQueryPath(stmt, ast.ToPredicateVertex, ast.ToPredicateEdge)
	if err != nil {
		return nil, err
	}

	query := NewQuery(stmt, path)

	return query, nil
}
