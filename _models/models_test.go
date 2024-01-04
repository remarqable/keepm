package _models

import (
	"testing"
)

func TestUserAuthenticate(t *testing.T) {
	user := new(User)

	authenticatedUser, err := user.Authenticate("asim@remarqable.io", "Ch0ngeme!")
	if err != nil {
		t.Fatalf("Failed to authenticate user: %v", err)
	}

	if authenticatedUser == nil {
		t.Fatalf("Authentication FAILED")
	}
}

func TestUserGet(t *testing.T) {
	user := new(User)

	// Get Details for a user given an ID
	got, err := user.Get(1)
	if err != nil {
		t.Fatalf("User.Get() failed with error: %v", err)
	}

	if got.FirstName != "Asim" {
		t.Errorf("got %q, want %q", got.FirstName, "Asim")
	}
}

func TestAccountGet(t *testing.T) {
	account := &Account{}

	got, err := account.Get(1)
	if err != nil {
		t.Fatalf("Failed to get account: %v", err)
	}

	want := "RemarQable Software"
	if got.Name != want {
		t.Errorf("got %q, want %q", got.Name, want)
	}
}

func TestAccountGetUsers(t *testing.T) {
	account := Account{ID: 1}

	users, err := account.GetUsers(1)
	if err != nil {
		t.Fatalf("Account.GetUsers() failed with error: %v", err)
	}

	expectedUsersCount := 7
	if len(users) != expectedUsersCount {
		t.Fatalf("Expected %d users, but got %d", expectedUsersCount, len(users))
	}
}

func TestContactGetAll(t *testing.T) {
	contact := Contact{}

	contacts, err := contact.GetAll(1)
	if err != nil {
		t.Fatalf("Contact.GetAll() failed with error: %v", err)
	}

	expectedContactsCount := 6
	if len(contacts) != expectedContactsCount {
		t.Fatalf("Expected %d contacts, but got %d", expectedContactsCount, len(contacts))
	}
}
