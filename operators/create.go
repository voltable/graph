package operators

import (
	//"github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/proto"
	"github.com/voltable/graph/widecolumnstore"
)

var _ Nullary = (*Create)(nil)

// Create fetches all tuples with a specific label.
type Create struct {
	iterator widecolumnstore.Iterator
}

// NewCreate returns a Create
func NewCreate(storage widecolumnstore.Storage, label string) (*Create, error) {
	// TODO need to marshal the prefix right
	k := &widecolumnstore.Key{}
	prefix, err := proto.Marshal(k)
	if err != nil {
		return nil, err
	}

	return &Create{
		iterator: storage.HasPrefix(prefix),
	}, nil
}

func (s *Create) Next() widecolumnstore.Iterator {
	return s.iterator
}

func (s *Create) Op() {}
