package operators

import "github.com/voltable/graph/widecolumnstore"

var _ Nullary = (*LabelScan)(nil)

// LabelScan fetches all tuples with a specific label.
type LabelScan struct {
	iterator widecolumnstore.Iterator
}

// NewLabelScan returns a LabelScan
func NewLabelScan(storage widecolumnstore.Storage, label string) *LabelScan {
	// TODO need to marshal the prefix right
	prefix := []byte(label)

	return &LabelScan{
		iterator: storage.HasPrefix(prefix),
	}
}

func (s *LabelScan) Next() widecolumnstore.Iterator {
	return s.iterator
}

func (s *LabelScan) Op() {}
