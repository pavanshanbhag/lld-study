package main

import (
	"fmt"

	"chainofresponsibility"
)

func main() {
	auth := &chainofresponsibility.AuthenticationHandler{}
	authz := &chainofresponsibility.AuthorizationHandler{}
	validate := &chainofresponsibility.ValidationHandler{}

	auth.SetNext(authz)
	authz.SetNext(validate)

	authRequest := &chainofresponsibility.Request{
		Type: "auth",
		Data: "valid_user",
	}

	authzRequest := &chainofresponsibility.Request{
		Type: "authz",
		Data: "admin_resource",
	}

	validateRequest := &chainofresponsibility.Request{
		Type: "validate",
		Data: "test_data",
	}

	invalidRequest := &chainofresponsibility.Request{
		Type: "unknown",
		Data: "invalid_data",
	}

	fmt.Println("Processing authentication request:")
	auth.Handle(authRequest)

	fmt.Println("\nProcessing authorization request:")
	auth.Handle(authzRequest)

	fmt.Println("\nProcessing validation request:")
	auth.Handle(validateRequest)

	fmt.Println("\nProcessing invalid request:")
	auth.Handle(invalidRequest)
}
