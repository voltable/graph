package operations

import (
	"bitbucket.org/rossmerr/caudex/graphs"
	"bitbucket.org/rossmerr/caudex/graphs/boltdb/internal"
	"github.com/boltdb/bolt"
)

func Delete(db *bolt.DB, c *[]graphs.Vertex) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(internal.BucketGraph))
		for _, vertex := range *c {
			b.Delete([]byte(vertex.ID))
		}
		return nil
	})
}
