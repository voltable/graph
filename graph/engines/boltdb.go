package engines

import (
	"log"
	"time"

	"github.com/boltdb/bolt"
)

const (
	bucketGraph = "graph"
	bucketLabel = "label"
	bucketIndex = "index"
)

type boltdb struct {
	bolt *bolt.DB
}

func (e boltdb) CreateIndexes() {

	// create the bucket to hold the Adjacency list.
	e.bolt.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketGraph))
		if err != nil {
			log.Fatal(err)
			panic("Failed to create graph")
		}

		_, err = tx.CreateBucketIfNotExists([]byte(bucketIndex))
		if err != nil {
			log.Fatal(err)
			panic("Failed to create index")
		}

		return nil
	})
}

func (e boltdb) Close() {
	defer e.bolt.Close()
}

func (e boltdb) Open(path string) error {
	b, err := bolt.Open(path+".db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	e.bolt = b
	return err
}
