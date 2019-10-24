package operators

import (
	"github.com/voltable/graph"
	"github.com/voltable/graph/operators/ir"
	"github.com/voltable/graph/widecolumnstore"
)

var _ Nullary = (*Create)(nil)

// Create fetches all tuples with a specific label.
type Create struct {
	nodes []*ir.Node
	storage widecolumnstore.Storage
}

// NewCreate returns a Create
func NewCreate(storage widecolumnstore.Storage, nodes []*ir.Node) (*Create, error) {
	return &Create{
		storage: storage,
		nodes: nodes,
	}, nil
}

func (s *Create) Next() (widecolumnstore.Iterator, graph.Statistics) {
	statistics := graph.NewStatistics()

	keyValues := make([]*widecolumnstore.KeyValue, 0)

	for _, n := range s.nodes {
		keyValues = append(keyValues, &widecolumnstore.KeyValue{
			Key: &widecolumnstore.Key{
				RowKey:       n.Id[:],
				ColumnFamily: ID,
			},
			Value: nil,
		})

		keyValues = append(keyValues, &widecolumnstore.KeyValue{
			Key: &widecolumnstore.Key{
				RowKey:          n.Id[:],
				ColumnFamily:    Label,
				ColumnQualifier: []byte(n.Label),
			},
			Value: nil,
		})

		for key, value := range n.Properties  {
			keyValues = append(keyValues, &widecolumnstore.KeyValue{
				Key: &widecolumnstore.Key{
					RowKey:          n.Id[:],
					ColumnFamily:    Properties,
					ColumnQualifier: []byte(key),
				},
				Value: widecolumnstore.NewAny(value),
			})
		}

		statistics.DbHits.CreateNode += 1
	}

	statistics.Rows += len(keyValues)

	_ = s.storage.Create(keyValues...)

	return func() (widecolumnstore.KeyValue, bool) {
		return widecolumnstore.KeyValue{}, false
	}, statistics
}

func (s *Create) Op() {}
