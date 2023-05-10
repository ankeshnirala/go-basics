package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type EnvStruct struct {
	key   string
	value string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// envVar := os.Getenv("MONGODB_URL")

	envVars := os.Environ()

	// Create a map to hold key-value pairs
	envMap := make(map[string]string)

	// Split each environment variable into key-value pairs and add them to the map
	for _, envVar := range envVars {
		parts := strings.SplitN(envVar, "=", 2)
		key := parts[0]
		value := parts[1]
		envMap[key] = value
	}

	// Convert the map to JSON
	jsonBytes, err := json.Marshal(envMap)
	if err != nil {
		panic(err)
	}

	// Print the JSON string
	jsonString := string(jsonBytes)

	// fmt.Println("Logging one env var", envVar)
	// fmt.Println("Logging all env var", envVars)
	fmt.Println("Logging all env var", jsonString)
}
