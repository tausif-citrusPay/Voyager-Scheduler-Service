package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"log"
	"os"
)

func handler(ctx context.Context) {

	log.Print("Consuming from worm hole")
	if ctx != nil {
		log.Print("Context: ", ctx)
	}
	var TASK_QUEUE_URL = os.Getenv("TASK_QUEUE_URL")
	log.Print("TASK_QUEUE_URL: ", TASK_QUEUE_URL)
	//var WORKER_LAMBDA_NAME = os.Getenv("WORKER_LAMBDA_NAME")
	var AWS_REGION = os.Getenv("AWS_REGION")
	log.Print("AWS_REGION: ", AWS_REGION)
	var max_no_messages int64 = 10

	// receive messages from worm hole
	region := "us-east-1"

	awsSession := session.Must(session.NewSession(&aws.Config{Region: aws.String(region)}))

	sqsClient := sqs.New(awsSession)

	var messages = &sqs.ReceiveMessageOutput{}

	messages, err := sqsClient.ReceiveMessage(&sqs.ReceiveMessageInput{QueueUrl: &TASK_QUEUE_URL, MaxNumberOfMessages: &max_no_messages})

	if err != nil {
		log.Fatal("Error while getting messages from Worm Hole")
	}

	if len(messages.Messages) > 0 {
		log.Print("Received something")
	}
	// call worker lambda

}

func main() {

	log.Print("Voyager has reached in the interstellar space")
	lambda.Start(handler)
}
