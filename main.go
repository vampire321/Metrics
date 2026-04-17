package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

//COUNTER(odometer) - A counter is a cumulative metric that represents a single numerical value that only ever goes up. It is typically used to count the number of events or occurrences of something, such as the number of requests received by a server or the number of errors that have occurred. Counters can be reset to zero, but they cannot decrease in value.
var (
	requestTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name:"myapp_total_request",
		Help:"The total number of requests received",
	},[]string{"method","path","status"})

	 // GAUGE — current active connections
    activeConns = promauto.NewGauge(prometheus.GaugeOpts{
        Name: "myapp_active_connections",
        Help: "Current active connections",
    })

	// HISTOGRAM — request duration distribution
    requestDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
        Name:    "myapp_request_duration_seconds",
        Help:    "Request duration in seconds",
        Buckets: []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5},
    }, []string{"path"})
)

func handleRequest(w http.ResponseWriter, r *http.Request){
	start := time.Now()
	activeConns.Inc()
	defer activeConns.Dec()

	time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)

    status := 200
    if rand.Intn(10) == 0 { // 10% error rate
        status = 500
    }

	requestTotal.WithLabelValues(r.Method, r.URL.Path,
        fmt.Sprintf("%d", status)).Inc()

	//record histogram
	requestDuration.WithLabelValues(r.URL.Path).
	Observe(time.Since(start).Seconds())

	w.WriteHeader(status)
}

func main(){
	http.HandleFunc("/api/test", handleRequest)
    http.Handle("/metrics", promhttp.Handler())

    fmt.Println("Visit http://localhost:2112/metrics")
    fmt.Println("Call http://localhost:2112/api/test a few times first")
    http.ListenAndServe(":2112", nil)
}

