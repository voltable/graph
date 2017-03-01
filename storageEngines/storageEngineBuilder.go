package storageEngines

import (
	"github.com/RossMerr/Caudex.Graph"
	"github.com/RossMerr/Caudex.Graph/storageEngines/boltdb"
	"github.com/RossMerr/Caudex.Graph/storageEngines/memorydb"
)

// StorageEngineType the backend persistence storage engine
type StorageEngineType int

const (
	// Bolt use Bolt as the storage engine (Default)
	Bolt StorageEngineType = iota
	// Memory use in memory for the storage engine
	Memory StorageEngineType = iota
)

func BuildGraphDefault(o *graphs.Options) graphs.StorageEngine {
	return BuildStorageEngine(Bolt, o)
}

func BuildStorageEngine(e StorageEngineType, o *graphs.Options) graphs.StorageEngine {
	if e == Memory {
		return memorydb.NewStorageEngine(o)
	}
	return boltdb.NewStorageEngine(o)
}
