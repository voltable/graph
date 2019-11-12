package expressions

import "reflect"

// ConditionalExpression represents an expression that has a conditional operator.
type ConditionalExpression struct {
	test Expression
	ifTrue Expression
	ifFalse Expression
	kind reflect.Kind
}

func (s *ConditionalExpression) String() string {
	return ExpressionToString(s)
}

var _ Expression = (*ConditionalExpression)(nil)

func (s *ConditionalExpression) Kind() reflect.Kind {
	return s.kind
}

func (s *ConditionalExpression) GetTest() Expression {
	return s.test
}

func (s *ConditionalExpression) GetIfTrue() Expression {
	return s.ifTrue
}

func (s *ConditionalExpression) GetIfFalse() Expression {
	return s.ifFalse
}


func (s *ConditionalExpression) Update(test, ifTrue, ifFalse Expression) (Expression, error) {
	if test == s.test && ifTrue == s.ifTrue && ifFalse == s.ifFalse {
		return s, nil
	}
	return Condition(test, ifTrue, ifFalse)
}

func (s *ConditionalExpression) Reduce() (Expression, error) {
	return baseReduce(s)
}

func (s *ConditionalExpression) ReduceAndCheck() (Expression, error) {
	return baseReduceAndCheck(s)
}

func (s *ConditionalExpression) Accept(visitor ExpressionVisitor) (Expression, error) {
	return visitor.VisitConditional(s)
}

func (s *ConditionalExpression) VisitChildren(visitor ExpressionVisitor) (Expression, error) {
	return baseVisitChildren(s, visitor)
}

// Condition creates a ConditionalExpression.
func Condition(test, ifTrue, ifFalse Expression) (*ConditionalExpression, error) {
	if ifTrue.Kind() != reflect.Bool {
		return nil, ArgumentMustBeBoolean
	}

	if ifTrue.Kind() != ifFalse.Kind() {
		return nil, ArgumentTypesMustMatch
	}

	return &ConditionalExpression{
		test: test,
		ifTrue: ifTrue,
		ifFalse:ifFalse,
	}, nil
}

// IfThen condition creates a ConditionalExpression.
func IfThen(test, ifTrue Expression)  (*ConditionalExpression, error) {
	return Condition(test, ifTrue, Empty())
}

// IfThenElse condition creates a ConditionalExpression.
func IfThenElse(test, ifTrue, ifFalse Expression)  (*ConditionalExpression, error) {
	return Condition(test, ifTrue, ifFalse)
}