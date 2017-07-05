package ast_test

import (
	"testing"

	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
)

func Test_MathPrecedence(t *testing.T) {

	add := &ast.AddExpr{}
	if ast.MathPrecedence(add) != 2 {
		t.Errorf("add expected")
	}

	subtract := &ast.SubtractExpr{}
	if ast.MathPrecedence(subtract) != 2 {
		t.Errorf("subtract expected")
	}

	divide := &ast.DivideExpr{}
	if ast.MathPrecedence(divide) != 3 {
		t.Errorf("divide expected")
	}

	multiply := &ast.MultiplyExpr{}
	if ast.MathPrecedence(multiply) != 3 {
		t.Errorf("multiply expected")
	}

	modulo := &ast.ModuloExpr{}
	if ast.MathPrecedence(modulo) != 4 {
		t.Errorf("modulo expected")
	}

	power := &ast.PowerExpr{}
	if ast.MathPrecedence(power) != 4 {
		t.Errorf("power expected")
	}
}
