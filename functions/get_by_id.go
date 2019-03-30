package main

import (
	"errors"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	palette "github.com/sanisoclem/iro-iro/api"
)

var (
	// ErrNotExists means the id was not found
	ErrNotExists = errors.New("Resource does not exists")
)

func handler(request events.APIGatewayProxyRequest) (*palette.Palette, error) {
	db := palette.CreateDB()

	res, err := db.GetByID(request.Body)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func main() {
	lambda.Start(handler)
}
