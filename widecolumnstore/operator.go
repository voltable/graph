package widecolumnstore

type Operator interface {
	op()
	Next(i ...Iterator) Iterator
}
