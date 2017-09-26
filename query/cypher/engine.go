package cypher

import (
	"fmt"
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
		Traversal: query.NewTraversal(i),
		Storage:   i,
		Parts:     NewParts(),
	}
}

// Engine is a implementation of the Query interface used to pass cypher queries
type Engine struct {
	Parser    parser.Parser
	Traversal CypherTraversal
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
		f := qe.toFontier(forEach)
		f = qe.Traversal.Travers(f, part.Path)
		forEach, _ = qe.filter(f, part)
	}

	results := qe.toVertices(forEach)
	query := query.NewQuery(q, results)

	return query, nil

}

func (qe Engine) filter(i query.IteratorFrontier, part *QueryPart) (enumerables.Iterator, error) {
	return func() (interface{}, bool) {
		for frontier, ok := i(); ok; frontier, ok = i() {
			// We only need the first array of vertices from the frontier as the rest aren't the the optimal path
			// Need to get the variable on the vertex so I can run the AST over the array
			vertices, _, _ := frontier.Pop()
			for _, v := range vertices {
				fmt.Printf(v.Variable)
			}

			return nil, false
		}
		return nil, false
	}, nil
}

func (qe Engine) toVertices(i enumerables.Iterator) []interface{} {
	results := make([]interface{}, 0)
	for item, ok := i(); ok; item, ok = i() {
		results = append(results, item)
	}
	return results
}

func (qe Engine) toFontier(i enumerables.Iterator) query.IteratorFrontier {
	return func() (*query.Frontier, bool) {
		for item, ok := i(); ok; item, ok = i() {
			if v, is := item.(*vertices.Vertex); is {
				f := query.NewFrontier(v)
				return &f, true
			}
		}
		return nil, false
	}
}
