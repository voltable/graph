package ast_test

import (
	"testing"

	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
)

func Test_NotPrecedence(t *testing.T) {
	c := ast.NotExpr{}

	if ast.NotPrecedence(c) != 13 {
		t.Errorf("not expected")
	}
}
