/*
In this example, we import the github.com/sirupsen/logrus package and create a new logger using logrus.New().

We set the log level to Info using logger.SetLevel(logrus.InfoLevel).

We create an HTTP server using the http.Server struct and set its handler to a function that logs incoming requests using the 
logger.WithFields() method and sends a response.

Finally, we start the server using server.ListenAndServe() and log any errors using logger.Fatal().
*/

package main

import (
    "net/http"
    "github.com/sirupsen/logrus"
)

func main() {
    // Create a new logger
    logger := logrus.New()

    // Set the log level to Info
    logger.SetLevel(logrus.InfoLevel)

    // Create a new HTTP server
    server := http.Server{
        Addr:    ":8080",
        Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            // Log the incoming request
            logger.WithFields(logrus.Fields{
                "method": r.Method,
                "path":   r.URL.Path,
                "ip":     r.RemoteAddr,
            }).Info("Received request")

            // Send a response
            w.Write([]byte("Hello, world!"))
        }),
    }

    // Start the server
    logger.Info("Starting server on :8080")
    err := server.ListenAndServe()
    if err != nil {
        logger.Fatal(err)
    }
}
