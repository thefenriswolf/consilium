package main

import (
	bolt "go.etcd.io/bbolt"
	"log"
)

func dbOpen(dbname string) *bolt.DB {
	db, err := bolt.Open(dbname, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func dbInsert(db *bolt.DB, bucketname string, key string, data []byte) {
	// start transaction
	tx, err := db.Begin(true)
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()
	// create bucket
	bucket, err := tx.CreateBucketIfNotExists([]byte(bucketname))
	if err != nil {
		log.Fatal(err)
	}
	// write entry
	err = bucket.Put([]byte(key), data)
	if err != nil {
		log.Fatal(err)
	}
	// commit transaction
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}

func WriteDB(dbname string, bucketname string, key string, data []byte) {
	db := dbOpen(dbname)
	defer db.Close()
	dbInsert(db, bucketname, key, data)
}
