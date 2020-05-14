# kontext

This library provides
 * opinionated gRPC middleware
 * function to parse user info from `context.Context`
 * unified layer to handle logging and tracing
 * function to aid in tracing instrumentation around asynchronously processed events, such as Kafka records
 * logic for errors with custom codes and descriptions, and for sending them across the wire
