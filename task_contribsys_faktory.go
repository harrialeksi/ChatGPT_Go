/*
In this example, we import the github.com/contribsys/faktory/client and github.com/contribsys/faktory/job packages.

We create a new Faktory client using client.Open() and register a job type using faktoryClient.Register().

The job type takes a context.Context and a variadic list of arguments, and simply prints "Hello, world!" to the console.

We then enqueue a job using job.New() and faktoryClient.Push(). The job type and arguments are specified in the job.New() call.

Note that you will need to have a Faktory server running locally or remotely for this example to work.
*/

package main

import (
    "context"
    "fmt"

    "github.com/contribsys/faktory/client"
    "github.com/contribsys/faktory/job"
)

func main() {
    // Create a new Faktory client
    faktoryClient, err := client.Open()
    if err != nil {
        panic(err)
    }
    defer faktoryClient.Close()

    // Register a job type
    faktoryClient.Register("hello", func(ctx context.Context, args ...interface{}) error {
        fmt.Println("Hello, world!")
        return nil
    })

    // Enqueue a job
    job := job.New("hello", nil)
    _, err = faktoryClient.Push(job)
    if err != nil {
        panic(err)
    }
}
