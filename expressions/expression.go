package expressions

import "reflect"

// Expr all expression nodes implement the Expr interface.
type Expression  interface {
	//ExpressionVisitor
	// Reduces the expression node to a simpler expression.
	Reduce() (Expression, error)

	//ReduceAndCheck this expression node to a simpler expression
	ReduceAndCheck() (Expression, error)

	// Accept visitor to visit this expression node with.
	Accept(visitor ExpressionVisitor) (Expression, error)

	// VisitChildren reduces the expression node and then calls the visitor delegate on the reduced expression.
	VisitChildren(visitor ExpressionVisitor) (Expression, error)

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

func baseReduceAndCheck(base Expression) (Expression, error) {
	if !baseCanReduce(base) {
		return nil, MustBeReducible
	}

	newNode, err := base.Reduce()
	if err != nil {
		return nil, err
	}

	if newNode == nil {
		if newNode == base {
			return nil, MustReduceToDifferent
		}
	}

	if !AreReferenceAssignable(base, newNode) {
		return nil, ReducedNotCompatible
	}

	return newNode, nil
}

func baseReduce(base Expression) (Expression, error) {
	if baseCanReduce(base) {
		return nil, ReducibleMustOverrideReduce
	}

	return base, nil
}

func baseVisitChildren(base Expression, visitor ExpressionVisitor) (Expression, error) {
	if !baseCanReduce(base) {
		return nil, MustBeReducible
	}

	expr, err := base.ReduceAndCheck()
	if err != nil {
		return nil, err
	}

	return visitor.Visit(expr)
}

func baseAccept(base Expression, visitor ExpressionVisitor) (Expression, error) {
	return visitor.VisitExtension(base)
}
