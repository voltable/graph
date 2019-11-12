package expressions

// TerminalExpression is a Terminal symbol which cannot be broken down further e.g. a ParameterExpression
type TerminalExpression interface {
	Expression
	GetValue() interface{}
}

