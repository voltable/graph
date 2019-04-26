package query

import "github.com/voltable/graph/uuid"

// Iterator is an alias for function to iterate over data.
type Iterator func() (uuid.UUID, float64, bool)
