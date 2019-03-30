package main

import (
	"errors"
	// "fmt"
	// "io/ioutil"
	// "net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	// Cannot create a resource that already exists
	ErrAlreadyExists = errors.New("Resource already exists")
)

type MyResponse struct {
	Id string `json:"id:"`
}

func handler(request events.APIGatewayProxyRequest) (MyResponse, error) {
	return MyResponse{Id: "poop"}, nil
}

func main() {
	lambda.Start(handler)
}
