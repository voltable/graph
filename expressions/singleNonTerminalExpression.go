package expressions

// SingleNonTerminalExpression is a NonTerminal symbol which only can be broken down once e.g. a NotExpression
type SingleNonTerminalExpression interface {
	Expression
	GetValue() Expression
}

