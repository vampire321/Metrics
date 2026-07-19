# Metrics

This repository contains a small Go application that demonstrates how to expose Prometheus metrics from a web service. It is useful for learning and testing instrumentation patterns without requiring a full production setup.

## What it demonstrates

- A Counter for tracking total requests
- A Gauge for monitoring active connections
- A Histogram for measuring request latency

## Features

- A simple HTTP server with example endpoints
- Randomized response behavior to simulate real traffic
- Prometheus-compatible metrics output
- Basic request and performance monitoring examples

## Run locally

### Prerequisites

- Go 1.11 or higher

### Install dependencies

```bash
go get github.com/prometheus/client_golang/prometheus
go get github.com/prometheus/client_golang/prometheus/promhttp
go get github.com/prometheus/client_golang/prometheus/promauto
```

### Start the app

```bash
go run main.go
```

The server will run on port 2112:
- Metrics endpoint: http://localhost:2112/metrics
- Test endpoint: http://localhost:2112/api/test

### Generate traffic

```bash
curl http://localhost:2112/api/test
curl http://localhost:2112/metrics
```

## Example metrics

- Total requests: myapp_total_request_total
- Active connections: myapp_active_connections
- Request duration: myapp_request_duration_seconds_bucket

## Project structure

```text
promethus/
├── main.go
├── go.mod
├── go.sum
└── README.md
```

## License

This project is open source and intended for educational use.
