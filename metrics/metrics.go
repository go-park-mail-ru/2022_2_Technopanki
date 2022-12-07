package metrics

import "github.com/prometheus/client_golang/prometheus"

var SessionRequest = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "session",
}, []string{"status", "msg", "method"})

var SessionRequestDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Name: "session_request_duration",
}, []string{"method"})

var MailRequest = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "mail",
}, []string{"status", "msg", "method"})

var MailRequestDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Name: "mail_request_duration",
}, []string{"method"})
