/*
show logging request example using uber/zap

In this example, we define a middleware function called loggingMiddleware that takes an http.Handler as an argument and returns a new http.Handler. 
The middleware function logs information about each request using Zap, and then calls the next middleware or handler in the chain.

We create a new http.ServeMux instance and add our loggingMiddleware to it. The loggingMiddleware is wrapped around a simple handler function that 
writes "Hello, world!" to the response.

We then create an http.Server instance and pass in our ServeMux as the handler. We start the server by calling ListenAndServe.

In the loggingMiddleware function, we start by recording the current time using time.Now(). We then call the next middleware or handler in the chain 
using next.ServeHTTP(w, r).

After the request has been completed, we calculate the duration of the request using time.Since(startTime). We then create a new zap.Logger instance 
using zap.NewProduction() and log information about the request using logger.Info. We pass in the HTTP method, URL path, status code, and duration 
as fields using zap.String, zap.Int, and zap.Duration.

Finally, we call logger.Sync() to flush any buffered logs to the output destination.
*/

package main

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		// Call the next middleware or handler in the chain
		next.ServeHTTP(w, r)

		// Calculate the duration of the request
		duration := time.Since(startTime)

		// Log the request using Zap
		logger, _ := zap.NewProduction()
		defer logger.Sync()
		logger.Info("Request completed",
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
			zap.Int("status_code", w.(interface{ StatusCode() int }).StatusCode()),
			zap.Duration("duration", duration),
		)
	})
}

func main() {
	// Create a new router
	mux := http.NewServeMux()

	// Add the logging middleware to the router
	mux.Handle("/", loggingMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})))

	// Start the server
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}


