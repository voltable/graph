package expressions

// NonTerminalExpression is a NonTerminal symbol which can still be broken down e.g. a BooleanExpression
type NonTerminalExpression interface {
	Expression
	GetLeft() Expression
	GetRight() Expression
}

