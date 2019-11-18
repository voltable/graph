package expressions

var _ ExpressionVisitor = (*DynamicExpressionVisitor )(nil)

type DynamicExpressionVisitor  struct {
}

func NewDynamicExpressionVisitor () *DynamicExpressionVisitor {
	return &DynamicExpressionVisitor {
	}
}

func (s *DynamicExpressionVisitor ) VisitLambda(expr *LambdaExpression) Expression {
	return baseVisitLambda(s, expr)
}

func (s *DynamicExpressionVisitor ) VisitBinary(expr BinaryExpression) Expression {
	return baseVisitBinary(s, expr)
}

func (s *DynamicExpressionVisitor ) Visit(expr Expression) Expression {
	return baseVisit(s, expr)
}

func (s *DynamicExpressionVisitor ) VisitExtension(expr Expression) Expression {
	return expr.VisitChildren(s)
}

func (s *DynamicExpressionVisitor ) VisitParameter(expr *ParameterExpression) Expression {
	return expr
}

func (s *DynamicExpressionVisitor ) VisitConstant(expr *ConstantExpression) Expression {
	return expr
}

func (s *DynamicExpressionVisitor ) VisitConditional(expr *ConditionalExpression) Expression {
	return baseVisitConditional(s, expr)
}