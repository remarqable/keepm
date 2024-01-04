package _models

import (
	"database/sql"
	"keepm/globals"
	"time"
)

type Contact struct {
	ID        int
	AccountID int
	FirstName string
	LastName  string
	Company   string
	Title     string
	Phone     string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Get retrieves a Contact by their id.
// Returns a Contact object and an error. In case the contact is not found,
// sql.ErrNoRows is returned.
func (c *Contact) Get(id int) (*Contact, error) {
	var AccountID int
	var FirstName, LastName, Company, Title, Phone, Email string
	var CreatedAt, UpdatedAt time.Time

	sqlStatement := `SELECT 
                        id, 
                        account_id, 
                        first_name, 
                        last_name, 
                        company, 
                        title, 
                        phone, 
                        email, 
                        created_at, 
                        updated_at 
                    FROM 
                        "contact" 
                    WHERE 
                        id = $1;`

	row := globals.DB.QueryRow(sqlStatement, id)
	switch err := row.Scan(&id, &AccountID, &FirstName, &LastName, &Company, &Title, &Phone, &Email, &CreatedAt, &UpdatedAt); err {
	case sql.ErrNoRows:
		return nil, sql.ErrNoRows
	case nil:
		c.ID = id
		c.AccountID = AccountID
		c.FirstName = FirstName
		c.LastName = LastName
		c.Company = Company
		c.Title = Title
		c.Phone = Phone
		c.Email = Email
		c.CreatedAt = CreatedAt
		c.UpdatedAt = UpdatedAt
		return c, nil
	default:
		return nil, err
	}
}

// GetAll retrieves all contacts for an account.
// Returns a slice of Contact objects and an error.
// In case the account is not found, sql.ErrNoRows is returned.
func (c Contact) GetAll(id int) ([]Contact, error) {
	var contacts []Contact

	sqlStatement := `SELECT 
						id, 
						account_id, 
						first_name, 
						last_name, 
						company, 
						title, 
						phone, 
						email, 
						created_at, 
						updated_at 
					FROM 
						"contact" 
					WHERE 
						account_id=$1;`

	rows, err := globals.DB.Query(sqlStatement, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id, AccountID int
		var FirstName, LastName, Company, Title, Phone, Email string
		var CreatedAt, UpdatedAt time.Time

		err = rows.Scan(&id, &AccountID, &FirstName, &LastName, &Company, &Title, &Phone, &Email, &CreatedAt, &UpdatedAt)
		if err != nil {
			return nil, err
		}
		contact := Contact{id, AccountID, FirstName, LastName, Company, Title, Phone, Email, CreatedAt, UpdatedAt}
		contacts = append(contacts, contact)
	}
	return contacts, nil

}
