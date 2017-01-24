package operations

import (
	"encoding/json"

	"bitbucket.org/rossmerr/caudex/graphs"
	"bitbucket.org/rossmerr/caudex/graphs/boltdb/internal"
	"github.com/boltdb/bolt"
)

func Add(db *bolt.DB, c *[]graphs.Vertex) error {
	var err error
	var buf []byte
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(internal.BucketGraph))
		for _, vertex := range *c {
			if buf, err = json.Marshal(vertex); err != nil {
				b.Put([]byte(vertex.ID), buf)
			} else {
				return err
			}
		}
		return nil
	})
}
