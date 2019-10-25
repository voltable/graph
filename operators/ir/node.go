package ir

import (
	"github.com/google/uuid"
	"github.com/voltable/graph/widecolumnstore"
)

type Node struct {
	Id         uuid.UUID
	Variable   string
	Label      string
	Properties map[string]interface{}
}

func (n *Node) Marshal(a *Actions) []*widecolumnstore.KeyValue {

	keyValues := make([]*widecolumnstore.KeyValue, 0)

	keyValues = append(keyValues, &widecolumnstore.KeyValue{
		Key: &widecolumnstore.Key{
			RowKey:       n.Id[:],
			ColumnFamily: Vertex,
			ColumnQualifier: ID,
		},
		Value: nil,
	})

	a.Nodes += 1

	if n.Label != EmptyString {

		keyValues = append(keyValues, &widecolumnstore.KeyValue{
			Key: &widecolumnstore.Key{
				RowKey:          n.Id[:],
				ColumnFamily:    Label,
				ColumnQualifier: []byte(n.Label),
			},
			Value: nil,
		})

		a.Labels += 1
	}

	for key, value := range n.Properties {

		if key != EmptyString {
			keyValues = append(keyValues, &widecolumnstore.KeyValue{
				Key: &widecolumnstore.Key{
					RowKey:          n.Id[:],
					ColumnFamily:    Properties,
					ColumnQualifier: []byte(key),
				},
				Value: widecolumnstore.NewAny(value),
			})
			a.Properties += 1
		}
	}

	return keyValues
}
