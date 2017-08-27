package expressions

// XORSwap Swaps the boolean
func XORSwap(b bool) bool {
	return !b
}

// XOR Exclusive or
func XOR(left bool, right bool) bool {
	return (left || right) && !(left && right)
}
