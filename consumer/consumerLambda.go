package main

import (
	"errors"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"log"
	"os"
	"github.com/aws/aws-lambda-go/lambda/messages"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

type workerReqData struct {
	FunctionName   string
	InvocationType string
	Payload        string
}

func handler() error {

	log.Print("Consuming from worm hole")
	var TASK_QUEUE_URL = os.Getenv("TASK_QUEUE_URL")
	log.Print("TASK_QUEUE_URL: ", TASK_QUEUE_URL)
	var WORKER_LAMBDA_NAME = os.Getenv("WORKER_LAMBDA_NAME")
	var AWS_REGION = os.Getenv("AWS_REGION")
	log.Print("AWS_REGION: ", AWS_REGION)
	var max_no_messages int64 = 10

	// receive messages from worm hole
	region := "us-east-1"

	awsSession := session.Must(session.NewSession(&aws.Config{Region: aws.String(region)}))

	sqsClient := sqs.New(awsSession)

	var messagesList = &sqs.ReceiveMessageOutput{}

	messagesList, err := sqsClient.ReceiveMessage(&sqs.ReceiveMessageInput{QueueUrl: &TASK_QUEUE_URL, MaxNumberOfMessages: &max_no_messages})

	if err != nil {
		log.Fatal("Error while getting messages from Worm Hole")
		return errors.New("Error while getting messages from Worm Hole!")
	}

	if len(messagesList.Messages) > 0 {
		log.Print("Received something")
		req := &messages.InvokeRequest{InvokedFunctionArn:WORKER_LAMBDA_NAME, Payload:[]byte("Hard Coded message")}
		response := &messages.InvokeResponse{}
		err := lambda.Function{}.Invoke(req,response)
		if err != nil {
			log.Print("Error occured while invoking worker Lambda ", WORKER_LAMBDA_NAME)
		}
	}
	log.Print("Processing completed")
	return nil

	// call worker lambda

}

func main() {
	log.Print("Voyager has reached in the interstellar space")
	lambda.Start(handler)
}
