package keyvalue

type Storage interface {
	ForEach() Iterator
	HasPrefix([]byte) Iterator
}
