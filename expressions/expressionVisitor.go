package expressions

import (
	"github.com/pkg/errors"
	"reflect"
)

type ExpressionVisitor interface{
	Visit(expr Expression) (Expression, error)
	VisitExtension(expr Expression) (Expression, error)
	VisitParameter(expr *ParameterExpression) (Expression, error)
	VisitConstant(expr *ConstantExpression) (Expression, error)
	VisitConditional(expr *ConditionalExpression) (Expression, error)
	VisitBinary(expr BinaryExpression) (Expression, error)
}

func baseVisit(base ExpressionVisitor, expr Expression) (Expression, error) {
	if expr != nil {
		return expr.Accept(base)
	}
	return nil, nil
}

func baseVisitParameter(base ExpressionVisitor, expr *ParameterExpression) (Expression, error) {
	return expr, nil
}


func baseVisitExtension(base ExpressionVisitor, expr Expression) (Expression, error) {
	return expr.VisitChildren(base)
}

func baseVisitConstant(base ExpressionVisitor, expr Expression) (Expression, error) {
	return expr, nil
}

func baseVisitConditional(base ExpressionVisitor, expr *ConditionalExpression) (Expression, error) {
	test, err := base.Visit(expr.GetTest())
	if err != nil {
		return nil, err
	}

	ifTrue, err := base.Visit(expr.GetIfTrue())
	if err != nil {
		return nil, err
	}

	ifFalse, err := base.Visit(expr.GetIfFalse())
	if err != nil {
		return nil, err
	}

	return expr.Update(test, ifTrue, ifFalse)
}


// baseVisitBinary Visits the children of the BinaryExpression node
func baseVisitBinary(base ExpressionVisitor, expr BinaryExpression) (Expression, error) {
	// Walk children in evaluation order: left, conversion, right
	left, err := base.Visit(expr.GetLeft())
	if err != nil {
		return nil, err
	}

	var right Expression
	right, err = base.Visit(expr.GetRight())
	if err != nil {
		return nil, err
	}

	var conversion Expression
	conversion, err = base.Visit(expr.GetConversion())
	if err != nil {
		return nil, err
	}
	if lambda, ok := conversion.(*LambdaExpression); ok {

		var after BinaryExpression
		after, err = expr.Update(left.(TerminalExpression), lambda, right.(TerminalExpression))
		if err != nil {
			return nil, err
		}

		return validateBinary(expr, after)
	}

	return nil, ArgumentTypesMustBeLambda
}

func validateBinary(before, after BinaryExpression) (BinaryExpression, error) {
	if before != after {
		err := validateChildType(before.GetLeft().Kind(), after.GetLeft().Kind(), "VisitBinary")
		if err != nil {
			return nil, err
		}

		err = validateChildType(before.GetRight().Kind(), after.GetRight().Kind(), "VisitBinary")
		if err != nil {
			return nil, err
		}
	}
	return after, nil
}

func validateChildType(before, after reflect.Kind, name string) error {
	if before == after {
		return nil
	}

	return errors.Wrap(MustRewriteChildToSameType, name)
}