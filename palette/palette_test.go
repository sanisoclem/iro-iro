package palette

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	dynamock "github.com/gusaul/go-dynamock"
)

func getDbMock() (*PaletteDynamoDb, *dynamock.DynaMock) {
	inner, mock := dynamock.New()
	svc := PaletteDynamoDb{inner}
	return &svc, mock
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}

func TestCreate(t *testing.T) {

	t.Run("Return ID when in the happy path", func(t *testing.T) {
		db, mock := getDbMock()
		mock.ExpectPutItem().WillReturns(dynamodb.PutItemOutput{})

		res, err := db.Create(NewPalette{})
		if err != nil || res == nil {
			t.Fatal("Failed to return ID")
		}
	})

	t.Run("Return nil ID and error if failed", func(t *testing.T) {
		db, _ := getDbMock()
		// -- dont put an expectation to trigger an error (weird)

		res, err := db.Create(NewPalette{})

		if res != nil || err == nil {
			t.Fatal("Returned an ID or no error even if failed")
		}
	})
}

func TestGetByID(t *testing.T) {

	t.Run("Return palette and no error when success", func(t *testing.T) {
		// I dont understand the success criteria of the marshalling enough
		// to make this work
	})

	t.Run("Return nil palette and error when fail", func(t *testing.T) {
		db, _ := getDbMock()
		// -- dont put an expectation to trigger an error (weird)

		res, err := db.GetByID("id")

		if res != nil || err == nil {
			t.Fatal("Returned an ID or no error even if failed")
		}
	})
}

func TestBuildPalette(t *testing.T) {

	t.Run("Created and accessed are same", func(t *testing.T) {
		res := buildPalette(&NewPalette{})

		if res.Accessed != res.Created {
			t.Fatal("Created and accessed are not equal")
		}
	})

	t.Run("Id should be populated and >= 16 chars", func(t *testing.T) {
		res := buildPalette(&NewPalette{})

		if len(res.ID) < 16 {
			t.Fatal("ID is too short")
		}
	})

	t.Run("Labels name and colors must be copied", func(t *testing.T) {
		payload := &NewPalette{
			Labels: []string{"test1", "test2"},
			Name:   "test",
			Colors: []Color{Color{
				HexValue: "#fff",
				Name:     "test",
			}},
		}
		res := buildPalette(payload)

		equals(t, payload.Colors, res.Colors)
		equals(t, payload.Name, res.Name)
		equals(t, payload.Labels, res.Labels)
	})
}
