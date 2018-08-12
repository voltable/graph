package keyvalue

import "github.com/RossMerr/Caudex.Graph/uuid"

// RelationshipPrefix builds up the key for a prefix search to find any relationship
func RelationshipPrefix(id *uuid.UUID) []byte {
	return append(id[:], append(US, Relationship...)...)
}
