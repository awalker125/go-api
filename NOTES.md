
Here it talks about wrapping in a struct <https://go.dev/doc/tutorial/database-access> instead of a global.

Bing reckons this is how you would do it

```go
package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    "github.com/go-sql-driver/mysql"
)

// DB wraps the sql.DB connection pool.
type DB struct {
    conn *sql.DB
}

// NewDB initializes a new database connection and returns a DB struct.
func NewDB(cfg mysql.Config) (*DB, error) {
    db, err := sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        return nil, err
    }

    if err := db.Ping(); err != nil {
        return nil, err
    }

    return &DB{conn: db}, nil
}

func main() {
    // Capture connection properties.
    cfg := mysql.Config{
        User:   os.Getenv("DBUSER"),
        Passwd: os.Getenv("DBPASS"),
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "recordings",
    }

    // Initialize the database connection.
    db, err := NewDB(cfg)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Connected!")
    // You can now use db.conn to interact with the database.
}
```

Explanation on "*" and "&". "*" means we expect a pointer. "&" creates a pointer.

```go
package main

import "fmt"

type Person struct {
    Name string
}

func main() {
    p := Person{Name: "Alice"}
    fmt.Println("Before:", p.Name)

    changeName(&p)
    fmt.Println("After:", p.Name)
}

func changeName(p *Person) {
    p.Name = "Bob"
}
```

In this example, changeName takes a pointer to a Person struct and modifies the Name field. The changes are reflected in the original Person instance because p is a pointer.
