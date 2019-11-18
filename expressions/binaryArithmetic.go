package expressions

// BinaryArithmetic operations
type BinaryArithmetic Binary

const (
	// An addition operation, such as (a + b).
	add BinaryArithmetic = iota
	// A division operation, such as (a / b).
	divide
	// An arithmetic remainder operation, such as (a % b).
	modulo
	// A multiplication operation, such as (a * b).
	multiply
	// A mathematical operation that raises a number to a power, such as (a ^ b).
	power
	// A subtraction operation, such as (a - b)
	subtract
)
