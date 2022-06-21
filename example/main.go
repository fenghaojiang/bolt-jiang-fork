package main

import (
	"log"

	"github.com/boltdb/bolt"
)

func main() {
	db, err := bolt.Open("example/test.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
