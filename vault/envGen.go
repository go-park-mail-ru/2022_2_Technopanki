package main

import (
	"flag"
	"fmt"
	"github.com/hashicorp/vault/api"
	"log"
	"os"
	"strings"
)

const secretPath = "jobflow/passwords"
const vaultAddress = "http://localhost:8200"

func main() {
	client, clientErr := api.NewClient(&api.Config{
		Address: fmt.Sprintf(vaultAddress),
	})

	if clientErr != nil {
		log.Fatalln(clientErr)
	}

	var token string
	flag.StringVar(&token, "token", "", "token for vault connecting")
	flag.Parse()
	client.SetToken(token)
	secretValues, err := client.Logical().Read(secretPath)
	if err != nil {
		log.Fatalln("get", err)
	}

	data := make([]string, 0, len(secretValues.Data))
	for name, value := range secretValues.Data {
		valueStr, ok := value.(string)
		if !ok {
			log.Fatalln("invalid data in vault")
		}

		data = append(data, fmt.Sprintf("%s=%s\n", name, valueStr))
	}

	fileEnv, OpenErr := os.Create(".env")
	if OpenErr != nil {
		log.Fatalln("Unable to create/open file:", OpenErr)
	}

	defer func(fileEnv *os.File) {
		syncErr := fileEnv.Sync()
		if syncErr != nil {
			log.Fatalln("Error with sync file", syncErr)
		}

		closeErr := fileEnv.Close()
		if closeErr != nil {
			log.Fatalln("Cannot close file:", closeErr)
		}
	}(fileEnv)

	_, writeErr := fileEnv.Write([]byte(strings.Join(data, "")))
	if writeErr != nil {
		log.Fatalln("Unable to write data:", writeErr)
	}
}
