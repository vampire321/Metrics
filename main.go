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
)