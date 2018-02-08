package main

import (
	"log"
	"github.com/aws/aws-lambda-go/lambda"
	lambdaService "github.com/aws/aws-sdk-go/service/lambda"
)

/*
    @author: tausif
*/

func EnquiryHandler(output string) error {

	log.Print("Executing the Enquiry Worker")

	log.Print("Payload received: " , output)
	log.Print("Output: " , lambdaService.InvokeInput{}.Payload)


	log.Print("Processing completed")
	return nil


}

func main() {

	log.Print("Voyager has reached near Saturn")
	lambda.Start(EnquiryHandler)
}