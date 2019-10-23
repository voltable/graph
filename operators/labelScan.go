package operators

import (
	"github.com/golang/protobuf/proto"
	"github.com/voltable/graph"
	"github.com/voltable/graph/widecolumnstore"
)

var _ Nullary = (*LabelScan)(nil)

// LabelScan fetches all tuples with a specific label.
type LabelScan struct {
	iterator widecolumnstore.Iterator
}

// NewLabelScan returns a LabelScan
func NewLabelScan(storage widecolumnstore.Storage, label string) (*LabelScan, error) {
	// TODO need to marshal the prefix right
	k := &widecolumnstore.Key{}

	prefix := []byte(label)

	prefix, err := proto.Marshal(k)
	if err != nil {
		return nil, err
	}

	return &LabelScan{
		iterator: storage.HasPrefix(prefix),
	}, nil
}

func (s *LabelScan) Next() (widecolumnstore.Iterator, graph.Statistics) {

	statistics := graph.NewStatistics()

	return s.iterator, statistics
}

func (s *LabelScan) Op() {}
