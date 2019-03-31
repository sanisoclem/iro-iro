package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.CloudWatchEvent) (*interface{}, error) {
	log.Printf("Trigged worker ==== %+v", request)
	return nil, nil
}

func main() {
	lambda.Start(handler)
}
