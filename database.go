package main

import (
	"errors"
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

func dbTx(db *bolt.DB) *bolt.Tx {
	// start transaction
	tx, err := db.Begin(true)
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

func dbInsert(db *bolt.DB, bucketname string, key string, data []byte) error {
	transaction := dbTx(db)
	defer transaction.Rollback()
	bucket := dbBucketCreate(transaction, bucketname)
	success := dbInsertKVPair(bucket, key, data)
	if !success {
		return errors.New("inserting data into DB failed")
	}
	dbCommit(transaction)
	return nil
}

func WriteDB(dbname string, bucketname string, key string, data []byte) {
	db := dbOpen(dbname)
	defer db.Close()
	err := dbInsert(db, bucketname, key, data)
	if err != nil {
		log.Fatal(err)
	}
}
