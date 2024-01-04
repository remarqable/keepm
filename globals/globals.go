package globals

import (
	"database/sql"
	"fmt"
	"os"
)

// DEBUG is a global variable that is set to true if the DEBUG environment
// variable is set to true.
var DEBUG bool = os.Getenv("DEBUG") == "true"

// DB is a global variable that holds the database connection.
var DB *sql.DB

// AppDown is a global variable that is set to true if the application is down.
var AppDown bool = false

// LOG is a global function that prints to the console if DEBUG is set to true.
func LOG(str ...interface{}) {
	if DEBUG {
		fmt.Println(str...)
	}
}
