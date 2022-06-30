package storage

import "github.com/rossmerr/graphblas/constraints"

type Each[T constraints.Number] interface {
	Each() Iterator
}

type HasPrefix[T constraints.Number] interface {
	HasPrefix(Triple[T]) Iterator
}

type OperatorStorage[T constraints.Number] interface {
	Each[T]
	HasPrefix[T]
}
type Storage[T constraints.Number] interface {
	Each[T]
	HasPrefix[T]
	// Count number of keys/value pairs
	Count() int
	//Add(*Mutation)
	//	Query(string) Iterator[T]

	//	Create(triples ...*SparseValue) error
}
