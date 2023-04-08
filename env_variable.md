show environment variables example using joho/godotenv

Sure, here's an example of how to use the joho/godotenv package in Go to load environment variables from a .env file:

First, install the joho/godotenv package by running go get github.com/joho/godotenv.

Create a .env file in the root of your project with the environment variables you want to use. For example:

```
DB_HOST=localhost
DB_PORT=5432
DB_NAME=mydatabase
DB_USER=myusername
DB_PASSWORD=mypassword
```

In your Go code, import the github.com/joho/godotenv package and call the godotenv.Load() function to load the environment variables from the .env file:

```go
package main

import (
    "fmt"
    "os"

    "github.com/joho/godotenv"
)

func main() {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        fmt.Println("Error loading .env file:", err)
        return
    }

    // Get environment variables
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")

    // Use environment variables
    fmt.Printf("DB_HOST: %s\n", dbHost)
    fmt.Printf("DB_PORT: %s\n", dbPort)
    fmt.Printf("DB_NAME: %s\n", dbName)
    fmt.Printf("DB_USER: %s\n", dbUser)
    fmt.Printf("DB_PASSWORD: %s\n", dbPassword)
}
```

Run the Go code and you should see the environment variables printed to the console.

Note that the godotenv.Load() function only needs to be called once, at the start of your program. The os.Getenv() function is used to retrieve the 
values of individual environment variables. If an environment variable is not set, os.Getenv() will return an empty string.
