package cypher_test

import (
	"fmt"
	"io"
	"reflect"
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

type FakeParser struct {
	err error
}

func (p *FakeParser) Parse(r io.Reader) (ast.Stmt, error) {
	return nil, p.err
}

func NewFakePaser(err error) parser.Parser {

	return &FakeParser{err: err}
}

func Test_Parser(t *testing.T) {

	toPredicateVertex := func(*ast.VertexPatn) query.PredicateVertex {
		return func(v *vertices.Vertex) bool {
			return false
		}
	}
	toPredicateEdge := func(patn *ast.EdgePatn) query.PredicateEdge {
		return func(e *vertices.Edge) bool {
			return false
		}
	}

	tests := []struct {
		e               *cypher.Engine
		p               parser.Parser
		predicateVertex func(*ast.VertexPatn) query.PredicateVertex
		predicateEdge   func(patn *ast.EdgePatn) query.PredicateEdge
		path            func(stmt ast.Stmt) (query.Path, error)
		s               string
		err             string
	}{
		{
			e:               cypher.NewEngine(),
			p:               NewFakePaser(nil),
			predicateVertex: toPredicateVertex,
			predicateEdge:   toPredicateEdge,
			s:               "str",
		},
		{
			e:               cypher.NewEngine(),
			p:               NewFakePaser(fmt.Errorf("paser error")),
			predicateVertex: toPredicateVertex,
			predicateEdge:   toPredicateEdge,
			s:               "str",
			err:             "paser error",
		},
		{
			e:               cypher.NewEngine(),
			p:               NewFakePaser(nil),
			predicateVertex: toPredicateVertex,
			predicateEdge:   toPredicateEdge,
			s:               "str",
			path: func(stmt ast.Stmt) (query.Path, error) {
				return nil, fmt.Errorf("path error")
			},
			err: "path error",
		},
	}

	for i, tt := range tests {
		tt.e.Parser = tt.p
		tt.e.ToPredicateEdge = tt.predicateEdge
		tt.e.ToPredicateVertex = tt.predicateVertex

		if tt.path != nil {
			tt.e.ToPath = tt.path
		}
		_, err := tt.e.Parse(tt.s)
		if !reflect.DeepEqual(tt.err, errstring(err)) {
			t.Errorf("%d. %q: error mismatch:\n  exp=%s\n  got=%s\n\n", i, tt.s, tt.err, err)
		}
	}
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

// errstring returns the string representation of an error.
func errstring(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}
