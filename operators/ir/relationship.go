package ir

import (
	"github.com/google/uuid"
	"github.com/voltable/graph/operators/ir/delimiters"
	"github.com/voltable/graph/widecolumnstore"
)

type Relationship struct {
	Id uuid.UUID
	Variable Variable
	Type Type
	Properties *Properties
}

func (n *Relationship) Marshal(a *Actions) []*widecolumnstore.KeyValue {

	keyValues := make([]*widecolumnstore.KeyValue, 0)

	keyValues = append(keyValues, &widecolumnstore.KeyValue{
		Key: &widecolumnstore.Key{
			RowKey:          n.Id[:],
			ColumnFamily:    delimiters.ID,
			ColumnQualifier: delimiters.Edge,
		},
		Value: nil,
	})

	a.Relationships += 1

	if n.Type != EmptyString {
		keyValues = append(keyValues, &widecolumnstore.KeyValue{
			Key: &widecolumnstore.Key{
				RowKey:          n.Id[:],
				ColumnFamily:    delimiters.EdgeType,
				ColumnQualifier: []byte(n.Type),
			},
			Value: nil,
		})

		a.Types += 1
	}

	if n.Properties != nil {
		for key, value := range n.Properties.Map.Items {
			if key != EmptyString {
				keyValues = append(keyValues, &widecolumnstore.KeyValue{
					Key: &widecolumnstore.Key{
						RowKey:          n.Id[:],
						ColumnFamily:    delimiters.EdgePoperties,
						ColumnQualifier: []byte(key),
					},
					Value: widecolumnstore.NewAny(value),
				})

				a.Properties += 1
			}
		}
	}

	// Transpose

	keyValues = append(keyValues, &widecolumnstore.KeyValue{
		Key: &widecolumnstore.Key{
			RowKey:          delimiters.TID,
			ColumnFamily:    delimiters.TEdge,
			ColumnQualifier: n.Id[:],
		},
		Value: nil,
	})

	return keyValues
}