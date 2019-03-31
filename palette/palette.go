package palette

import (
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/rs/xid"
)

var (
	tableName = os.Getenv("TABLE_NAME")
)

// Color represents a hex color
type Color struct {
	Name     string
	HexValue string
}

// Palette represents color palette
type Palette struct {
	ID      string
	Name    string
	Labels  []string
	Created time.Time
	Accessed  time.Time
	Colors  []Color
}

// NewPalette represents information required to create a new palette
type NewPalette struct {
	Name   string
	Labels []string
	Colors []Color
}

type PaletteDynamoDb struct {
	db dynamodbiface.DynamoDBAPI
}

func CreateDB() *PaletteDynamoDb {
	// Build the Dynamo client object
	sess := session.Must(session.NewSession())
	svc := PaletteDynamoDb{dynamodb.New(sess)}
	return &svc
}

// Create a color palette in the db
func (c *PaletteDynamoDb) Create(palette NewPalette) (*string, error) {
	// create a valid palette
	item := buildPalette(&palette)

	// Marshall the Item into a Map DynamoDB can deal with
	// why can't this be proved during compile time??
	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		fmt.Println("Got error marshalling map:")
		fmt.Println(err.Error())
		return nil, err
	}

	// Create Item in table and return
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = c.db.PutItem(input)
	if err != nil {
		return nil, err
	}

	return &item.ID, nil

}

// GetByID wraps up the DynamoDB calls to fetch a palette by Id
func (c *PaletteDynamoDb) GetByID(id string) (*Palette, error) {
	item := Palette{}

	// Perform the query
	fmt.Println("Trying to read from table: ")
	result, err := c.db.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				N: aws.String(id),
			},
		},
	})
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	// Unmarshall the result in to an Item
	err = dynamodbattribute.UnmarshalMap(result.Item, item)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return &item, nil
}

func buildPalette(palette *NewPalette) *Palette {
	timestamp := time.Now().UTC()

	// generate an id
	item := Palette{
		ID:      xid.New().String(),
		Created: timestamp,
		Accessed:  timestamp,
		Labels:  palette.Labels,
		Name:    palette.Name,
		Colors:  palette.Colors,
	}

	return &item
}
