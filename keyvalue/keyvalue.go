package keyvalue

// NewKeyValue returns a new KeyValue
func NewKeyValue(i interface{}, bytes ...[]byte) *KeyValue {
	kv := &KeyValue{
		Value: NewAny(i),
	}
	for _, b := range bytes {
		kv.Key = append(kv.Key, b...)
	}
	return kv
}
