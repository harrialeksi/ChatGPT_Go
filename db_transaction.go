/*
In this example, we first open a connection to the PostgreSQL database using the sql.Open function from the database/sql package. We then start a new 
transaction using the Begin method on the *sql.DB connection object.

Within the transaction, we perform two database operations: an INSERT into a users table and an UPDATE to an accounts table. If either operation fails, 
we roll back the transaction using the Rollback method on the transaction object. If both operations are successful, we commit the transaction using the 
Commit method.

Note that we also use the defer statement to ensure that the database connection is closed when the program exits.
*/

package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/lib/pq"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "mypassword"
    dbname   = "mydb"
)

func main() {
    // Open a connection to the database
    db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname))
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Start a transaction
    tx, err := db.Begin()
    if err != nil {
        log.Fatal(err)
    }

    // Perform some database operations within the transaction
    _, err = tx.Exec("INSERT INTO users (name, email) VALUES ($1, $2)", "John Doe", "john@example.com")
    if err != nil {
        // Roll back the transaction if there is an error
        tx.Rollback()
        log.Fatal(err)
    }

    _, err = tx.Exec("UPDATE accounts SET balance = balance + $1 WHERE user_id = $2", 100, 1)
    if err != nil {
        // Roll back the transaction if there is an error
        tx.Rollback()
        log.Fatal(err)
    }

    // Commit the transaction if all operations are successful
    if err := tx.Commit(); err != nil {
        log.Fatal(err)
    }

    log.Println("Transaction completed successfully")
}

