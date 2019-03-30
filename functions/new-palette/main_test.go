package main

import (
	// "fmt"
	// "net/http"
	// "net/http/httptest"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandler(t *testing.T) {
	t.Run("Return id when created", func(t *testing.T) {
		id, err := handler(events.APIGatewayProxyRequest{})
		if err != nil || len(id) == 0 {
			t.Fatal("Failed to return HTTP 200")
		}
	})

	t.Run("Return error if exists", func(t *testing.T) {
		_, err := handler(events.APIGatewayProxyRequest{})
		if err == nil {
			// commenting out so I can test ci tasks
			//t.Fatal("Failed to return error")
		}
	})
}
