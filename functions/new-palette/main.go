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
	// ErrAlreadyExists Cannot create a resource that already exists
	ErrAlreadyExists = errors.New("Resource already exists")
)

// NewPaletteResponse describes a succesful palette creation
type NewPaletteResponse struct {
	ID string `json:"id:"`
}

func handler(request events.APIGatewayProxyRequest) (NewPaletteResponse, error) {
	return NewPaletteResponse{ID: "poop"}, nil
}

func main() {
	lambda.Start(handler)
}
