package main

import (
	"fmt"
	"log"

	bolt "go.etcd.io/bbolt"
)

func basicOprations() {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	// 0600 => read & write permission
	// nil => use default options
	db, err := bolt.Open("diary.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Database opened successfully!")

	// Create a bucket
	err = db.Update(func(tx *bolt.Tx) error {
		// Create a bucket named "Entries" if it doesn't exist
		_, err := tx.CreateBucketIfNotExists([]byte("Entries"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Bucket created successfully!")

}
