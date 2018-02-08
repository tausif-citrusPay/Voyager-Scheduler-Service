package main

import (
	"log"
	"github.com/aws/aws-lambda-go/lambda"
)

/*
    @author: tausif
*/

func EnquiryHandler(output map[string]interface{}) error {

	log.Print("Executing the Enquiry Worker")

	log.Print("Output: " , output)
	log.Print("Payload received: " , output["Payload"])


	log.Print("Processing completed")
	return nil


}

func main() {

	log.Print("Voyager has reached near Saturn")
	lambda.Start(EnquiryHandler)
}