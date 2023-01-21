package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	vegeta "github.com/tsenart/vegeta/v12/lib"
)

const createURL = "http://localhost:8080/api/vacancy"

const session = "session=657f5d3b-c202-4cb7-8861-0ce7d1454b19"

func getBody() []byte {
	title := "aaaa"
	location := "Москва"
	salary := rand.Int()%100000 + 20000
	vacancy := fmt.Sprintf(`{    
    "title": "%s",
    "location": "%s",
    "salary": %d,
    "experience": "от 1 до 3 лет",
    "format": "online"
}`, title, location, salary)
	return []byte(vacancy)
}

func main() {
	rand.Seed(time.Now().Unix())
	rate := vegeta.Rate{Freq: 5000, Per: time.Second}
	duration := 2 * time.Second
	header := http.Header{}
	header.Add("Cookie", session)
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: http.MethodPost,
		URL:    createURL,
		Header: header,
		Body:   getBody(),
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
