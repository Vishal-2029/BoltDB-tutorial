package main

import (
	"fmt"

	bolt "go.etcd.io/bbolt"
)

func listAllEntries(db *bolt.DB) error {
	return db.View(func(tx *bolt.Tx) error {
		// Get the bucket
		b := tx.Bucket([]byte("entries"))
		if b == nil {
			return fmt.Errorf("bucket doesn't exist")
		}

		fmt.Println("\n=== ALL DIARY ENTRIES ===")

		// Create a cursor to iterate through all entries
		cursor := b.Cursor()
		count := 0

		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			count++
			fmt.Printf("\n--- Entry %d (ID: %s) ---\n", count, string(k))
			fmt.Println(string(v))
			fmt.Println("----------------------")
		}

		if count == 0 {
			fmt.Println("No entries found. Start writing your diary!")
		} else {
			fmt.Printf("Total entries: %d\n", count)
		}

		return nil
	})
}

func getEntry(db *bolt.DB, id string) error {
	return db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("entries"))
		if b == nil {
			return fmt.Errorf("bucket doesn't exist")
		}

		value := b.Get([]byte(id))
		if value == nil {
			return fmt.Errorf("entry with ID %s not found", id)
		}

		fmt.Printf("\n--- Entry (ID: %s) ---\n", id)
		fmt.Println(string(value))
		fmt.Println("----------------------")

		return nil
	})
}
