package expressions

type Binary int

// A BinaryExpression node represents a binary expression.
type BinaryExpression interface {
	NonTerminalExpression
	binary()
	Type() Binary
	Update(left TerminalExpression, conversion *LambdaExpression, right TerminalExpression) BinaryExpression
	GetConversion() *LambdaExpression
}

func baseUpdate(b BinaryExpression, left TerminalExpression, conversion *LambdaExpression, right TerminalExpression) BinaryExpression {
	if left == b.GetLeft() && right == b.GetRight() && conversion == b.GetConversion() {
		return b
	}

	return makeBinary(b.Type(), left, right, conversion)
}

func makeBinary(t Binary, left, right TerminalExpression, conversion *LambdaExpression) BinaryExpression{
	switch t {
	case Binary(add):
		return Add(left, right)
	default:
		return Subtract(left, right)
	}

	//todo add all other case's
}
