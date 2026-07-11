package main

import (
	"fmt"

	"design-patterns/golang/builder"
)

func main() {
	getRequest, err := builder.NewHttpRequestBuilder("https://api.example.com/users")
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}

	getRequest.
		Method("GET").
		Header("Accept", "application/json").
		QueryParam("page", "1").
		QueryParam("limit", "10").
		Timeout(5000)

	request := getRequest.Build()
	fmt.Println("GET Request:", request)

	postRequest, err := builder.NewHttpRequestBuilder("https://api.example.com/users")
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}

	postRequest.
		Method("POST").
		Header("Content-Type", "application/json").
		Header("Authorization", "Bearer token123").
		Body(`{"name": "John Doe", "email": "john@example.com"}`).
		Timeout(10000)

	request = postRequest.Build()
	fmt.Println("\nPOST Request:", request)
}
