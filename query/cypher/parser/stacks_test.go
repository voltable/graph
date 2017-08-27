package parser_test

import (
	"testing"

	"github.com/RossMerr/Caudex.Graph/expressions"
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

	n2 := &ast.ComparisonExpr{Comparison: expressions.GTE}
	exprStack = exprStack.Push(n2)

	n3 := &ast.Ident{Data: 1}
	exprStack = exprStack.Push(n3)

	n4 := &ast.BooleanExpr{Boolean: ast.AND}
	exprStack = exprStack.Push(n4)

	n5 := &ast.PropertyStmt{Variable: "n", Value: "number"}
	exprStack = exprStack.Push(n5)

	n6 := &ast.ComparisonExpr{Comparison: expressions.LTE}
	exprStack = exprStack.Push(n6)

	n7 := &ast.Ident{Data: 10}
	exprStack = exprStack.Push(n7)

	result, _ := exprStack.Shunt()

	if result != n4 {
		t.Errorf("found %#v expected %#v", result, n4)
	}

	if n4.Left != n2 {
		t.Errorf("found %#v expected %#v", n4.Left, n2)
	}

	if n2.Left != n1 {
		t.Errorf("found %#v expected %#v", n2.Left, n1)
	}

	if n2.Right != n3 {
		t.Errorf("found %#v expected %#v", n2.Right, n3)
	}

	if n4.Right != n6 {
		t.Errorf("found %#v expected %#v", n4.Right, n6)
	}

	if n6.Left != n5 {
		t.Errorf("found %#v expected %#v", n6.Left, n5)
	}

	if n6.Right != n7 {
		t.Errorf("found %#v expected %#v", n6.Right, n7)
	}
}

// n.number >= 1 AND NOT n.number = 10
// should result in the AST tree :-
//					 AND
//			   _____/	\____
//			  /				 \
//		  __>=__ 		  	 NOT
//		 /	    \			  |
//  n.number     1  		__=__
//						   /	 \
//					   n.number   10
func TestBasic_UpdateStackWithNot(t *testing.T) {
	exprStack := make(parser.StackExpr, 0)

	n1 := &ast.PropertyStmt{Variable: "n", Value: "number"}
	exprStack = exprStack.Push(n1)

	n2 := &ast.ComparisonExpr{Comparison: expressions.GTE}
	exprStack = exprStack.Push(n2)

	n3 := &ast.Ident{Data: 1}
	exprStack = exprStack.Push(n3)

	n4 := &ast.BooleanExpr{Boolean: ast.AND}
	exprStack = exprStack.Push(n4)

	n5 := &ast.NotExpr{}
	exprStack = exprStack.Push(n5)

	n6 := &ast.PropertyStmt{Variable: "n", Value: "number"}
	exprStack = exprStack.Push(n6)

	n7 := &ast.ComparisonExpr{Comparison: expressions.EQ}
	exprStack = exprStack.Push(n7)

	n8 := &ast.Ident{Data: 10}
	exprStack = exprStack.Push(n8)

	result, _ := exprStack.Shunt()

	if result != n4 {
		t.Errorf("found %#v expected %#v", result, n4)
	}

	if n4.Left != n2 {
		t.Errorf("found %#v expected %#v", n4.Left, n2)
	}

	if n2.Left != n1 {
		t.Errorf("found %#v expected %#v", n2.Left, n1)
	}

	if n2.Right != n3 {
		t.Errorf("found %#v expected %#v", n2.Right, n3)
	}

	if n4.Right != n5 {
		t.Errorf("found %#v expected %#v", n4.Right, n5)
	}

	if n5.Left != n7 {
		t.Errorf("found %#v expected %#v", n5.Left, n7)
	}

	if n7.Left != n6 {
		t.Errorf("found %#v expected %#v", n7.Left, n6)
	}

	if n7.Right != n8 {
		t.Errorf("found %#v expected %#v", n7.Right, n8)
	}
}

// (n.name = 'Peter' AND n.age < 30 XOR n.age > 30) AND n.name = 'Tobias'
// Not really a valid query but logically correct and should result in the AST tree :-
//
//	                       _________AND_________
//		                  /	    	            \
//		          _______XOR_______		      ___=___
//	             /    	           \	     /	     \
//	            /		            \     n.name   'Tobias'
//        _____AMD_____		      ___>___
//       /	           \	     /       \
//	    /		        \       n.age    30
//   ___=___           ___<___
//  |      \		  |       \
// n.name 'Peter'   n.age     30
func TestBasic_Parentheses(t *testing.T) {
	exprStack := make(parser.StackExpr, 0)

	n1 := &ast.ParenthesesExpr{Parentheses: ast.LPAREN}
	exprStack = exprStack.Push(n1)

	n2 := &ast.PropertyStmt{Variable: "n", Value: "name"}
	exprStack = exprStack.Push(n2)

	n3 := &ast.ComparisonExpr{Comparison: expressions.EQ}
	exprStack = exprStack.Push(n3)

	n4 := &ast.Ident{Data: "Peter"}
	exprStack = exprStack.Push(n4)

	n5 := &ast.BooleanExpr{Boolean: ast.AND}
	exprStack = exprStack.Push(n5)

	n6 := &ast.PropertyStmt{Variable: "n", Value: "age"}
	exprStack = exprStack.Push(n6)

	n7 := &ast.ComparisonExpr{Comparison: expressions.LT}
	exprStack = exprStack.Push(n7)

	n8 := &ast.Ident{Data: 30}
	exprStack = exprStack.Push(n8)

	n9 := &ast.BooleanExpr{Boolean: ast.XOR}
	exprStack = exprStack.Push(n9)

	n10 := &ast.PropertyStmt{Variable: "n", Value: "age"}
	exprStack = exprStack.Push(n10)

	n11 := &ast.ComparisonExpr{Comparison: expressions.GT}
	exprStack = exprStack.Push(n11)

	n12 := &ast.Ident{Data: 30}
	exprStack = exprStack.Push(n12)

	n13 := &ast.ParenthesesExpr{Parentheses: ast.RPAREN}
	exprStack = exprStack.Push(n13)

	n14 := &ast.BooleanExpr{Boolean: ast.AND}
	exprStack = exprStack.Push(n14)

	n15 := &ast.PropertyStmt{Variable: "n", Value: "name"}
	exprStack = exprStack.Push(n15)

	n16 := &ast.ComparisonExpr{Comparison: expressions.EQ}
	exprStack = exprStack.Push(n16)

	n17 := &ast.Ident{Data: "Tobias"}
	exprStack = exprStack.Push(n17)

	result, _ := exprStack.Shunt()

	if result != n14 {
		t.Errorf("found %#v expected %#v", result, n14)
	}

	if n14.Left != n9 {
		t.Errorf("found %#v expected %#v", n14.Left, n9)
	}

	if n9.Left != n5 {
		t.Errorf("found %#v expected %#v", n9.Left, n5)
	}

	if n5.Left != n3 {
		t.Errorf("found %#v expected %#v", n5.Left, n3)
	}

	if n3.Left != n2 {
		t.Errorf("found %#v expected %#v", n3.Left, n2)
	}

	if n3.Right != n4 {
		t.Errorf("found %#v expected %#v", n3.Right, n4)
	}

	if n5.Right != n7 {
		t.Errorf("found %#v expected %#v", n5.Right, n7)
	}

	if n7.Left != n6 {
		t.Errorf("found %#v expected %#v", n7.Left, n6)
	}

	if n7.Right != n8 {
		t.Errorf("found %#v expected %#v", n7.Right, n8)
	}

	if n9.Right != n11 {
		t.Errorf("found %#v expected %#v", n9.Right, n11)
	}

	if n11.Left != n10 {
		t.Errorf("found %#v expected %#v", n11.Left, n11)
	}

	if n11.Right != n12 {
		t.Errorf("found %#v expected %#v", n11.Right, n12)
	}

	if n14.Right != n16 {
		t.Errorf("found %#v expected %#v", n14.Right, n16)
	}

	if n16.Left != n15 {
		t.Errorf("found %#v expected %#v", n16.Left, n15)
	}

	if n16.Right != n17 {
		t.Errorf("found %#v expected %#v", n16.Right, n17)
	}
}

// n.name = 'Peter' XOR (n.age < 30 AND n.name = 'Tobias') OR NOT (n.name = 'Tobias' OR n.name = 'Peter')
// should result in the AST tree :-
//						     	  	   OR
//				   ___________________/  \____________________
//				  /						 	  		 		  \
//		 	_____XOR_____							 		  NOT
//		   /	 	     \					   		  		   |
//		__=__		 	__AND__				 			 ______OR_____
//	   /   	 |	       /       \		    			/			  \
// n.name  'Peter'  __<__     __=__					 __=__		 	 __=__
//				   /	|	 |	   \				/	  \	    	/ 	  \
//				  /		|	 |	    \			n.name  'Tobias'  n.name  'Peter'
// 		    	n.age  30   n.name  'Tobias'
func TestDeep_UpdateStack(t *testing.T) {
	exprStack := make(parser.StackExpr, 0)

	n1 := &ast.PropertyStmt{Variable: "n", Value: "name"}
	exprStack = exprStack.Push(n1)

	n2 := &ast.ComparisonExpr{Comparison: expressions.EQ}
	exprStack = exprStack.Push(n2)

	n3 := &ast.Ident{Data: "Peter"}
	exprStack = exprStack.Push(n3)

	n4 := &ast.BooleanExpr{Boolean: ast.XOR}
	exprStack = exprStack.Push(n4)

	n5 := &ast.ParenthesesExpr{Parentheses: ast.LPAREN}
	exprStack = exprStack.Push(n5)

	n6 := &ast.PropertyStmt{Variable: "n", Value: "age"}
	exprStack = exprStack.Push(n6)

	n7 := &ast.ComparisonExpr{Comparison: expressions.LT}
	exprStack = exprStack.Push(n7)

	n8 := &ast.Ident{Data: 30}
	exprStack = exprStack.Push(n8)

	n9 := &ast.BooleanExpr{Boolean: ast.AND}
	exprStack = exprStack.Push(n9)

	n10 := &ast.PropertyStmt{Variable: "n", Value: "name"}
	exprStack = exprStack.Push(n10)

	n11 := &ast.ComparisonExpr{Comparison: expressions.EQ}
	exprStack = exprStack.Push(n11)

	n12 := &ast.Ident{Data: "Tobias"}
	exprStack = exprStack.Push(n12)

	n13 := &ast.ParenthesesExpr{Parentheses: ast.RPAREN}
	exprStack = exprStack.Push(n13)

	n14 := &ast.BooleanExpr{Boolean: ast.OR}
	exprStack = exprStack.Push(n14)

	n15 := &ast.NotExpr{}
	exprStack = exprStack.Push(n15)

	n16 := &ast.ParenthesesExpr{Parentheses: ast.LPAREN}
	exprStack = exprStack.Push(n16)

	n17 := &ast.PropertyStmt{Variable: "n", Value: "name"}
	exprStack = exprStack.Push(n17)

	n18 := &ast.ComparisonExpr{Comparison: expressions.EQ}
	exprStack = exprStack.Push(n18)

	n19 := &ast.Ident{Data: "Tobias"}
	exprStack = exprStack.Push(n19)

	n20 := &ast.BooleanExpr{Boolean: ast.OR}
	exprStack = exprStack.Push(n20)

	n21 := &ast.PropertyStmt{Variable: "n", Value: "name"}
	exprStack = exprStack.Push(n21)

	n22 := &ast.ComparisonExpr{Comparison: expressions.EQ}
	exprStack = exprStack.Push(n22)

	n23 := &ast.Ident{Data: "Peter"}
	exprStack = exprStack.Push(n23)

	n24 := &ast.ParenthesesExpr{Parentheses: ast.RPAREN}
	exprStack = exprStack.Push(n24)

	result, _ := exprStack.Shunt()

	if result != n14 {
		t.Errorf("found %#v expected %#v", result, n14)
	}

	if n14.Left != n4 {
		t.Errorf("found %#v expected %#v", n14.Left, n4)
	}

	if n4.Left != n2 {
		t.Errorf("found %#v expected %#v", n4.Left, n2)
	}

	if n2.Left != n1 {
		t.Errorf("found %#v expected %#v", n2.Left, n1)
	}

	if n4.Right != n9 {
		t.Errorf("found %#v expected %#v", n4.Right, n9)
	}

	if n9.Left != n7 {
		t.Errorf("found %#v expected %#v", n9.Left, n7)
	}

	if n7.Left != n6 {
		t.Errorf("found %#v expected %#v", n7.Left, n6)
	}

	if n7.Right != n8 {
		t.Errorf("found %#v expected %#v", n7.Right, n8)
	}

	if n9.Right != n11 {
		t.Errorf("found %#v expected %#v", n9.Right, n11)
	}

	if n11.Left != n10 {
		t.Errorf("found %#v expected %#v", n11.Left, n10)
	}

	if n11.Right != n12 {
		t.Errorf("found %#v expected %#v", n11.Right, n12)
	}

	if n14.Right != n15 {
		t.Errorf("found %#v expected %#v", n14.Right, n15)
	}

	if n15.Left != n20 {
		t.Errorf("found %#v expected %#v", n15.Left, n20)
	}

	if n20.Left != n18 {
		t.Errorf("found %#v expected %#v", n20.Left, n18)
	}

	if n18.Left != n17 {
		t.Errorf("found %#v expected %#v", n18.Left, n17)
	}

	if n18.Right != n19 {
		t.Errorf("found %#v expected %#v", n18.Right, n19)
	}

	if n20.Right != n22 {
		t.Errorf("found %#v expected %#v", n20.Right, n22)
	}

	if n22.Left != n21 {
		t.Errorf("found %#v expected %#v", n22.Left, n22)
	}

	if n22.Right != n23 {
		t.Errorf("found %#v expected %#v", n22.Right, n23)
	}
}
