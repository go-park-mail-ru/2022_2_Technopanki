package main

import (
	"HeadHunter/entity"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	objExample := entity.Employer{
		Name:     "Tinkoff",
		Email:    "tinkoff@gmail.com",
		Password: "345",
	}
	jsonExample, marshErr := json.Marshal(objExample)
	if marshErr != nil {
		log.Fatal(marshErr)
	}
	resp, err := http.Post("http://localhost:8080/employers", "application/json",
		bytes.NewBuffer(jsonExample))

	if err != nil {
		log.Fatal(err)
	}

	var res map[string]interface{}
	if resp.Status != "201 Created" {
		log.Fatal("Error")
	}
	json.NewDecoder(resp.Body).Decode(&res)

	fmt.Println(res["json"])
}
