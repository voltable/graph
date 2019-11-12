package operators

import (
	"github.com/golang/protobuf/proto"
	"github.com/voltable/graph"
	"github.com/voltable/graph/operators/ir"
	"github.com/voltable/graph/operators/ir/delimiters"
	"github.com/voltable/graph/widecolumnstore"
)

var _ Nullary = (*AllNodesScan )(nil)

// AllNodesScan operator reads all nodes from the node store. The variable that will contain the nodes is seen in the arguments. Any query using this operator is likely to encounter performance problems on a non-trivial database.
type AllNodesScan   struct {
	storage    widecolumnstore.Storage
	statistics *graph.Statistics
	variable   ir.Variable
}

// NewAllNodesScan returns a AllNodesScan
func NewAllNodesScan (storage widecolumnstore.Storage, statistics *graph.Statistics, variable ir.Variable) (*AllNodesScan , error) {
	return &AllNodesScan {
		storage: storage,
		statistics: statistics,
		variable: variable,
	}, nil
}

func (s *AllNodesScan ) Next() widecolumnstore.Iterator {
	 kv := &widecolumnstore.KeyValue{
		Key: &widecolumnstore.Key{
			RowKey:          delimiters.TID,
			ColumnFamily:    delimiters.TVertex,
			ColumnQualifier: nil,
		},
		Value: nil,
	}

	prefix, _ := proto.Marshal(kv)

	return s.storage.HasPrefix(prefix)
}

func (s *AllNodesScan ) Op() {}