package keyvalue

import "github.com/RossMerr/Caudex.Graph/uuid"

// RelationshipPrefix builds up the key for a prefix search to find any relationship
func RelationshipPrefix(id *uuid.UUID) []byte {
	var bytes = id[:]
	bytes = append(bytes, US...)
	bytes = append(bytes, Relationship...)
	return bytes
}
