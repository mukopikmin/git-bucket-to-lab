package gitbucket

import (
	"fmt"
	"log"
	"testing"
)

func TestGetUsers(t *testing.T) {
	c := NewClient("http://localhost:8080", "855a9c623ef34a433f9118c0ddc52ec79b956d54")

	users, err := c.GetUsers()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(users)
}

// func TestCreateUser(t *testing.T) {
// 	c := NewClient("http://localhost:8080", "855a9c623ef34a433f9118c0ddc52ec79b956d54")

// 	login := faker.FirstNameMale()
// 	email := faker.Email()

// 	user, err := c.CreateUser(login, email)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	if user.Login != login {
// 		log.Fatal("assertion error")
// 	}
// }
