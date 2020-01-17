package expressions

import "github.com/voltable/graph/expressions/stack"

func Compile(expression *LambdaExpression) Delegate {
	s := &stack.Stack{}

	return func(...interface{}) interface{} {

		switch v := expression.body.(type) {
		case *ConstantExpression:
			return v.GetValue()
		case *BinaryArithmeticExpression:
			v.Interpret(s)
		}

		return s.Pop()
	}
}
