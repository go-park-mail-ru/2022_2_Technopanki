package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	vegeta "github.com/tsenart/vegeta/v12/lib"
)

const createURL = "http://localhost:8080/api/vacancy"

const session = "session=b94219cb-ae8f-4fae-8132-708b55e4cbba"

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
	rate := vegeta.Rate{Freq: 10000, Per: time.Second}
	duration := 100 * time.Second
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
