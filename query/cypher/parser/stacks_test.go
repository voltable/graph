package parser_test

import (
	"testing"

	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
	"github.com/RossMerr/Caudex.Graph/query/cypher/parser"
)

// n.number >= 1 AND n.number <= 10
// should result in the AST tree :-
//					 AND
//			   _____/	\_____
//			  /				  \
//		  __>=__ 		  	 __<=__
//		 /	    \			/	   \
//  n.number     1  	n.number   10
func TestBasic_UpdateStack(t *testing.T) {
	exprStack := make(parser.StackExpr, 0)

	n1 := &ast.PropertyStmt{Variable: "n", Value: "number"}
	exprStack = exprStack.Push(n1)

	n2 := &ast.ComparisonExpr{Comparison: ast.GTE}
	exprStack = exprStack.Push(n2)

	n3 := &ast.Ident{Data: 1}
	exprStack = exprStack.Push(n3)

	n4 := &ast.BooleanExpr{Boolean: ast.AND}
	exprStack = exprStack.Push(n4)

	n5 := &ast.PropertyStmt{Variable: "n", Value: "number"}
	exprStack = exprStack.Push(n5)

	n6 := &ast.ComparisonExpr{Comparison: ast.NEQ}
	exprStack = exprStack.Push(n6)

	n7 := &ast.Ident{Data: 10}
	exprStack = exprStack.Push(n7)

	result, _ := exprStack.Shunt()

	if result != n4 {
		t.Errorf("found %s expected %s", result, n4)
	}

	if n4.X != n2 {
		t.Errorf("found %s expected %s", n4.X, n2)
	}

	if n2.X != n1 {
		t.Errorf("found %s expected %s", n2.X, n1)
	}

	if n2.Y != n3 {
		t.Errorf("found %s expected %s", n2.Y, n3)
	}

	if n4.Y != n6 {
		t.Errorf("found %s expected %s", n4.Y, n6)
	}

	if n6.X != n5 {
		t.Errorf("found %s expected %s", n6.X, n5)
	}

	if n6.Y != n7 {
		t.Errorf("found %s expected %s", n6.Y, n7)
	}
}

// n.number >= 1 AND NOT n.number = 10
// should result in the AST tree :-
//					 AND
//			   _____/	\____
//			  /				 \
//		  __>=__ 		  	 NOT-----(Y is always nil on a NOT)
//		 /	    \			  |
//  n.number     1  		__=__
//						   /	 \
//					   n.number   10
func TestBasic_UpdateStackWithNot(t *testing.T) {
	exprStack := make(parser.StackExpr, 0)

	n1 := &ast.PropertyStmt{Variable: "n", Value: "number"}
	exprStack = exprStack.Push(n1)

	n2 := &ast.ComparisonExpr{Comparison: ast.GTE}
	exprStack = exprStack.Push(n2)

	n3 := &ast.Ident{Data: 1}
	exprStack = exprStack.Push(n3)

	n4 := &ast.BooleanExpr{Boolean: ast.AND}
	exprStack = exprStack.Push(n4)

	n5 := &ast.BooleanExpr{Boolean: ast.NOT}
	exprStack = exprStack.Push(n5)

	n6 := &ast.PropertyStmt{Variable: "n", Value: "number"}
	exprStack = exprStack.Push(n6)

	n7 := &ast.ComparisonExpr{Comparison: ast.EQ}
	exprStack = exprStack.Push(n7)

	n8 := &ast.Ident{Data: 10}
	exprStack = exprStack.Push(n8)

	result, _ := exprStack.Shunt()

	if result != n4 {
		t.Errorf("found %s expected %s", result, n4)
	}

	if n4.X != n2 {
		t.Errorf("found %s expected %s", n4.X, n2)
	}

	if n2.X != n1 {
		t.Errorf("found %s expected %s", n2.X, n1)
	}

	if n2.Y != n3 {
		t.Errorf("found %s expected %s", n2.Y, n3)
	}

	if n4.Y != n5 {
		t.Errorf("found %s expected %s", n4.Y, n5)
	}

	if n5.X != n7 {
		t.Errorf("found %s expected %s", n5.X, n7)
	}

	if n5.Y != nil {
		t.Errorf("found %s expected nil", n5.Y)
	}

	if n7.X != n6 {
		t.Errorf("found %s expected %s", n7.X, n6)
	}

	if n7.Y != n8 {
		t.Errorf("found %s expected %s", n7.Y, n8)
	}
}

// n.name = 'Peter' XOR (n.age < 30 AND n.name = 'Tobias') OR NOT (n.name = 'Tobias' OR n.name = 'Peter')
// should result in the AST tree :-
//						     	  	   OR
//				   ___________________/  \____________________
//				  /						 	  		 		  \
//		 	_____XOR_____							 		  NOT-----(Y is always nil on a NOT)
//		   /	 	     \					   		  		   |
//		__=__		 	__AND__				 			 ______OR_____
//	   /   	 |	       /       \		    			/			  \
// n.name  'Peter'  __<__     __=__					 __=__		 	 __=__
//				   /	|	 |	   \				/	  \	    	/ 	  \
//				  /		|	 |	    \			n.name  'Tobias'  n.name  'Peter'
//

// 			n.age  30   n.name  'Tobias'
// func TestDeep_UpdateStack(t *testing.T) {
// 	exprStack := make(parser.StackExpr, 0)

// 	n1 := &ast.PropertyStmt{Variable: "n", Value: "name"}
// 	exprStack = exprStack.Push(n1)

// 	n2 := &ast.ComparisonExpr{Comparison: ast.EQ}
// 	exprStack = exprStack.Push(n2)

// 	n3 := &ast.Ident{Data: "Peter"}
// 	exprStack = exprStack.Push(n3)

// 	n4 := &ast.BooleanExpr{Boolean: ast.XOR}
// 	exprStack = exprStack.Push(n4)

// 	n5 := &ast.ParenthesesExpr{Parentheses: ast.LPAREN}
// 	exprStack = exprStack.Push(n5)

// 	n6 := &ast.PropertyStmt{Variable: "n", Value: "age"}
// 	exprStack = exprStack.Push(n6)

// 	n7 := &ast.ComparisonExpr{Comparison: ast.LT}
// 	exprStack = exprStack.Push(n7)

// 	n8 := &ast.Ident{Data: 30}
// 	exprStack = exprStack.Push(n8)

// 	n9 := &ast.BooleanExpr{Boolean: ast.AND}
// 	exprStack = exprStack.Push(n9)

// 	n10 := &ast.PropertyStmt{Variable: "n", Value: "name"}
// 	exprStack = exprStack.Push(n10)

// 	n11 := &ast.Ident{Data: "Tobias"}
// 	exprStack = exprStack.Push(n11)

// 	n12 := &ast.ParenthesesExpr{Parentheses: ast.RPAREN}
// 	exprStack = exprStack.Push(n12)

// 	n13 := &ast.BooleanExpr{Boolean: ast.OR}
// 	exprStack = exprStack.Push(n13)

// 	n14 := &ast.BooleanExpr{Boolean: ast.NOT}
// 	exprStack = exprStack.Push(n14)

// 	n15 := &ast.ParenthesesExpr{Parentheses: ast.LPAREN}
// 	exprStack = exprStack.Push(n15)

// 	n16 := &ast.PropertyStmt{Variable: "n", Value: "name"}
// 	exprStack = exprStack.Push(n16)

// 	n17 := &ast.ComparisonExpr{Comparison: ast.EQ}
// 	exprStack = exprStack.Push(n17)

// 	n18 := &ast.Ident{Data: "Tobias"}
// 	exprStack = exprStack.Push(n18)

// 	n19 := &ast.BooleanExpr{Boolean: ast.OR}
// 	exprStack = exprStack.Push(n19)

// 	n20 := &ast.PropertyStmt{Variable: "n", Value: "name"}
// 	exprStack = exprStack.Push(n20)

// 	n21 := &ast.ComparisonExpr{Comparison: ast.EQ}
// 	exprStack = exprStack.Push(n21)

// 	n22 := &ast.Ident{Data: "Peter"}
// 	exprStack = exprStack.Push(n22)

// 	n23 := &ast.ParenthesesExpr{Parentheses: ast.RPAREN}
// 	exprStack = exprStack.Push(n23)

// 	exprStack.Shunt()

// 	// n.name = 'Peter' XOR (n.age < 30 AND n.name = 'Tobias') OR NOT (n.name = 'Tobias' OR n.name = 'Peter')

// }
