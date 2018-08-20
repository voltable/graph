package widecolumnstore

type Storage interface {
	Each() Iterator
	HasPrefix([]byte) Iterator
	// Count number of keys/value pairs
	Count() int
	//Add(*Mutation)
}
