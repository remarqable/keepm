package util

import (
	"strings"
)

// EmptyUserPass checks if the username or password is empty after trimming spaces.
// Returns true if either username or password is effectively empty.
func EmptyUserPass(username, password string) bool {
	return strings.Trim(username, " ") == "" || strings.Trim(password, " ") == ""
}

func Render() {
	// Function implementation needed
}
