package query

import (
	"bytes"
	"errors"

	"github.com/google/uuid"
	"github.com/voltable/graph/arch"
	"github.com/voltable/graph/encoding/wcs"
	"github.com/voltable/graph/widecolumnstore"
)

func Weight(s widecolumnstore.KeyValue) float64 {
	key := &widecolumnstore.Key{}
	key.Unmarshal(s.Key)

	if bytes.Equal(key.Column.Family, wcs.Relationship) {
		to := widecolumnstore.Unmarshal(s.Value)
		return to.(float64)
	}

	if bytes.Equal(key.ID, wcs.Relationship) {
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

	if bytes.Equal(key.Column.Family, wcs.Relationship) {
		return uuid.FromBytes(key.Column.Qualifier)
	}

	if bytes.Equal(key.ID, wcs.Relationship) {
		to := widecolumnstore.Unmarshal(s.Value)
		return to.(uuid.UUID), nil
	}

	return uuid.Nil, ErrIDNotFound
}

// ParseKeyToUUID looks for the UUID in the Key
func ParseKeyToUUID(key widecolumnstore.Key) (uuid.UUID, error) {
	if bytes.Equal(key.Column.Family, wcs.Label) {
		return uuid.FromBytes(key.ID)
	}

	if bytes.Equal(key.Column.Family, wcs.Properties) {
		return uuid.FromBytes(key.ID)
	}

	if bytes.Equal(key.Column.Family, wcs.Relationship) {
		return uuid.FromBytes(key.ID)
	}

	if bytes.Equal(key.Column.Family, wcs.Relationshipproperties) {
		return uuid.FromBytes(key.ID)
	}

	// Transpose

	if bytes.Equal(key.ID, wcs.TLabel) {
		return uuid.FromBytes(key.Column.Qualifier)
	}

	if bytes.Equal(key.ID, wcs.TProperties) {
		return uuid.FromBytes(key.Column.Qualifier)
	}

	if bytes.Equal(key.Column.Family, wcs.TRelationship) {
		return uuid.FromBytes(key.ID)
	}

	if bytes.Equal(key.ID, wcs.TRelationshipproperties) {
		return uuid.FromBytes(key.Column.Qualifier)
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

	if bytes.Equal(key.Column.Family, wcs.Properties) {

		if value == string(key.Column.Qualifier) {
			return widecolumnstore.Unmarshal(s.kv.Value)
		}
	}

	return nil
}
