package expressions

import "reflect"

// Expr all expression nodes implement the Expr interface.
type Expression  interface {
	//ExpressionVisitor
	// Reduces the expression node to a simpler expression.
	Reduce() Expression

	//ReduceAndCheck this expression node to a simpler expression
	ReduceAndCheck() Expression

	// Accept visitor to visit this expression node with.
	Accept(visitor ExpressionVisitor) Expression

	// VisitChildren reduces the expression node and then calls the visitor delegate on the reduced expression.
	VisitChildren(visitor ExpressionVisitor) Expression

	Kind() reflect.Kind

	String() string
}

// baseCanReduce indicates whether the expression can be reduced.
func baseCanReduce(base Expression) bool {
	if _, ok := base.(TerminalExpression); ok {
		return false
	}

	return true
}

func baseReduceAndCheck(base Expression) Expression {
	if !baseCanReduce(base) {
		panic(MustBeReducible)
	}

	newNode := base.Reduce()

	if newNode == nil {
		if newNode == base {
			panic(MustReduceToDifferent)
		}
	}

	if !AreReferenceAssignable(base, newNode) {
		panic(ReducedNotCompatible)
	}

	return newNode
}

func baseReduce(base Expression) Expression {
	if baseCanReduce(base) {
		panic(ReducibleMustOverrideReduce)
	}

	return base
}

func baseVisitChildren(base Expression, visitor ExpressionVisitor) Expression  {
	if !baseCanReduce(base) {
		panic(MustBeReducible)
	}

	expr := base.ReduceAndCheck()
	return visitor.Visit(expr)
}

func baseAccept(base Expression, visitor ExpressionVisitor) Expression {
	return visitor.VisitExtension(base)
}
