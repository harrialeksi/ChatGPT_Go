Sure, here are some examples of how to perform CRUD operations (CREATE, READ, UPDATE, DELETE) in PostgreSQL using pgx library in Go:

Assuming you have already connected to the PostgreSQL database using pgx and have a connection conn:

CREATE
```go
func createUser(name string, email string, age int) error {
    // Prepare the SQL statement
    stmt := "INSERT INTO users(name, email, age) VALUES($1, $2, $3)"

    // Execute the SQL statement with the parameters
    _, err := conn.Exec(context.Background(), stmt, name, email, age)
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
    stmt := "SELECT id, name, email, age FROM users WHERE id=$1"

    // Query the database with the parameter
    row := conn.QueryRow(context.Background(), stmt, id)
    err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Age)
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
    stmt := "UPDATE users SET name=$1, email=$2, age=$3 WHERE id=$4"

    // Execute the SQL statement with the parameters
    _, err := conn.Exec(context.Background(), stmt, name, email, age, id)
    if err != nil {
        return err
    }

    return nil
}
```

DELETE
```go
func deleteUser(id int) error {
    // Prepare the SQL statement
    stmt := "DELETE FROM users WHERE id=$1"

    // Execute the SQL statement with the parameter
    _, err := conn.Exec(context.Background(), stmt, id)
    if err != nil {
        return err
    }

    return nil
}
```

Note that in these examples, User is a struct with fields ID, Name, Email, and Age, corresponding to the columns in the users table in the database. 
Also, error handling has been omitted for brevity. In a real application, you should handle errors appropriately.
