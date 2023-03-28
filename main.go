package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	// Your API Key
	// Define your environment variable or directly replace with the keys
	apiKey := os.Getenv("OPENAI_API_KEY")
    organization := os.Getenv("OPENAI_ORGANIZATION")

	// Message to ChatGPT
	input := "a brief summary of The Lord of the Rings"

	// Building a HTTP POST request
	requestBody, err := json.Marshal(map[string]interface{}{
		"model":       "text-davinci-003",
		"prompt":      input,
		"max_tokens":  4000,
		"temperature": 1.0,
	})
	if err != nil {
		fmt.Println("Error building request body: ", err.Error())
		return
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/completions", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error creating request ", err.Error())
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("OpenAI-Organization", organization)

	// Execute the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error request execution: ", err.Error())
		return
	}
	defer resp.Body.Close()

	// Read the answer
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading answer: ", err.Error())
		return
	}

	// Show answer in JSON format
	fmt.Println(string(body))
}
