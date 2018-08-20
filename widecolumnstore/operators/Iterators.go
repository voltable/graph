package operators

type Iterator interface {
	Next(Iterator) (interface{}, bool)
}
