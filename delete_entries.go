package main

import (
	"fmt"

	bolt "go.etcd.io/bbolt"
)

func deleteEntry(db *bolt.DB, id string) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("entries"))
		if b == nil {
			return fmt.Errorf("bucket doesn't exist")
		}

		// Check if entry exists before deleting
		existing := b.Get([]byte(id))
		if existing == nil {
			return fmt.Errorf("entry with ID %s not found", id)
		}

		// Delete the entry
		err := b.Delete([]byte(id))
		if err != nil {
			return err
		}

		fmt.Printf("Entry %s deleted successfully!\n", id)
		return nil
	})
}
