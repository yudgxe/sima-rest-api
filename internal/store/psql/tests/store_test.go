package psql_test

import (
	"os"
	"testing"
)

var databaseURL string

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_TEST_URL")
	if databaseURL == "" {
		databaseURL = "host=localhost port=5432 user=postgres password=00007890 dbname=db_api_test sslmode=disable"
	}
	os.Exit(m.Run())
}
