package expressions

func Compile(expression *LambdaExpression) Delegate {
	return func([]interface{}) interface{} {
		c := expression.body.(*ConstantExpression)
		return c.GetValue()
	}
}