package main

import (
	"fmt"
	"github.com/hashicorp/vault/api"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

func main() {
	client, clientErr := api.NewClient(&api.Config{
		Address: fmt.Sprintf("http://localhost:8200"),
	})

	if clientErr != nil {
		log.Fatalln(clientErr)
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
	secretValues, err := client.Logical().Read("jobflow/passwords")
	if err != nil {
		log.Fatalln("get", err)
	}

	data := fmt.Sprintf("TOKEN=%s\n", token)
	for name, value := range secretValues.Data {
		valueStr, ok := value.(string)
		if !ok {
			log.Fatalln("invalid data in vault")
		}

		data = strings.Join([]string{data, fmt.Sprintf("%s=%s", name, valueStr)}, "")
	}

	fileEnv, OpenErr := os.Create(".env")
	if OpenErr != nil {
		log.Fatalln("Unable to create/open file:", OpenErr)
	}

	defer func(fileEnv *os.File) {
		closeErr := fileEnv.Close()
		if closeErr != nil {
			log.Fatalln("Cannot close file:", closeErr)
		}
	}(fileEnv)

	_, writeErr := fileEnv.Write([]byte(data))
	if writeErr != nil {
		log.Fatalln("Unable to write data:", writeErr)
	}
}
