package operators

type Iterator interface {
	Next() (interface{}, bool)
}
