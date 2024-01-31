package main

import (
	"testing"
)

type dbtest struct {
	dbname     string
	bucketname string
	key        string
	data       []byte
}

var dbtestList = []dbtest{
	dbtest{"./test/testdb.db", "testbucket", "td", []byte("testdata")},
	dbtest{"./test/testdb.db", "testbucket", "tj", []byte(`{"birds":{"pigeon":"likes to perch on rocks","eagle":"bird of prey"`)},
}

func TestWriteDB(t *testing.T) {
	for _, item := range dbtestList {
		WriteDB(item.dbname, item.bucketname, item.key, item.data)
	}
}
