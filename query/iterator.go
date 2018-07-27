package query

import "github.com/RossMerr/Caudex.Graph/uuid"

// Iterator is an alias for function to iterate over data.
type Iterator func() (interface{}, bool)

// IteratorUUID is an alias for function to iterate over uuid's.
type IteratorUUID func() (uuid.UUID, bool)
