package operations

import (
	"encoding/json"

	"bitbucket.org/rossmerr/caudex/graphs"
	"bitbucket.org/rossmerr/caudex/graphs/boltdb/internal"
	"github.com/boltdb/bolt"
)

func Find(db *bolt.DB, ID string) (*graphs.Vertex, error) {
	var err error
	var buf []byte
	var v graphs.Vertex
	return &v, db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(internal.BucketGraph))
		buf = b.Get([]byte(ID))

		if err = json.Unmarshal(buf, v); err == nil {
			return nil
		}

		return err
	})
}
