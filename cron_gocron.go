/*
In this example, we import the github.com/go-co-op/gocron package and create a new scheduler using gocron.NewScheduler(time.UTC).

We then schedule a task to run every minute using the s.Every(1).Minute().Do() method, which takes an integer representing the frequency and 
a function to execute.

The function simply prints the current time to the console.

We then start the scheduler using s.StartAsync() and let it run for 5 minutes using time.Sleep(5 * time.Minute). Finally, we stop the scheduler 
using s.Stop().
*/

package main

import (
    "fmt"
    "time"

    "github.com/go-co-op/gocron"
)

func main() {
    // Create a new gocron scheduler
    s := gocron.NewScheduler(time.UTC)

    // Schedule a task to run every minute
    s.Every(1).Minute().Do(func() {
        fmt.Println("Task executed at:", time.Now())
    })

    // Start the scheduler
    s.StartAsync()

    // Run the scheduler for 5 minutes
    time.Sleep(5 * time.Minute)

    // Stop the scheduler
    s.Stop()
}
