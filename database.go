package main

import (
	bolt "go.etcd.io/bbolt"
	"log"
)

func dbOpen() *bolt.DB {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func dbBucket(bucket string, db *bolt.DB) {
	// start transaction
	tx, err := db.Begin(true)
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()
	// create bucket
	_, err = tx.CreateBucketIfNotExists([]byte(bucket))
	if err != nil {
		log.Fatal(err)
	}
	// commit transaction
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}

func WriteDB() {
	db := dbOpen()
	defer db.Close()
	dbBucket("testbucket3", db)
}
