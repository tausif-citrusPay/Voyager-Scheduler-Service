package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"net/http"
)

/*
   @author: tausif
*/

func EnquiryHandler(output string) error {

	log.Print("Executing the Enquiry Worker")

	log.Print("Payload received: ", output)
	url := "https://pay.blazepay.in/info"

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return err
	}

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return err
	}

	log.Print("BLAZEPAY VERSION :", resp)

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()


	log.Print("Processing completed")
	return nil

}

func main() {

	log.Print("Voyager has reached near Saturn")
	lambda.Start(EnquiryHandler)
}
