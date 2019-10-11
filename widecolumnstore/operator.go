package widecolumnstore

// Operator usually have 0, 1, or 2 arguments
type Operator interface {
	Op()
}

// Nullary an operation of arity 0, and hence call it nullary
type Nullary interface {
	Operator
	Next() Iterator
}

// Unary operation takes one argument
type Unary interface {
	Operator
	Next(i Iterator) Iterator
}

// Binary operation takes two arguments
type Binary interface {
	Operator
	Next(x, y Iterator) Iterator
}
