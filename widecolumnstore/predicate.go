package widecolumnstore

// Predicate apply the predicate over the key/value
type Predicate func(interface{}) bool

var EmptyPredicate = func(interface{}) bool {
	return true
}
