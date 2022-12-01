package metrics

import "github.com/prometheus/client_golang/prometheus"

var SessionRequest = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "session",
}, []string{"status", "msg"})

var SessionRequestDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Name: "session_request_duration",
}, []string{"method"})
