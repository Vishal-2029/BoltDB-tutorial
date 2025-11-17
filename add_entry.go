package main

import (
	"fmt"
	"time"

	bolt "go.etcd.io/bbolt"
)

func addEntry(db *bolt.DB, title string, content string) error {
	return db.Update(func(tx *bolt.Tx) error {
		// Get the entries bucket
		b := tx.Bucket([]byte("entries"))
		if b == nil {
			return fmt.Errorf("bucket doesn't exist")
		}

		// Create a unique ID using timestamp
		id := fmt.Sprintf("%d", time.Now().UnixNano())

		// Create our entry data
		entry := fmt.Sprintf("Title: %s\nContent: %s\nDate: %s",
			title, content, time.Now().Format("2006-01-02 15:04:05"))

		// Store in database: key = id, value = entry
		err := b.Put([]byte(id), []byte(entry))
		if err != nil {
			return err
		}

		fmt.Printf("Entry added successfully! ID: %s\n", id)
		return nil
	})
}
