/*
In this example, we create a new zerolog logger with the console output and add a timestamp to each log entry. We define a handler function that logs 
incoming requests using the hlog.FromRequest function, which creates a new zerolog logger that is attached to the request context. We also define a 
middleware handler using hlog.AccessHandler that logs the request duration, status code, and response size.

We create a new http.Server instance with the middleware and handler function, and start it using the ListenAndServe method. We also use a defer 
statement to log when the server is stopped.

When a request is made to the server, the logging middleware will log the request details and the handler function will log that the request was 
handled. The logs will include the timestamp, request method and URL, and other details such as the response status and size.

Note that you can customize the log output format and level using the zerolog package's various configuration methods, such as Level, TimeFormat, 
and Caller.
*/

package main

import (
	"net/http"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
)

func main() {
	// Create a new logger with the console output
	logger := zerolog.New(zerolog.ConsoleWriter()).With().Timestamp().Logger()

	// Define the handler function
	handler := func(w http.ResponseWriter, r *http.Request) {
		// Log the incoming request
		hlog.FromRequest(r).Info().Msg("Incoming request")

		// Serve the response
		w.Write([]byte("Hello, world!"))
	}

	// Create a new middleware handler that logs requests
	logMiddleware := hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
		hlog.FromRequest(r).Info().
			Dur("duration", duration).
			Int("status", status).
			Int("size", size).
			Msg("Handled request")
	})

	// Create a new server with the logging middleware and handler function
	server := &http.Server{
		Addr:    ":8080",
		Handler: logMiddleware(http.HandlerFunc(handler)),
	}

	// Start the server and log when it is stopped
	logger.Info().Msgf("Starting server on %s", server.Addr)
	defer logger.Info().Msgf("Stopping server on %s", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		logger.Error().Err(err).Msg("Server error")
	}
}
