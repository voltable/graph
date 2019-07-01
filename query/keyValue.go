package query

import (
	"bytes"
	"errors"

	"github.com/voltable/graph/arch"
	"github.com/voltable/graph/uuid"
	"github.com/voltable/graph/widecolumnstore"
)

func Weight(s widecolumnstore.KeyValue) float64 {
	key := &widecolumnstore.Key{}
	key.Unmarshal(s.Key)

	if bytes.Equal(key.Column.Family, Relationship) {
		to := widecolumnstore.Unmarshal(s.Value)
		return to.(float64)
	}

	if bytes.Equal(key.ID, Relationship) {
		return arch.DecodeFloat64Bytes(key.Column.Extended).(float64)
	}

	return 0
}

var (
	ErrIDNotFound = errors.New("ID not found in KeyValue")
)

// To looks only for a ID as part of a Relationship
func To(s widecolumnstore.KeyValue) (uuid.UUID, error) {
	key := &widecolumnstore.Key{}
	key.Unmarshal(s.Key)

	if bytes.Equal(key.Column.Family, Relationship) {
		return uuid.SliceToUUID(key.Column.Qualifier), nil
	}

	if bytes.Equal(key.ID, Relationship) {
		to := widecolumnstore.Unmarshal(s.Value)
		return to.(uuid.UUID), nil
	}

	return uuid.Nil, ErrIDNotFound
}

// ParseKeyToUUID looks for the UUID in the Key
func ParseKeyToUUID(key widecolumnstore.Key) (uuid.UUID, error) {
	if bytes.Equal(key.Column.Family, Label) {
		return uuid.SliceToUUID(key.ID), nil
	}

	if bytes.Equal(key.Column.Family, Properties) {
		return uuid.SliceToUUID(key.ID), nil
	}

	if bytes.Equal(key.Column.Family, Relationship) {
		return uuid.SliceToUUID(key.ID), nil
	}

	if bytes.Equal(key.Column.Family, Relationshipproperties) {
		return uuid.SliceToUUID(key.ID), nil
	}

	// Transpose

	if bytes.Equal(key.ID, TLabel) {
		return uuid.SliceToUUID(key.Column.Qualifier), nil
	}

	if bytes.Equal(key.ID, TProperties) {
		return uuid.SliceToUUID(key.Column.Qualifier), nil
	}

	if bytes.Equal(key.Column.Family, TRelationship) {
		return uuid.SliceToUUID(key.ID), nil
	}

	if bytes.Equal(key.ID, TRelationshipproperties) {
		return uuid.SliceToUUID(key.Column.Qualifier), nil
	}

	return uuid.Nil, ErrIDNotFound
}

// UUID looks for the UUID in the KeyValue
func UUID(s *widecolumnstore.KeyValue) (uuid.UUID, error) {
	key := widecolumnstore.Key{}
	key.Unmarshal(s.Key)
	return ParseKeyToUUID(key)
}

type KeyValueWrapper struct {
	kv *widecolumnstore.KeyValue
}

func NewKeyValueWrapper(kv *widecolumnstore.KeyValue) *KeyValueWrapper {
	return &KeyValueWrapper{kv}
}

func (s *KeyValueWrapper) Interpret(value string) interface{} {
	key := &widecolumnstore.Key{}
	key.Unmarshal(s.kv.Key)

	if bytes.Equal(key.Column.Family, Properties) {

		if value == string(key.Column.Qualifier) {
			return widecolumnstore.Unmarshal(s.kv.Value)
		}
	}

	return nil
}
