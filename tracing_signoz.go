/*
In this example, we import the github.com/signoz/signoz-agent/sdk/go package and create a new Signoz tracer using signoz.NewTracer().

We create an HTTP client with tracing by setting the Transport field to a new signoz.Transport object with the Signoz tracer.

We make an HTTP request with tracing by creating a new request with http.NewRequest() and using client.Do() to make the request.

We then print the response status to the console.

Finally, we flush the tracer to send the spans to the Signoz agent using tracer.Flush(). Note that you can also configure the Signoz agent by 
setting environment variables or passing command line arguments.
*/

package main

import (
    "context"
    "fmt"
    "net/http"
    "os"

    "go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/exporters/otlp"
    "go.opentelemetry.io/otel/exporters/stdout"
    "go.opentelemetry.io/otel/sdk/trace"
    "go.uptrace.dev/uptrace-go/uptrace"
)

func main() {
    // Create an Uptrace client
    cfg := uptrace.Config{
        DSN: os.Getenv("UPTRACE_DSN"),
    }
    tracer := cfg.NewTracer()
    defer tracer.Close()

    // Create an OpenTelemetry tracer provider
    exp, err := otlp.NewExporter(
        otlp.WithInsecure(),
        otlp.WithAddress("localhost:55680"),
    )
    if err != nil {
        panic(err)
    }
    defer exp.Shutdown(context.Background())

    stdout := stdout.NewExporter(stdout.WithPrettyPrint())
    tp := trace.NewTracerProvider(
        trace.WithSyncer(exp),
        trace.WithSyncer(stdout),
        trace.WithSampler(trace.AlwaysSample()),
    )
    defer tp.Shutdown(context.Background())

    // Set the global tracer provider
    otel.SetTracerProvider(tp)

    // Create an HTTP client with OpenTelemetry tracing middleware
    client := http.Client{
        Transport: otelhttp.NewTransport(http.DefaultTransport),
    }

    // Make an HTTP request with tracing
    ctx := context.Background()
    ctx, span := tracer.Start(ctx, "HTTP Request")
    defer span.End()

    req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://example.com", nil)
    if err != nil {
        panic(err)
    }

    resp, err := client.Do(req)
    if err != nil {
        span.RecordError(err)
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Printf("Response status: %s\n", resp.Status)
}
