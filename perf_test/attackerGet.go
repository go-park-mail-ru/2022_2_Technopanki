package main

import (
	"fmt"
	vegeta "github.com/tsenart/vegeta/v12/lib"
	"math/rand"
	"net/http"
	"time"
)

const getURL = "http://localhost:8080/api/vacancy/%d"

func main() {
	rand.Seed(time.Now().Unix())
	rate := vegeta.Rate{Freq: 20000, Per: time.Second}
	duration := 50 * time.Second
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: http.MethodGet,
		URL:    fmt.Sprintf(getURL, rand.Int()%1000),
	})
	attacker := vegeta.NewAttacker()

	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
		metrics.Add(res)
	}
	metrics.Close()

	fmt.Printf("99th percentile: %s\n", metrics.Latencies.P99)
	fmt.Printf("RPS: %f\n", metrics.Rate)
}
