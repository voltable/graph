package cypher_test

import (
	"io"
	"testing"

	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/query/cypher"
	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
	"github.com/RossMerr/Caudex.Graph/query/cypher/parser"
	"github.com/RossMerr/Caudex.Graph/vertices"
)

func Test_Filter(t *testing.T) {
	se := &cypher.Engine{}

	q := &query.Query{Path: &cypher.Root{}}

	se.Filter(q)
}

type ParserMock struct{}

func (p *ParserMock) Parse(r io.Reader) (ast.Stmt, error) {
	return nil, nil
}

func NewPaserMock() parser.IParser {

	return &ParserMock{}
}

func Test_Parser(t *testing.T) {
	se := cypher.NewEngine()
	se.IParser = NewPaserMock()
	se.ToPredicateVertex = func(*ast.VertexPatn) query.PredicateVertex {
		return func(v *vertices.Vertex) bool {
			return false
		}
	}
	se.ToPredicateEdge = func(patn *ast.EdgePatn) query.PredicateEdge {
		return func(e *vertices.Edge) bool {
			return false
		}
	}

	se.Parser("str")
}

func Test_ToQueryPath(t *testing.T) {
	edgePatn := &ast.EdgePatn{Body: &ast.EdgeBodyStmt{LengthMinimum: 2, LengthMaximum: 5}}
	vertexPatn := &ast.VertexPatn{Variable: "bar", Edge: edgePatn}
	var b bool

	toPredicateVertex := func(*ast.VertexPatn) query.PredicateVertex {
		return func(v *vertices.Vertex) bool {
			return b
		}
	}

	toPredicateEdge := func(patn *ast.EdgePatn) query.PredicateEdge {
		return func(e *vertices.Edge) bool {
			return b
		}
	}

	want := &cypher.Root{}
	vertexPath := &query.PredicateVertexPath{PredicateVertex: toPredicateVertex(vertexPatn)}
	vertexPath.SetNext(&query.PredicateEdgePath{PredicateEdge: toPredicateEdge(edgePatn)})
	want.SetNext(vertexPath)

	engine := cypher.NewEngine()
	engine.ToPredicateEdge = toPredicateEdge
	engine.ToPredicateVertex = toPredicateVertex

	got, _ := engine.ToQueryPath(&ast.MatchStmt{Pattern: vertexPatn})

	v, _ := got.Next().(query.VertexNext)
	if v == nil {
		t.Errorf("VertexNext")
	}

	pv, _ := v.(*query.PredicateVertexPath)

	if pv == nil {
		t.Errorf("PredicateVertexPath")
	}

	e, _ := pv.Next().(query.EdgeNext)
	if e == nil {
		t.Errorf("EdgeNext")
	}

	pe, _ := e.(*query.PredicateEdgePath)
	if pe == nil {
		t.Errorf("PredicateEdgePath")
	}

	if pe.Length().Minimum != 2 {
		t.Errorf("Minimum")
	}

	if pe.Length().Maximum != 5 {
		t.Errorf("Maximum")
	}

	last, _ := pe.Next().(query.VertexNext)
	if last != nil {
		t.Errorf("VertexNext")
	}
}

func Test_IsPattern(t *testing.T) {

	if _, ok := cypher.IsPattern(&ast.DeleteStmt{}); !ok {
		t.Errorf("DeleteStmt")
	}

	if _, ok := cypher.IsPattern(&ast.CreateStmt{}); !ok {
		t.Errorf("CreateStmt")
	}

	if _, ok := cypher.IsPattern(&ast.OptionalMatchStmt{}); !ok {
		t.Errorf("OptionalMatchStmt")
	}

	if _, ok := cypher.IsPattern(&ast.MatchStmt{}); !ok {
		t.Errorf("MatchStmt")
	}

	if _, ok := cypher.IsPattern(&ast.WhereStmt{}); ok {
		t.Errorf("WhereStmt")
	}
}
