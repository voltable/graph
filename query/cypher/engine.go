package cypher

import (
	"fmt"
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
	e := NewEngine()
	return e, nil
}

func NewEngine() *Engine {
	e := &Engine{
		Parser:            parser.NewParser(),
		ToPredicateVertex: ast.ToPredicateVertex,
		ToPredicateEdge:   ast.ToPredicateEdge,
	}

	e.ToPath = e.ToQueryPath
	return e
}

// Engine is a implementation of the Query interface used to pass cypher queries
type Engine struct {
	Parser            parser.Parser
	ToPredicateVertex func(*ast.VertexPatn) query.PredicateVertex
	ToPredicateEdge   func(patn *ast.EdgePatn) query.PredicateEdge
	ToPath            func(stmt ast.Stmt) (query.Path, error)
}

var _ query.Engine = (*Engine)(nil)

// Parse in a cypher query as a string and get back Query that is abstracted from the cypher AST
func (qe Engine) Parse(q string) (query.Path, error) {
	stmt, err := qe.Parser.Parse(strings.NewReader(q))
	if err != nil {
		return nil, err
	}
	path, err := qe.ToPath(stmt)
	if err != nil {
		return nil, err
	}

	return path, nil
}

// Filter is used to run any final part of the AST on the result set
func (qe Engine) Filter(q *query.Query) error {
	if root, ok := q.Path.(*Root); ok {
		fmt.Printf("%v", root)
	}

	return nil
}

// ToQueryPath converts a cypher.Stmt to a QueryPath the queryPath is used to walk the graph
func (qe Engine) ToQueryPath(stmt ast.Stmt) (query.Path, error) {
	q, _ := NewPath(stmt)
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
			} else {
				break
			}
		}
	}

	return q, nil
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
