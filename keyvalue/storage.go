package keyvalue

type Storage interface {
	ForEach() Iterator
	Fetch(string) (*KeyValue, error)
	HasPrefix([]byte) Iterator
}
