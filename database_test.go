package database_golang_database

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
)

func TestEmpty(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("postgres", "root:@tcp(localhost:5432)/youdb")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
