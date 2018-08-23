package widecolumnstore

var _ Operator = (*Scan)(nil)

// Scan is a source operator that accesses a table and returns the set of all its tuples
type Scan struct {
	iterator Iterator
}

// NewScan returns a Scan
func NewScan(storage Storage) *Scan {
	return &Scan{
		iterator: storage.Each(),
	}
}

func (s *Scan) Next(i ...Iterator) Iterator {
	return s.iterator
}

func (s *Scan) op() {}
