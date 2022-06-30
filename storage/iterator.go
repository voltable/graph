package storage

// Iterator is an alias for function to iterate over data.
type Iterator func() (SparseValue, bool)
