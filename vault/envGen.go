package main

import (
	"fmt"
	"github.com/hashicorp/vault/api"
	"log"
	"os"
)

func main() {
	client, err := api.NewClient(&api.Config{
		Address: fmt.Sprintf("http://localhost:8200"),
	})

	if err != nil {
		log.Fatal(err)
	}

	token := os.Getenv("TOKEN")
	client.SetToken(token)
	secretValues, err := client.Logical().Read("secret/data/jf")
	if err != nil {
		log.Println(err)
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
