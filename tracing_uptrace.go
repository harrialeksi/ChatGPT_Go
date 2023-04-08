/*
show distributed tracing example in go using uptrace
compare example in https://github.com/uptrace/uptrace/tree/master/example/gin-gorm

In this example, we import the go.uptrace.dev/uptrace-go package and create an uptrace.Config object with the DSN from an environment variable. 
We create an Uptrace tracer using cfg.NewTracer() and defer calling tracer.Close() to ensure that it is properly closed.

We then create an OpenTelemetry tracer provider with an OTLP exporter and a stdout exporter. We set the global tracer provider to this provider using 
otel.SetTracerProvider().

We create an HTTP client with OpenTelemetry tracing middleware using otelhttp.NewTransport(). We then make an HTTP request with tracing using 
tracer.Start() to create a new span, http.NewRequestWithContext() to create the request, and client.Do() to make the request.

We then print the response status to the console.

Note that this is a simple example and you can customize the tracer provider and exporters to fit your own use case.

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
