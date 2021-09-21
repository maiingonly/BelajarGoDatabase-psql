package database_golang_database

import (
	"context"
	"fmt"
	"testing"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	// update := "UPDATE users SET id= 1, first_name='otong' WHERE email='tong@otong.com';"
	delete := "DELETE FROM users WHERE id=2;"
	// users := "INSERT INTO users (id, first_name, last_name, age, email) VALUES (2,'dinda', 'sudinda', 20, 'din@dinda.com');"
	_, err := db.ExecContext(ctx, delete)
	if err != nil {
		panic(err)
	}

	fmt.Println("success insert new customer")
}
