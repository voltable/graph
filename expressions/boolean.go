package expressions

// XORSwap Swaps the boolean
func XORSwap(b bool) bool {
	return !b
}

// XORExclusive Exclusive or
func XORExclusive(left bool, right bool) bool {
	return (left || right) && !(left && right)
}

type Boolean int

const (
	// AND a bitwise or logical AND operation, such as (a And b).
	AND Boolean = iota
	// OR a bitwise or logical OR operation, such as (a Or b).
	OR
	// XOR a bitwise or logical XOR operation, such as (a Xor b).
	XOR

// NOT a bitwise complement or logical negation operation (Not a).
//NOT
)
