package operators

import (
	"github.com/google/uuid"
	"github.com/voltable/graph"
	"github.com/voltable/graph/widecolumnstore"
)

var _ Nullary = (*Create)(nil)

// Create fetches all tuples with a specific label.
type Create struct {
	keyValues []*widecolumnstore.KeyValue
	storage widecolumnstore.Storage
}

// NewCreate returns a Create
func NewCreate(storage widecolumnstore.Storage, id uuid.UUID, variable string, label string, properties map[string]interface{}) (*Create, error) {

	keyValues := make([]*widecolumnstore.KeyValue, 0)

	keyValues = append(keyValues, &widecolumnstore.KeyValue{
		Key: &widecolumnstore.Key{
			RowKey:       id[:],
			ColumnFamily: ID,
		},
		Value: nil,
	})

	keyValues = append(keyValues, &widecolumnstore.KeyValue{
		Key: &widecolumnstore.Key{
			RowKey:          id[:],
			ColumnFamily:    Label,
			ColumnQualifier: []byte(label),
		},
		Value: nil,
	})

	for key, value := range properties  {
		keyValues = append(keyValues, &widecolumnstore.KeyValue{
			Key: &widecolumnstore.Key{
				RowKey:          id[:],
				ColumnFamily:    Properties,
				ColumnQualifier: []byte(key),
			},
			Value: widecolumnstore.NewAny(value),
		})
	}

	return &Create{
		storage: storage,
		keyValues:keyValues,
	}, nil
}

func (s *Create) Next() (widecolumnstore.Iterator, graph.Statistics) {
	statistics := graph.NewStatistics()

	statistics.DbHits.CreateNode += 1
	statistics.Rows += len(s.keyValues)

	_ = s.storage.Create(s.keyValues...)

	return func() (widecolumnstore.KeyValue, bool) {
		return widecolumnstore.KeyValue{}, false
	}, statistics
}

func (s *Create) Op() {}
