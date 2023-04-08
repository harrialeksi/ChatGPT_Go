/*
In this example, we import the github.com/robfig/cron package and create a new cron scheduler using cron.New(). We then schedule a task to 
run every minute using the c.AddFunc() method, which takes a cron expression and a function to execute.

The cron expression * * * * * specifies that the task should run every minute, and the function simply prints the current time to the console.

We then start the cron scheduler using c.Start() and let it run for 5 minutes using time.Sleep(5 * time.Minute). Finally, we stop the cron 
scheduler using c.Stop().
*/

package main

import (
    "fmt"
    "time"

    "github.com/robfig/cron"
)

func main() {
    // Create a new cron scheduler
    c := cron.New()

    // Schedule a task to run every minute
    c.AddFunc("* * * * *", func() {
        fmt.Println("Task executed at:", time.Now())
    })

    // Start the cron scheduler
    c.Start()

    // Run the scheduler for 5 minutes
    time.Sleep(5 * time.Minute)

    // Stop the cron scheduler
    c.Stop()
}
