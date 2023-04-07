package infrastructure

import (
	"encoding/binary"
	"log"
	"path/filepath"
	"time"

	"github.com/boltdb/bolt"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/oivinig/cli-todoer/internal/app/interfaces"
)

type DBHandler struct {
	Conn *bolt.DB
}
var bucketName = []byte("tasks")

func Connect() (interfaces.DBHandler, error) {
	dbHandler := &DBHandler{}
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	
	db, err := bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second}) 
	if err != nil {
		log.Fatal(err)
	}
	dbHandler.Conn = db
	return dbHandler, err
}

func (db *DBHandler) Create(value string) error {
	var id int
	return db.Conn.Update(func (tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(bucketName)
		if err != nil {
			return err
		}

		id64, _ := bucket.NextSequence()
		id = int(id64)
		key := itob(id)

		err = bucket.Put(key, []byte(value))
		if err != nil {
			return err
		}
		return nil
	})
}

func (db *DBHandler) Retrieve() ([][]byte, [][]byte, error) {
	var keys [][]byte
	var values [][]byte
	err := db.Conn.View(func (tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		c := b.Cursor()
		for k, v  := c.First(); k != nil; k, v = c.Next() {
			keys = append(keys, k)
			values = append(values, v)
		}	
		return nil
	})
	if err != nil {
		return nil, nil, err
	}
	return keys, values, nil
}

func (db *DBHandler) Delete(key int) error {
	return db.Conn.Update(func (tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		return b.Delete(itob(key))
	})
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
