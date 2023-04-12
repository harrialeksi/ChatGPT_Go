/*
show an example in go how to interact with chatgpt api
https://platform.openai.com/account/api-keys

In this example, we create a ChatRequest struct to represent the request body and a ChatResponse struct to represent the response body.

We set up the API endpoint URL and request body, and use the http.NewRequest() method to create an HTTP request with the appropriate headers and 
authorization token.

We then send the API request using the http.Client.Do() method and parse the response body using the ioutil.ReadAll() and json.Unmarshal() methods.

Finally, we print the response text to the console.

Note that you will need to replace YOUR_API_KEY with your actual API key for this example to work.
*/

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ChatRequest struct {
	Text string `json:"text"`
}

type ChatResponse struct {
	Text string `json:"text"`
}

func main() {
	// Set up the API endpoint URL
	url := "https://api.openai.com/v1/engines/davinci-codex/completions"

	// Set up the API request body
	requestBody, err := json.Marshal(ChatRequest{
		Text: "Hello, ChatGPT!",
	})
	if err != nil {
		panic(err)
	}

	// Set up the API request
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		panic(err)
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer YOUR_API_KEY")

	// Send the API request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	// Parse the API response
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	var chatResponse ChatResponse
	err = json.Unmarshal(responseBody, &chatResponse)
	if err != nil {
		panic(err)
	}

	// Print the response
	fmt.Println(chatResponse.Text)
}
