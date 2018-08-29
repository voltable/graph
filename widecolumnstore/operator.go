package widecolumnstore

type Operator interface {
	Op()
}

type Binary interface {
	Operator
	Next(x, y Iterator) Iterator
}

type Unary interface {
	Operator
	Next(i Iterator) Iterator
}
