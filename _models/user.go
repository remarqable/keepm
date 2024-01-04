package _models

import (
	"database/sql"
	globals "keepm/globals"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int
	AccountID int
	Email     string
	Hash      string
	LastName  string
	FirstName string
	Role      string
	Phone     string
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Get retrieves a User by their id.
// Returns a User object and an error. In case the user is not found,
// sql.ErrNoRows is returned.
func (u User) Get(ID int) (User, error) {
	var Email, LastName, FirstName string
	var CreatedAt, UpdatedAt time.Time
	var AccountID int
	var Hash, Role, Phone, Title string

	sqlStatement := `SELECT 
						id, 
						account_id, 
						email, 
						hash, 
						last_name, 
						first_name, 
						role, 
						phone, 
						title, 
						created_at, 
						updated_at 
					FROM 
						"user" 
					WHERE 
						id=$1;`

	row := globals.DB.QueryRow(sqlStatement, ID)
	switch err := row.Scan(&ID, &AccountID, &Email, &Hash, &LastName, &FirstName,
		&Role, &Phone, &Title, &CreatedAt, &UpdatedAt); err {
	case sql.ErrNoRows:
		return User{}, err // User not found
	case nil:
		user := User{ID, AccountID, Email, Hash, LastName, FirstName, Role, Phone,
			Title, CreatedAt, UpdatedAt}
		return user, nil // User found, no error
	default:
		return User{}, err // An error occurred during query
	}
}

// Add creates a new user record with the provided details.
// Returns true on successful addition.
// TODO: Implement the function to handle user addition.
func (u User) Add(Email string, password string, LastName string, FirstName string) (success bool) {
	// Implementation needed
	return true
}

// Authenticate checks if the provided email and password match a user record.
// Returns a User object and an error. In case the user is not found,
func (u *User) Authenticate(email, password string) (*User, error) {
	var id int
	var accountID int
	var hash string
	var lastName, firstName, role, phone, title string
	var createdAt, updatedAt time.Time

	sqlStatement := `SELECT 
						id, 
						account_id,
						email, 
						hash, 
						last_name, 
						first_name, 
						role, 
						phone, 
						title, 
						created_at, 
						updated_at 
					FROM 
						"user" 
					WHERE 
						email=$1;`

	row := globals.DB.QueryRow(sqlStatement, email)

	switch err := row.Scan(&id, &accountID, &email, &hash, &lastName, &firstName,
		&role, &phone, &title, &createdAt, &updatedAt); err {
	case sql.ErrNoRows:
		globals.LOG("User not found")
		return nil, nil // User not found
	case nil:
		err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
		if err == nil {
			globals.LOG("Password matches")
			// Password matches, construct user object
			u.ID = id
			u.AccountID = accountID
			u.Email = email
			u.Hash = hash
			u.LastName = lastName
			u.FirstName = firstName
			u.Role = role
			u.Phone = phone
			u.Title = title
			u.CreatedAt = createdAt
			u.UpdatedAt = updatedAt
			return u, nil
		} else {
			globals.LOG("Password does not match")
			return nil, nil // Password does not match
		}
	default:
		globals.LOG("Error")
		return nil, err // An error occurred during query
	}
}
