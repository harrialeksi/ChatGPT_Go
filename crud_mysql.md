show CRUD (CREATE, READ, UPDATE, DELETE) examples in mysql using go language

Sure, here are some examples of how to perform CRUD operations (CREATE, READ, UPDATE, DELETE) in MySQL using Go language:

Assuming you have already connected to the MySQL database using a package like database/sql and have a connection db:

CREATE
```go
func createUser(name string, email string, age int) error {
    // Prepare the SQL statement
    stmt, err := db.Prepare("INSERT INTO users(name, email, age) VALUES(?, ?, ?)")
    if err != nil {
        return err
    }
    defer stmt.Close()

    // Execute the SQL statement with the parameters
    _, err = stmt.Exec(name, email, age)
    if err != nil {
        return err
    }

    return nil
}
```

READ
```go
func getUser(id int) (User, error) {
    var user User

    // Prepare the SQL statement
    stmt, err := db.Prepare("SELECT id, name, email, age FROM users WHERE id=?")
    if err != nil {
        return user, err
    }
    defer stmt.Close()

    // Query the database with the parameter
    err = stmt.QueryRow(id).Scan(&user.ID, &user.Name, &user.Email, &user.Age)
    if err != nil {
        return user, err
    }

    return user, nil
}
```

UPDATE
```go
func updateUser(id int, name string, email string, age int) error {
    // Prepare the SQL statement
    stmt, err := db.Prepare("UPDATE users SET name=?, email=?, age=? WHERE id=?")
    if err != nil {
        return err
    }
    defer stmt.Close()

    // Execute the SQL statement with the parameters
    _, err = stmt.Exec(name, email, age, id)
    if err != nil {
        return err
    }

    return nil
}
```

DELETE
```
func deleteUser(id int) error {
    // Prepare the SQL statement
    stmt, err := db.Prepare("DELETE FROM users WHERE id=?")
    if err != nil {
        return err
    }
    defer stmt.Close()

    // Execute the SQL statement with the parameter
    _, err = stmt.Exec(id)
    if err != nil {
        return err
    }

    return nil
}
```

Note that in these examples, User is a struct with fields ID, Name, Email, and Age, corresponding to the columns in the users table in the database. 
Also, error handling has been omitted for brevity. In a real application, you should handle errors appropriately.
