# Personal Diary App - BoltDB Tutorial

A simple command-line personal diary application built with Go and BoltDB to demonstrate CRUD operations with an embedded key-value database.

## 📚 What is BoltDB?

BoltDB is a pure Go embedded key-value database. It's perfect for projects that need a simple, fast database without the overhead of a separate database server.

**Key Features:**
- No external dependencies
- ACID transactions
- Simple API
- Single file database
- No configuration needed

## 🚀 Getting Started

### Prerequisites

- Go 1.16 or higher
- Basic understanding of Go programming

### Installation

1. Clone or download this project
2. Install BoltDB dependency:
```bash
go get go.etcd.io/bbolt
```

3. Run the application:
```bash
go run complete_app.go
```

## 📁 Project Structure

```
Bolt-DB/
├── example/
│   └── complete_app.go     # Full interactive BoltDB app
│
├── add_entry.go            # Standalone example: Add entries
├── basic.go                # Standalone example: Basic BoltDB flow
├── delete_entries.go       # Standalone example: Delete entries
├── read_entries.go         # Standalone example: Read entries
├── main.go                 # Demo runner or combined example
│
├── diary.db                # Auto-generated database file
├── go.mod                  # Module definition
├── go.sum                  # Dependencies
└── README.md               # Documentation

```

## 🎯 Learning Path

Follow these files in order to learn BoltDB step by step:

### 1. **basic.go** - Database Basics
Learn how to:
- Open a BoltDB database
- Create buckets (like tables in SQL)
- Use proper error handling

```bash
go run basic.go
```

### 2. **add_entry.go** - Create Operation
Learn how to:
- Insert data into the database
- Generate unique IDs
- Format and store structured data

### 3. **read_entries.go** - Read Operations
Learn how to:
- Retrieve all entries using cursors
- Get specific entries by ID
- Iterate through database records

### 4. **delete_entries.go** - Delete Operation
Learn how to:
- Remove entries from the database
- Check if entries exist before deletion
- Handle not-found scenarios

### 5. **main.go** - Complete Demo
See all operations working together:
```bash
go run main.go
```

### 6. **complete_app.go** - Full Application
A complete interactive CLI application with a menu system:
```bash
go run complete_app.go
```

## 🎮 Using the Complete App

When you run `complete_app.go`, you'll see this menu:

```
=== PERSONAL DIARY APP ===
1. Add new entry
2. View all entries
3. Find entry by ID
4. Delete entry
5. Exit
```

**Example Usage:**

1. **Add a new entry:**
   - Choose option 1
   - Enter a title
   - Enter your diary content
   - Entry is saved with a unique ID

2. **View all entries:**
   - Choose option 2
   - See all your diary entries with their IDs

3. **Find specific entry:**
   - Choose option 3
   - Enter the entry ID (from the list)
   - View that specific entry

4. **Delete an entry:**
   - Choose option 4
   - Enter the entry ID
   - Confirm deletion

## 🔑 Key Concepts Learned

### 1. **Buckets**
- Buckets are like tables in SQL databases
- Store related data together
- Create with `CreateBucketIfNotExists()`

### 2. **Transactions**
- **Read transactions:** `db.View()` - for reading data
- **Write transactions:** `db.Update()` - for modifying data
- Automatic rollback on error

### 3. **Keys and Values**
- Both keys and values are byte slices (`[]byte`)
- Convert strings with `[]byte(string)`
- Convert back with `string([]byte)`

### 4. **Cursors**
- Used to iterate through all records
- Methods: `First()`, `Next()`, `Last()`, `Prev()`

## 📝 Code Examples

### Opening a Database
```go
db, err := bolt.Open("diary.db", 0600, nil)
if err != nil {
    log.Fatal(err)
}
defer db.Close()
```

### Creating a Bucket
```go
err = db.Update(func(tx *bolt.Tx) error {
    _, err := tx.CreateBucketIfNotExists([]byte("entries"))
    return err
})
```

### Adding Data
```go
err := db.Update(func(tx *bolt.Tx) error {
    b := tx.Bucket([]byte("entries"))
    return b.Put([]byte("key"), []byte("value"))
})
```

### Reading Data
```go
err := db.View(func(tx *bolt.Tx) error {
    b := tx.Bucket([]byte("entries"))
    value := b.Get([]byte("key"))
    fmt.Println(string(value))
    return nil
})
```

### Iterating Through All Entries
```go
err := db.View(func(tx *bolt.Tx) error {
    b := tx.Bucket([]byte("entries"))
    cursor := b.Cursor()
    
    for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
        fmt.Printf("Key: %s, Value: %s\n", k, v)
    }
    return nil
})
```

## 🛠️ Customization Ideas

Try extending this project:

1. **Add search functionality** - Search entries by keyword
2. **Add tags** - Categorize your entries
3. **Export to text file** - Backup your diary
4. **Add password protection** - Encrypt your entries
5. **Add date filtering** - View entries from specific dates
6. **Statistics** - Count total entries, entries per month, etc.

## ⚠️ Important Notes

- The `diary.db` file contains all your data
- Keep backups of your database file
- BoltDB allows only one write transaction at a time
- Always use `defer db.Close()` to close the database properly

## 🐛 Troubleshooting

**Error: "bucket doesn't exist"**
- Make sure to create the bucket before using it
- Run the initialization code first

**Error: "timeout"**
- Only one process can open the database at a time
- Close other instances of your app

**Database file locked**
- Make sure no other program is using the database
- Check if previous program instance closed properly

## 📚 Further Learning

- [BoltDB GitHub](https://github.com/etcd-io/bbolt)
- [BoltDB Documentation](https://pkg.go.dev/go.etcd.io/bbolt)
- Try building other apps: TODO list, contact manager, note-taking app

## 📄 License

This is a tutorial project for learning purposes. Feel free to use and modify as needed.

## 🤝 Contributing

This is a learning project, but suggestions for improvements are welcome!

---

**Happy Coding! 🎉**

Start with `basic.go` and work your way through each file to master BoltDB!
