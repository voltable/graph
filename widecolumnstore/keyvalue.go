package widecolumnstore

import (
	"bytes"

	"github.com/RossMerr/Caudex.Graph/arch"
	"github.com/RossMerr/Caudex.Graph/uuid"
)

func (s *KeyValue) Weight() float64 {
	key := &Key{}
	key.Unmarshal(s.Key)

	if bytes.Equal(key.Column.Family, Relationship) {
		to := Unmarshal(s.Value)
		return to.(float64)
	}

	if bytes.Equal(key.ID, Relationship) {
		return arch.DecodeFloat64Bytes(key.Column.Extended).(float64)
	}

	return 0
}

func (s *KeyValue) To() *uuid.UUID {
	key := &Key{}
	key.Unmarshal(s.Key)

	if bytes.Equal(key.Column.Family, Relationship) {
		return uuid.SliceToUUID(key.Column.Qualifier)
	}

	if bytes.Equal(key.ID, Relationship) {
		to := Unmarshal(s.Value)
		return to.(*uuid.UUID)
	}

	return nil
}

// UUID looks for the UUID in the KeyValue
func (s *KeyValue) UUID() *uuid.UUID {
	key := &Key{}
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

func (s *KeyValue) Interpret(value string) interface{} {
	key := &Key{}
	key.Unmarshal(s.Key)

	if bytes.Equal(key.Column.Family, Properties) {

		if value == string(key.Column.Qualifier) {
			return Unmarshal(s.Value)
		}
	}

	return nil
}
