package operators

import (
	"github.com/RossMerr/Caudex.Graph/keyvaluestore"
)

// Scan is a source operator that accesses a table and returns the set of all its tuples
type Scan struct {
	iterator keyvaluestore.Iterator
}

// NewScan returns a Scan
func NewScan(storage keyvaluestore.Storage) *Scan {
	return &Scan{
		iterator: storage.Each(),
	}
}

func (s *Scan) Next() (interface{}, bool) {
	return s.iterator()
}
