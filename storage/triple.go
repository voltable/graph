package storage

import "github.com/rossmerr/graphblas/constraints"

type Triple[T constraints.Number] struct {
	Subject   string
	Predicate string
	Object    T
}

func (t *Triple[T]) Transpose() *Triple {
	return &Triple{
		Subject: t.Object,
	}
}
