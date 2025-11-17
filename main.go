package main

import (
	"fmt"
	"log"

	bolt "go.etcd.io/bbolt"
)

func main() {
	fmt.Println("Welcome to Personal Diary App!")

	// Open database
	db, err := bolt.Open("diary.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create bucket first
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("entries"))
		return err
	})
	if err != nil {
		log.Fatal(err)
	}

	// Test all operations
	fmt.Println("\n=== TESTING ALL OPERATIONS ===")

	// 1. Add entries
	fmt.Println("\n1. Adding entries:")
	addEntry(db, "My First Day", "Today I started learning BoltDB. It's amazing!")
	addEntry(db, "Go Programming", "I love how simple Go makes database operations.")
	addEntry(db, "My Diary App", "Building this diary app step by step!")

	// 2. Read all entries
	fmt.Println("\n2. Reading all entries:")
	listAllEntries(db)

	// 3. Demo: Get a specific entry (you'll see the IDs in the output above)
	fmt.Println("\n3. Demo - Try getting an entry by ID:")
	// Uncomment the line below and replace with actual ID from your output
	// getEntry(db, "1735762832123456789") // Replace with actual ID

	// 4. Demo: Delete an entry
	fmt.Println("\n4. Demo - Try deleting an entry:")
	// Uncomment the line below and replace with actual ID from your output
	// deleteEntry(db, "1735762832123456789") // Replace with actual ID

	// 5. Show final entries
	fmt.Println("\n5. Final entries list:")
	listAllEntries(db)
}
