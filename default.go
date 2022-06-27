package graph

func Zero[T any]() T {
	return *new(T)
}

func IsZero[T comparable](v T) bool {
	return v == *new(T)
}

func Default[T any]() T {
	return *new(T)
}
