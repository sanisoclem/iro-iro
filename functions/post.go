package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	palette "github.com/sanisoclem/iro-iro/api"
)

// NewPaletteResponse describes a succesful palette creation
type NewPaletteResponse struct {
	ID string `json:"id:"`
}

func handler(request events.APIGatewayProxyRequest) (*NewPaletteResponse, error) {
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

func main() {
	lambda.Start(handler)
}
