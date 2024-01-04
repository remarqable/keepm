package _models

import (
	"database/sql"
	"keepm/globals"
	"time"
)

type Account struct {
	ID        int
	Name      string
	URL       string
	Address   string
	City      string
	State     string
	Zip       string
	Country   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Get retrieves an Account by their id.
// Returns an Account object and an error. In case the account is not found,
// sql.ErrNoRows is returned.
func (a *Account) Get(id int) (*Account, error) {
	var Name, URL, Address, City, State, Zip, Country string
	var CreatedAt, UpdatedAt time.Time

	sqlStatement := `SELECT 
                        id, 
                        name, 
                        url, 
                        address, 
                        city, 
                        state, 
                        zip, 
                        country, 
                        created_at, 
                        updated_at 
                    FROM 
                        "account" 
                    WHERE 
                        id = $1;`

	row := globals.DB.QueryRow(sqlStatement, id)
	switch err := row.Scan(&id, &Name, &URL, &Address, &City, &State, &Zip, &Country, &CreatedAt, &UpdatedAt); err {
	case sql.ErrNoRows:
		return nil, sql.ErrNoRows
	case nil:
		a.ID = id
		a.Name = Name
		a.URL = URL
		a.Address = Address
		a.City = City
		a.State = State
		a.Zip = Zip
		a.Country = Country
		a.CreatedAt = CreatedAt
		a.UpdatedAt = UpdatedAt
		return a, nil
	default:
		return nil, err
	}
}

func (a Account) GetUsers(id int) ([]User, error) {
	var users []User

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
						account_id=$1;`
	rows, err := globals.DB.Query(sqlStatement, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var u User
		err = rows.Scan(&u.ID, &u.AccountID, &u.Email, &u.Hash, &u.LastName, &u.FirstName, &u.Role, &u.Phone, &u.Title, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
