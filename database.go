package main

import (
	"errors"
	"log"

	bolt "go.etcd.io/bbolt"
)

func dbOpen(dbname string) {
	var err error
	DB, err = bolt.Open(dbname, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func dbTx() *bolt.Tx {
	// start transaction
	tx, err := DB.Begin(true)
	if err != nil {
		log.Fatal(err)
	}
	return tx
}

func dbBucketCreate(transaction *bolt.Tx, bucketname string) *bolt.Bucket {
	// create bucket
	bucket, err := transaction.CreateBucketIfNotExists([]byte(bucketname))
	if err != nil {
		log.Fatal(err)
	}
	return bucket
}

func dbInsertKVPair(bucket *bolt.Bucket, key string, data []byte) bool {
	// write entry
	err := bucket.Put([]byte(key), data)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func dbCommit(transaction *bolt.Tx) {
	// commit transaction
	err := transaction.Commit()
	if err != nil {
		log.Fatal(err)
	}
}

func dbInsert(bucketname string, key string, data []byte) error {
	transaction := dbTx()
	defer transaction.Rollback()
	bucket := dbBucketCreate(transaction, bucketname)
	success := dbInsertKVPair(bucket, key, data)
	if !success {
		return errors.New("inserting data into DB failed")
	}
	dbCommit(transaction)
	return nil
}

// WriteDB is a wrapper for opening the DB and writing some test data
func WriteDB(dbname string, bucketname string, key string, data []byte) {
	dbOpen(dbname)

	err := dbInsert(bucketname, key, data)
	if err != nil {
		log.Fatal(err)
	}
}
