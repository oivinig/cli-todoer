package infrastructure

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"github.com/oivinig/cli-todoer/internal/app/interfaces"
)

type DBHandler struct {
	Conn *bolt.DB
}

var bucketName = []byte("TASKS") //TODO: create logic to receive the bucket name and use each name for each statuses

func Connect() (interfaces.DBHandler, error) {
	dbHandler := &DBHandler{}
	db, err := bolt.Open("/Users/vghffmnn/development/cli-todoer/bolt.db", 0644, nil) //FIXME: implement to get the path from env-var
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	dbHandler.Conn = db
	return dbHandler, err
}

func (db *DBHandler) Create(key []byte, value []byte) error {
	err := db.Conn.Update(func (tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(bucketName)
		if err != nil {
			return err
		}

		err = bucket.Put(key, value)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

// FIXME: change the logic to return an array of bytes (of all the values in the bucket)
func (db *DBHandler) Retrieve(key []byte) []byte {
	var val []byte
	err := db.Conn.View(func (tx *bolt.Tx) error {
		bucket := tx.Bucket(bucketName)
		if bucket == nil {
			return fmt.Errorf("bucket %q not found", bucketName)
		}

		val = bucket.Get(key)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return val
}


func (db *DBHandler) Update(key []byte, value []byte) {
	err := db.Conn.Update(func (tx *bolt.Tx) error {
		bucket := tx.Bucket(bucketName)
		if bucket == nil {
			return fmt.Errorf("bucket %q not found", bucketName)
		}
		
		err := bucket.Put(key, value)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}

func (db *DBHandler) Delete(key []byte) {
	err := db.Conn.Update(func (tx *bolt.Tx) error {
		bucket := tx.Bucket(bucketName)
		if bucket == nil {
			return fmt.Errorf("bucket %q not found", bucketName)
		}

		err := bucket.Delete(key)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
