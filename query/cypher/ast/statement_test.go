package ast_test

import (
	"testing"

	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
)

func Test_Precedence(t *testing.T) {

	boolean := &ast.BooleanExpr{}
	if ast.Precedence(boolean) != 20 {
		t.Errorf("boolean expected")
	}
}
