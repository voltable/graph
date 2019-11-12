package expressions

// XORSwap Swaps the boolean
func XORSwap(b bool) bool {
	return !b
}

// XORExclusive Exclusive or
func XORExclusive(left bool, right bool) bool {
	return (left || right) && !(left && right)
}

type Boolean Binary

const (
	// and a bitwise or logical and operation, such as (a And b).
	and Boolean = Boolean(subtract) + iota + 1
	// or a bitwise or logical or operation, such as (a Or b).
	or
	// xor a bitwise or logical xor operation, such as (a Xor b).
	xor

// NOT a bitwise complement or logical negation operation (Not a).
//NOT
)
