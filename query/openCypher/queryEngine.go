package openCypher

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	graph "github.com/voltable/graph"
	"github.com/voltable/graph/query"
	"github.com/voltable/graph/query/openCypher/parser"
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

const QueryType = graph.QueryType("openCypher")

func newEngine(i widecolumnstore.Storage) (query.QueryEngine, error) {
	e := NewQueryEngine(i)
	return e, nil
}

// NewQueryEngine creates a new QueryEngine
func NewQueryEngine(i widecolumnstore.Storage) *QueryEngine {
	return &QueryEngine{
		storage:i,
	}
}

// QueryEngine is a implementation of the Query interface used to pass cypher queries
type QueryEngine struct {
	storage widecolumnstore.Storage
}

var _ query.QueryEngine = (*QueryEngine)(nil)

// Parse in a cypher query as a string and get back Query that is abstracted from the cypher AST
func (qe QueryEngine) Parse(q string) (*graph.Query, error) {
	// Setup the input
	is := antlr.NewInputStream(q)

	// Create the Lexer
	lexer := parser.NewCypherLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewCypherParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	tree := p.OC_Cypher()

	statistics := graph.NewStatistics()

	// Finally parse the expression (by walking the tree)
	walker := newCypherWalker()
	antlr.ParseTreeWalkerDefault.Walk(&walker, tree)

	ir := walker.getIR()
	planner := NewQueryPlanner(qe.storage, statistics, ir)
	results := planner.Execute()

	return graph.NewQuery(q, results, statistics), nil
}
