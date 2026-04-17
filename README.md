# Prometheus Metrics Application

A Go application demonstrating Prometheus metrics collection with three types of metrics: Counter, Gauge, and Histogram.

## Overview

This project showcases how to implement and expose metrics using the Prometheus Go client library. It includes a simple HTTP server that collects metrics about request patterns and performance.

## Metrics Types

### 1. Counter (Type 1)
- **Metric**: `myapp_total_request`
- **Description**: Cumulative count of total requests received
- **Labels**: method, path, status
- **Use Case**: Track the total number of requests, errors, and successful responses over time

### 2. Gauge (Type 2)
- **Metric**: `myapp_active_connections`
- **Description**: Current number of active connections
- **Use Case**: Monitor real-time connection activity on the server

### 3. Histogram (Type 3)
- **Metric**: `myapp_request_duration_seconds`
- **Description**: Distribution of request duration in seconds
- **Labels**: path
- **Buckets**: 0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5 seconds
- **Use Case**: Analyze request latency distribution and identify performance bottlenecks

## Features

- Default status code: 200 (success)
- 10% random error rate (status 500)
- WithLabelValues: Maps actual request data to metrics
- Simulated request processing delay (0-200ms)

## Installation

### Prerequisites
- Go 1.11 or higher

### Dependencies
```bash
go get github.com/prometheus/client_golang/prometheus
go get github.com/prometheus/client_golang/prometheus/promhttp
go get github.com/prometheus/client_golang/prometheus/promauto
```

## Usage

1. Run the application:
```bash
go run main.go
```

2. The server will start on port 2112:
   - Metrics endpoint: `http://localhost:2112/metrics`
   - Test endpoint: `http://localhost:2112/api/test`

3. Generate some requests:
```bash
curl http://localhost:2112/api/test
```

4. View metrics in Prometheus format:
```bash
curl http://localhost:2112/metrics
```

## Example Prometheus Queries

Once metrics are exposed:

- Total requests: `myapp_total_request_total`
- Active connections: `myapp_active_connections`
- Request duration: `myapp_request_duration_seconds_bucket`

## Project Structure

```
promethus/
├── main.go        # Main application with metrics
├── go.mod         # Module definition
├── go.sum         # Dependency checksums
└── README.md      # This file
```

## License

This project is open source and available for educational purposes.
