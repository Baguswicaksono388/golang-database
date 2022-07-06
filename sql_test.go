package golang_database

import (
	"context"
	"fmt"
	"testing"
)

func TestExectSq(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO coba_customer(id, name) VALUES(1, 'Bagus')"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new Customer")
}
