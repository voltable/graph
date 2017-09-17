package cypher

import (
	"strings"

	"github.com/RossMerr/Caudex.Graph/enumerables"
	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
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
	e := &Engine{
		Parser:            parser.NewParser(),
		ToPredicateVertex: ast.ToPredicateVertex,
		ToPredicateEdge:   ast.ToPredicateEdge,
		Traversal:         query.NewTraversal(i),
		Storage:           i,
	}

	e.ToPart = e.ToQueryPart
	return e
}

// Engine is a implementation of the Query interface used to pass cypher queries
type Engine struct {
	Parser            parser.Parser
	ToPredicateVertex func(*ast.VertexPatn) query.PredicateVertex
	ToPredicateEdge   func(patn *ast.EdgePatn) query.PredicateEdge
	ToPart            func(stmt ast.Stmt) ([]QueryPart, error)
	Traversal         CypherTraversal
	Storage           storage.Storage
}

var _ query.Engine = (*Engine)(nil)

// Parse in a cypher query as a string and get back Query that is abstracted from the cypher AST
func (qe Engine) Parse(q string) (*query.Query, error) {
	stmt, err := qe.Parser.Parse(strings.NewReader(q))
	if err != nil {
		return nil, err
	}

	queryPart, err := qe.ToPart(stmt)
	if err != nil {
		return nil, err
	}

	forEach := qe.Storage.ForEach()
	for _, part := range queryPart {
		f := qe.toFontier(forEach)
		f = qe.Traversal.Travers(f, part.Path)
		forEach = qe.Filter(f, part.Where)
	}

	results := qe.toVertices(forEach)
	query := query.NewQuery(q, results)

	return query, nil

}

func (qe Engine) Filter(i func() query.IteratorFrontier, stmt ast.Stmt) func() enumerables.Iterator {
	return func() enumerables.Iterator {
		return func() (interface{}, bool) {
			next := i()
			for frontier, ok := next(); ok; frontier, ok = next() {
				return *frontier.Peek(), ok
			}
			return nil, false
		}
	}
}

func (qe Engine) toVertices(i func() enumerables.Iterator) []interface{} {
	next := i()
	results := make([]interface{}, 0)
	for item, ok := next(); ok; item, ok = next() {
		if frontier, is := item.(query.Frontier); is {
			results = append(results, *frontier.Peek())
		}
	}
	return results
}

func (qe Engine) toFontier(i func() enumerables.Iterator) func() query.IteratorFrontier {

	return func() query.IteratorFrontier {
		return func() (*query.Frontier, bool) {
			next := i()
			for item, ok := next(); ok; item, ok = next() {
				if v, is := item.(vertices.Vertex); is {
					f := query.NewFrontier(&v)
					return &f, true
				}
			}
			return nil, false
		}
	}
}

// QueryPart is one part of a explicitly separate query parts
type QueryPart struct {
	Path  query.Path
	Where ast.Stmt
}

// ToQueryPath converts a cypher.Stmt to a QueryPath the queryPath is used to walk the graph
func (qe Engine) ToQueryPart(stmt ast.Stmt) ([]QueryPart, error) {

	arr := make([]QueryPart, 0)
	q, _ := NewPath()
	qp := QueryPart{Path: q}
	arr = append(arr, qp)
	var next func(query.Path)
	next = q.SetNext
	if pattern, ok := IsPattern(stmt); ok {
		for pattern != nil {
			if v, ok := pattern.(*ast.VertexPatn); ok && v != nil {
				pvp := query.PredicateVertexPath{PredicateVertex: qe.ToPredicateVertex(v)}
				next(&pvp)
				next = pvp.SetNext
				pattern = v.Edge

			} else if e, ok := pattern.(*ast.EdgePatn); ok && e != nil {
				pvp := query.PredicateEdgePath{PredicateEdge: qe.ToPredicateEdge(e)}
				if e.Body != nil {
					pvp.SetLength(e.Body.LengthMinimum, e.Body.LengthMaximum)
				}
				next(&pvp)
				next = pvp.SetNext
				pattern = e.Vertex
				// don't like making the WhereStmt a pattern
			} else if w, ok := pattern.(*ast.WhereStmt); ok && w != nil {
				//todo this might not be right
				qp.Where = w
				break
			} else {
				break
			}
		}
	}
	return arr, nil
}

func IsPattern(item ast.Stmt) (ast.Patn, bool) {
	if b, ok := item.(*ast.DeleteStmt); ok {
		return b.Pattern, true
	} else if b, ok := item.(*ast.CreateStmt); ok {
		return b.Pattern, true
	} else if b, ok := item.(*ast.OptionalMatchStmt); ok {
		return b.Pattern, true
	} else if b, ok := item.(*ast.MatchStmt); ok {
		return b.Pattern, true
	}
	return nil, false
}
