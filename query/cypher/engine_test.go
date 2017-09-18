package cypher_test

import (
	"fmt"
	"io"
	"reflect"
	"testing"

	"github.com/RossMerr/Caudex.Graph/enumerables"
	"github.com/RossMerr/Caudex.Graph/query/cypher"
	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
	"github.com/RossMerr/Caudex.Graph/query/cypher/parser"
	"github.com/RossMerr/Caudex.Graph/storage"
	"github.com/RossMerr/Caudex.Graph/vertices"
)

type FakeParser struct {
	err error
}

func (p *FakeParser) Parse(r io.Reader) (ast.Stmt, error) {
	return nil, p.err
}

func NewFakePaser(err error) parser.Parser {

	return &FakeParser{err: err}
}

type FakeTraversal struct {
}

type FakeStorage struct {
}

func (s FakeStorage) Fetch() func(string) (*vertices.Vertex, error) {
	return nil
}

func (s FakeStorage) ForEach() enumerables.Iterator {
	return func() (item interface{}, ok bool) {

		return nil, false
	}
}

func NewFakeStorage() storage.Storage {
	return &FakeStorage{}
}

type FakeParts struct {
}

func (s FakeParts) ToQueryPart(stmt ast.Stmt) ([]*cypher.QueryPart, error) {
	return nil, nil
}

func NewFakeParts() cypher.Parts {
	return &FakeParts{}
}

func Test_Parser(t *testing.T) {

	tests := []struct {
		e     *cypher.Engine
		p     parser.Parser
		parts cypher.Parts
		path  func(stmt ast.Stmt) ([]cypher.QueryPart, error)
		s     string
		err   string
	}{
		{
			e:     cypher.NewEngine(NewFakeStorage()),
			p:     NewFakePaser(nil),
			parts: NewFakeParts(),
			s:     "str",
		},
		{
			e:     cypher.NewEngine(NewFakeStorage()),
			p:     NewFakePaser(fmt.Errorf("paser error")),
			parts: NewFakeParts(),
			s:     "str",
			err:   "paser error",
		},
	}

	for i, tt := range tests {
		tt.e.Parser = tt.p
		tt.e.Parts = tt.parts

		_, err := tt.e.Parse(tt.s)
		if !reflect.DeepEqual(tt.err, errstring(err)) {
			t.Errorf("%d. %q: error mismatch:\n  exp=%s\n  got=%s\n\n", i, tt.s, tt.err, err)
		}
	}
}

func Test_IsPattern(t *testing.T) {

	var tests = []struct {
		c      ast.Stmt
		result bool
	}{
		{
			c:      ast.DeleteStmt{},
			result: true,
		}, {
			c:      ast.CreateStmt{},
			result: true,
		}, {
			c:      ast.OptionalMatchStmt{},
			result: true,
		}, {
			c:      ast.MatchStmt{},
			result: true,
		}, {
			c:      ast.WhereStmt{},
			result: true,
		},
	}

	for i, tt := range tests {
		_, ok := cypher.IsPattern(&tt.c)
		if ok == tt.result {
			t.Errorf("%d. comparison mismatch:\n %v\n\n", i, tt.c)
		}
	}
}

// errstring returns the string representation of an error.
func errstring(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}
