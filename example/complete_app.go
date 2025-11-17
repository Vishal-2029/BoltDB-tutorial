package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	bolt "go.etcd.io/bbolt"
)

func main() {
	// Open database
	db, err := bolt.Open("diary.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize bucket
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("entries"))
		return err
	})
	if err != nil {
		log.Fatal(err)
	}

	// Interactive menu
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n=== PERSONAL DIARY APP ===")
		fmt.Println("1. Add new entry")
		fmt.Println("2. View all entries")
		fmt.Println("3. Find entry by ID")
		fmt.Println("4. Delete entry")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option (1-5): ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			addNewEntry(db, scanner)
		case "2":
			listAllEntries(db)
		case "3":
			findEntry(db, scanner)
		case "4":
			deleteEntryInteractive(db, scanner)
		case "5":
			fmt.Println("Goodbye! Your diary is saved.")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func addNewEntry(db *bolt.DB, scanner *bufio.Scanner) {
	fmt.Print("Enter title: ")
	scanner.Scan()
	title := scanner.Text()

	fmt.Print("Enter content: ")
	scanner.Scan()
	content := scanner.Text()

	err := addEntry(db, title, content)
	if err != nil {
		fmt.Printf("Error adding entry: %v\n", err)
	}
}

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

func findEntry(db *bolt.DB, scanner *bufio.Scanner) {
	fmt.Print("Enter entry ID: ")
	scanner.Scan()
	id := scanner.Text()

	err := getEntry(db, id)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

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

func deleteEntryInteractive(db *bolt.DB, scanner *bufio.Scanner) {
	fmt.Print("Enter entry ID to delete: ")
	scanner.Scan()
	id := scanner.Text()

	// Show the entry before deleting
	fmt.Println("Entry to delete:")
	err := getEntry(db, id)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Print("Are you sure you want to delete this entry? (y/n): ")
	scanner.Scan()
	confirm := strings.ToLower(scanner.Text())

	if confirm == "y" || confirm == "yes" {
		err = deleteEntry(db, id)
		if err != nil {
			fmt.Printf("Error deleting entry: %v\n", err)
		}
	} else {
		fmt.Println("Deletion cancelled.")
	}
}

// Enhanced addEntry function
func addEntry(db *bolt.DB, title string, content string) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("entries"))
		if b == nil {
			return fmt.Errorf("bucket doesn't exist")
		}

		id := fmt.Sprintf("%d", time.Now().UnixNano())

		entry := fmt.Sprintf("Title: %s\nContent: %s\nDate: %s",
			title, content, time.Now().Format("2006-01-02 15:04:05"))

		err := b.Put([]byte(id), []byte(entry))
		if err != nil {
			return err
		}

		fmt.Printf("✓ Entry added successfully! ID: %s\n", id)
		return nil
	})
}
