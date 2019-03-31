package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sanisoclem/iro-iro/palette"
)

// NewPaletteResponse describes a succesful palette creation
type NewPaletteResponse struct {
	ID string `json:"id:"`
}

func post(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	item := palette.NewPalette{}

	// create a transient error return because error handling in go is hard
	retval := events.APIGatewayProxyResponse{
		StatusCode:      500,
		IsBase64Encoded: false,
	}

	// deserialize json
	err := json.Unmarshal([]byte(request.Body), &item)
	if err != nil {
		return retval, err
	}

	// save thing to db
	db := palette.CreateDB()
	res, err := db.Create(item)
	if err != nil {
		return retval, err
	}

	// serialize json
	ser, err := json.Marshal(*res)
	if err != nil {
		return retval, err
	}

	// update the return value
	retval.Body = string(ser)
	retval.StatusCode = 200

	return retval, nil
}

func getByID(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	db := palette.CreateDB()

	retval := events.APIGatewayProxyResponse{
		StatusCode:      500,
		IsBase64Encoded: false,
	}

	res, err := db.GetByID(request.PathParameters["id"])
	if err != nil {
		return retval, err
	}

	// serialize to json
	ser, err := json.Marshal(*res)
	if err != nil {
		return retval, err
	}

	retval.StatusCode = 200
	retval.Body = string(ser)

	return retval, nil
}

func handler(request events.APIGatewayProxyRequest) (interface{}, error) {

	if request.HTTPMethod == http.MethodPost {
		return post(request)
	} else if request.HTTPMethod == http.MethodGet {
		return getByID(request)
	}

	log.Fatalf("Unknown route ==== %+v", request)
	return nil, nil // ???? why why why why is this needed
}

func main() {
	lambda.Start(handler)
}
