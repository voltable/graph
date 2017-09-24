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
			//part.Path
			//v := *frontier.Peek()
			// if where, is := v.(ast.NonTerminalExpr); is {
			// 	if where.Interpret(v) {
			// 		return v, true
			// 	}
			// }
			//			return v, ok

			return qe.filterFrontier(i, part.Path, frontier)
		}
		return nil, false
	}, nil
}

func (qe Engine) filterFrontier(i query.IteratorFrontier, path query.Path, f *query.Frontier) (*vertices.Vertex, bool) {

	edgePath := query.NewEdgePath(i, qe.Storage.Fetch())
	vertexPath := query.NewVertexPath(i, qe.Storage.Fetch())
	iterated := false
	var result interface{}
	for p := path.Next(); p != nil; p = p.Next() {

		if pv, ok := p.(*query.PredicateVertexPath); ok {

			edgePath = vertexPath.Node(pv.PredicateVertex)
			result, iterated = edgePath.Iterate()

		} else if pe, ok := p.(*query.PredicateEdgePath); ok {
			vertexPath = edgePath.Relationship(pe.PredicateEdge)
			result, iterated = vertexPath.Iterate()
		}
		if iterated {
			if _, is := result.(*query.Frontier); is {
				return nil, true
			}
		}
	}
	return nil, false
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
