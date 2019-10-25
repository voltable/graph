package ir

import (
	"github.com/google/uuid"
	"github.com/voltable/graph/widecolumnstore"
)

type Relationship struct {
	Id uuid.UUID
	Variable string
	Type string
	Properties map[string]interface{}
}

func (n *Relationship) Marshal(a *Actions) []*widecolumnstore.KeyValue {

	keyValues := make([]*widecolumnstore.KeyValue, 0)

	keyValues = append(keyValues, &widecolumnstore.KeyValue{
		Key: &widecolumnstore.Key{
			RowKey:       n.Id[:],
			ColumnFamily: ID,
			ColumnQualifier: Edge,
		},
		Value: nil,
	})

	a.Relationships += 1

	if n.Type != EmptyString {
		keyValues = append(keyValues, &widecolumnstore.KeyValue{
			Key: &widecolumnstore.Key{
				RowKey:          n.Id[:],
				ColumnFamily:    EdgeType,
				ColumnQualifier: []byte(n.Type),
			},
			Value: nil,
		})

		a.Types += 1
	}

	for key, value := range n.Properties {
		if key != EmptyString {
			keyValues = append(keyValues, &widecolumnstore.KeyValue{
				Key: &widecolumnstore.Key{
					RowKey:          n.Id[:],
					ColumnFamily:    EdgePoperties,
					ColumnQualifier: []byte(key),
				},
				Value: widecolumnstore.NewAny(value),
			})

			a.Properties += 1
		}
	}

	return keyValues
}