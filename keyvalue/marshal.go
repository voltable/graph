package keyvalue

// MarshalKeyValue interface for marshaling objects as a keyValue
type MarshalKeyValue interface {
	MarshalKeyValue() []*KeyValue
	MarshalKeyValueTranspose() []*KeyValue

	UnmarshalKeyValue(kv []*KeyValue)
	UnmarshalKeyValueTranspose(kv []*KeyValue)
}
