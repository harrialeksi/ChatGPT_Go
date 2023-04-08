/*
show hibiken/asynq example in go language

In this example, we first create a Redis client using asynq.NewRedisClient. We pass in the address of the Redis server as well as any options we 
want to use.

We then create a new task using asynq.NewTask. We pass in the task type as the first argument, and a map of data as the second argument. The data 
can be any type that can be serialized as JSON.

We enqueue the task using asynq.Enqueue. This returns an asynq.TaskInfo object that contains information about the enqueued task, such as the task ID 
and enqueue time.

We then create a task handler using asynq.HandlerFunc. This is a function that will be called whenever a task of the specified type is dequeued. 
In this example, the task handler simply retrieves data from the task payload and logs a message.

We create a worker using asynq.NewServer. We pass in an asynq.RedisClientOpt object with the address of the Redis server, an asynq.Config object 
with the desired concurrency level, and the task handler we created earlier.

Finally, we start the worker using w.Run(). This will block the current goroutine and process tasks until an error occurs or the worker is stopped.
*/

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/hibiken/asynq"
)

func main() {
	// Create a Redis connection
	r := asynq.NewRedisClient("localhost:6379", asynq.RedisOptions{})

	// Create a task with some data
	task := asynq.NewTask("send_email", map[string]interface{}{
		"to":      "jane@example.com",
		"subject": "Hello, Jane!",
		"body":    "How are you doing?",
	})

	// Enqueue the task
	info, err := asynq.Enqueue(r, task)
	if err != nil {
		log.Fatalf("Could not enqueue task: %v", err)
	}
	fmt.Printf("Enqueued task %v\n", info.TaskID)

	// Create a task handler
	h := asynq.HandlerFunc(func(ctx context.Context, task *asynq.Task) error {
		// Retrieve data from the task payload
		to := task.Payload["to"].(string)
		subject := task.Payload["subject"].(string)
		body := task.Payload["body"].(string)

		// Send the email
		fmt.Printf("Sending email to %s with subject %s and body %s\n", to, subject, body)

		return nil
	})

	// Create a worker to handle the task
	w := asynq.NewServer(
		asynq.RedisClientOpt{Addr: "localhost:6379"},
		asynq.Config{Concurrency: 1},
		h,
	)

	// Start the worker
	if err := w.Run(); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
