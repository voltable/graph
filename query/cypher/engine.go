package cypher

import (
	"strings"

	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
	"github.com/RossMerr/Caudex.Graph/query/cypher/parser"
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

func newEngine() (query.Engine, error) {
	e := NewEngine()
	return e, nil
}

func NewEngine() *Engine {
	e := &Engine{
		Parser:            parser.NewParser(),
		ToPredicateVertex: ast.ToPredicateVertex,
		ToPredicateEdge:   ast.ToPredicateEdge,
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
	Traversal         *query.Traversal
}

var _ query.Engine = (*Engine)(nil)

// Parse in a cypher query as a string and get back Query that is abstracted from the cypher AST
func (qe Engine) Parse(q string) (query.Path, error) {
	stmt, err := qe.Parser.Parse(strings.NewReader(q))
	if err != nil {
		return nil, err
	}
	path, err := qe.ToPart(stmt)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (qe Engine) Query(i func() query.Iterator, q string) (*query.Query, error) {
	stmt, err := qe.Parser.Parse(strings.NewReader(q))
	if err != nil {
		return nil, err
	}

	queryPart, err := qe.ToPart(stmt)

	for _, part := range queryPart {
		i = qe.Traversal.Travers(i, part.Path)
		i = qe.Filter(i, part.Where)
	}

	results := qe.toVertices(i)

	query := query.NewQuery(q, results)

	return query, nil

}

func (qe Engine) Filter(i func() query.Iterator, stmt ast.Stmt) func() query.Iterator {
	next := i()
	results := make([]interface{}, 0)
	for item, ok := next(); ok; item, ok = next() {
		if v, is := item.(*vertices.Vertex); is {
			results = append(results, v)
		}
	}
	return nil
}

func (qe Engine) toVertices(i func() query.Iterator) []interface{} {
	next := i()
	results := make([]interface{}, 0)
	for item, ok := next(); ok; item, ok = next() {
		if frontier, is := item.(query.Frontier); is {
			results = append(results, *frontier.Peek())
		}
	}
	return results
}

// QueryPart is one part of a explicitly separate query parts
type QueryPart struct {
	Path  query.Path
	Where ast.Stmt
}

// ToQueryPath converts a cypher.Stmt to a QueryPath the queryPath is used to walk the graph
func (qe Engine) ToQueryPart(stmt ast.Stmt) ([]QueryPart, error) {

	arr := make([]QueryPart, 1)
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
