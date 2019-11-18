package expressions

type Binary int

// A BinaryExpression node represents a binary expression.
type BinaryExpression interface {
	NonTerminalExpression
	binary()
	Type() Binary
	Update(left, right TerminalExpression) BinaryExpression
}

func baseUpdate(b BinaryExpression, left, right TerminalExpression) BinaryExpression {
	if left == b.GetLeft() && right == b.GetRight()  {
		return b
	}

	return makeBinary(b.Type(), left, right)
}

func makeBinary(t Binary, left, right TerminalExpression) BinaryExpression{
	switch t {
	case Binary(add):
		return Add(left, right)
	default:
		return Subtract(left, right)
	}

	//todo add all other case's
}
