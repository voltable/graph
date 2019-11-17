package expressions

import (
	"github.com/pkg/errors"
	"reflect"
)

type ExpressionVisitor interface{
	Visit(expr Expression) Expression
	VisitExtension(expr Expression) Expression
	VisitParameter(expr *ParameterExpression) Expression
	VisitConstant(expr *ConstantExpression) Expression
	VisitConditional(expr *ConditionalExpression) Expression
	VisitBinary(expr BinaryExpression) Expression
	VisitLambda(expr *LambdaExpression)Expression
}

func baseVisit(base ExpressionVisitor, expr Expression) Expression {
	if expr != nil {
		return expr.Accept(base)
	}
	return nil
}

func baseVisitParameter(base ExpressionVisitor, expr *ParameterExpression) Expression {
	return expr
}

func baseVisitExtension(base ExpressionVisitor, expr Expression) Expression {
	return expr.VisitChildren(base)
}

func baseVisitConstant(base ExpressionVisitor, expr Expression) Expression {
	return expr
}

func baseVisitConditional(base ExpressionVisitor, expr *ConditionalExpression) Expression {
	test := base.Visit(expr.GetTest())
	ifTrue:= base.Visit(expr.GetIfTrue())
	ifFalse := base.Visit(expr.GetIfFalse())
	return expr.Update(test, ifTrue, ifFalse)
}


// baseVisitBinary Visits the children of the BinaryExpression node
func baseVisitBinary(base ExpressionVisitor, expr BinaryExpression) Expression {
	// Walk children in evaluation order: left, conversion, right
	conversion := base.Visit(expr.GetConversion())

	if lambda, ok := conversion.(*LambdaExpression); ok {
		left := base.Visit(expr.GetLeft())
		right := base.Visit(expr.GetRight())
		after := expr.Update(left.(TerminalExpression), lambda, right.(TerminalExpression))
		return validateBinary(expr, after)
	}

	panic(ArgumentTypesMustBeLambda)
}

func validateBinary(before, after BinaryExpression) BinaryExpression {
	if before != after {
		err := validateChildType(before.GetLeft().Kind(), after.GetLeft().Kind(), "VisitBinary")
		if err != nil {
			panic(err)
		}

		err = validateChildType(before.GetRight().Kind(), after.GetRight().Kind(), "VisitBinary")
		if err != nil {
			panic(err)
		}
	}
	return after
}

func validateChildType(before, after reflect.Kind, name string) error {
	if before == after {
		return nil
	}

	return errors.Wrap(MustRewriteChildToSameType, name)
}

func baseVisitLambda(base ExpressionVisitor, expr *LambdaExpression) Expression {
	parameters := make([]*ParameterExpression, 0)
	for _, parameter:= range expr.parameters {
		p := base.Visit(parameter)
		if pe,  ok := p.(*ParameterExpression); ok {
			parameters = append(parameters, pe)
		}
	}

	return expr.Update(base.Visit(expr.body), parameters)
}