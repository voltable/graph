package widecolumnstore

type Each interface {
	Each() Iterator
}

type HasPrefix interface {
	HasPrefix([]byte) Iterator
}

type OperatorStorage interface {
	Each
	HasPrefix
}
type Storage interface {
	Each
	HasPrefix
	// Count number of keys/value pairs
	Count() int
	//Add(*Mutation)
	Query(Operator) Iterator

	Create(triples ...*KeyValue) error
}
