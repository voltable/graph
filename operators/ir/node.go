package ir

import (
	"github.com/google/uuid"
	"github.com/voltable/graph/operators/ir/delimiters"
	"github.com/voltable/graph/widecolumnstore"
)

type Node struct {
	Id         uuid.UUID
	Variable   Variable
	Label      Label
	Properties *Properties
}

func (n *Node) Marshal(a *Actions) []*widecolumnstore.KeyValue {

	keyValues := make([]*widecolumnstore.KeyValue, 0)

	keyValues = append(keyValues, &widecolumnstore.KeyValue{
		Key: &widecolumnstore.Key{
			RowKey:          n.Id[:],
			ColumnFamily:    delimiters.Vertex,
			ColumnQualifier: delimiters.ID,
		},
		Value: nil,
	})

	a.Nodes += 1

	if n.Label != EmptyString {

		keyValues = append(keyValues, &widecolumnstore.KeyValue{
			Key: &widecolumnstore.Key{
				RowKey:          n.Id[:],
				ColumnFamily:    delimiters.Label,
				ColumnQualifier: []byte(n.Label),
			},
			Value: nil,
		})

		a.Labels += 1
	}

	if n.Properties != nil  {
		for key, value := range n.Properties.Map.Items {
			if key != EmptyString {
				keyValues = append(keyValues, &widecolumnstore.KeyValue{
					Key: &widecolumnstore.Key{
						RowKey:          n.Id[:],
						ColumnFamily:    delimiters.Properties,
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
			ColumnFamily:    delimiters.TVertex,
			ColumnQualifier: n.Id[:],
		},
		Value: nil,
	})



	return keyValues
}
