package psql

import (
	"database/sql"
	"fmt"
	"strings"
	"testing"

	_ "github.com/lib/pq"
)

func GetTestStore(t *testing.T, databaseURL string) (*Store, func(...string)) {
	t.Helper()
	
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}

	s := New(db)

	return s, func(tables ...string) {
		defer db.Close()

		if len(tables) > 0 {
			db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", ")))
		}
	}
}

