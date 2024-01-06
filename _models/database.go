package _models

import (
	"context"
	"database/sql"
	"fmt"
	"keepm/globals"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// DB is a global variable that holds the database connection.
func init() {
	var err error
	globals.DB, err = DB()
	if err != nil {
		log.Fatalf("%v", err)
	}
}

// DB returns a database connection.
func DB() (*sql.DB, error) {
	// TODO: Load configuration from environment variables or .env file
	connectionString := `
		host=localhost 
		port=5432 
		user=postgres
		dbname=keepmdb 
		sslmode=disable`

	var err error
	globals.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	return globals.DB, nil
}

// transaction executes an array of SQL statements in a single transaction.
// This ensures that either all statements are successfully executed, or none
// are in case of an error.
func transaction(db *sql.DB, statements []string) {
	fmt.Println("Setting up database... ")
	var ctx = context.Background()
	var tx, err = db.BeginTx(ctx, nil)
	if err != nil {
		fmt.Println("Couldn't begin transaction")
		panic(err)
	}
	for _, statement := range statements {
		fmt.Println("Running statement:")
		fmt.Println(statement)

		_, err = tx.ExecContext(ctx, statement)
		if err != nil {
			tx.Rollback()
			if err != nil {

				panic(err)
			}
		}
	}
	err = tx.Commit()
	if err != nil {
		fmt.Println("Failed to commit transaction")
		panic(err)
	}
}
