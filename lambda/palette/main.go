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

func post(request events.APIGatewayProxyRequest) (*NewPaletteResponse, error) {
	var item palette.NewPalette

	err := json.Unmarshal([]byte(request.Body), &item)
	if err != nil {
		return nil, err
	}

	db := palette.CreateDB()

	res, err := db.Create(item)

	if err != nil {
		return nil, err
	}

	return &NewPaletteResponse{*res}, nil
}

func get_by_id(request events.APIGatewayProxyRequest) (*palette.Palette, error) {
	db := palette.CreateDB()

	res, err := db.GetByID(request.PathParameters["id"])

	if err != nil {
		return nil, err
	}

	return res, nil
}

func handler(request events.APIGatewayProxyRequest) (interface{}, error) {

	if request.HTTPMethod == http.MethodPost {
		return post(request)
	} else if request.HTTPMethod == http.MethodGet {
		return get_by_id(request)
	}

	log.Fatalf("Unknown route ==== %+v", request)
	return nil, nil // ???? why why why why is this needed
}

func main() {
	lambda.Start(handler)
}
