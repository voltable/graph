package query

import (
	"bytes"

	"github.com/RossMerr/Caudex.Graph/arch"
	"github.com/RossMerr/Caudex.Graph/uuid"
	"github.com/RossMerr/Caudex.Graph/widecolumnstore"
)

func Weight(s *widecolumnstore.KeyValue) float64 {
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

func To(s *widecolumnstore.KeyValue) *uuid.UUID {
	key := &widecolumnstore.Key{}
	key.Unmarshal(s.Key)

	if bytes.Equal(key.Column.Family, Relationship) {
		return uuid.SliceToUUID(key.Column.Qualifier)
	}

	if bytes.Equal(key.ID, Relationship) {
		to := widecolumnstore.Unmarshal(s.Value)
		return to.(*uuid.UUID)
	}

	return nil
}

// UUID looks for the UUID in the KeyValue
func UUID(s *widecolumnstore.KeyValue) *uuid.UUID {
	key := &widecolumnstore.Key{}
	key.Unmarshal(s.Key)

	if bytes.Equal(key.Column.Family, Label) {
		return uuid.SliceToUUID(key.ID)
	}

	if bytes.Equal(key.Column.Family, Properties) {
		return uuid.SliceToUUID(key.ID)
	}

	if bytes.Equal(key.Column.Family, Relationship) {
		return uuid.SliceToUUID(key.ID)
	}

	if bytes.Equal(key.Column.Family, Relationshipproperties) {
		return uuid.SliceToUUID(key.ID)
	}

	// Transpose

	if bytes.Equal(key.ID, TLabel) {
		return uuid.SliceToUUID(key.Column.Qualifier)
	}

	if bytes.Equal(key.ID, TProperties) {
		return uuid.SliceToUUID(key.Column.Qualifier)
	}

	if bytes.Equal(key.ID, TRelationship) {
		return uuid.SliceToUUID(key.Column.Qualifier)
	}

	if bytes.Equal(key.ID, TRelationshipproperties) {
		return uuid.SliceToUUID(key.Column.Qualifier)
	}

	return nil
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
