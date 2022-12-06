package main

import (
	"fmt"
	"github.com/hashicorp/vault/api"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	client, err := api.NewClient(&api.Config{
		Address: fmt.Sprintf("http://localhost:8200"),
	})
	if err != nil {
		log.Fatalln(err)
	}

	if envErr := godotenv.Load(); envErr != nil {
		log.Fatalln("error with load .env: ", envErr)
	}

	token, ok := os.LookupEnv("TOKEN")
	if !ok {
		log.Fatalln("token not found")
	}
	client.SetToken(token)
	fmt.Println(token)
	secretValues, err := client.Logical().Read("secret/jobflow")
	if err != nil {
		log.Fatalln("get", err)
	}
	for name, value := range secretValues.Data {
		valueStr, ok := value.(string)
		if !ok {
			log.Fatalln("invalid data in vault")
		}

		setErr := os.Setenv(name, valueStr)
		if setErr != nil {
			log.Fatalln(setErr)
		}
	}
}
