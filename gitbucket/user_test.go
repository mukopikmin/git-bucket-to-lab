package gitbucket

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestGetUsers(t *testing.T) {
	c := NewClient("http://localhost:8080", os.Getenv("GITBUCKE_TOKEN"))

	users, err := c.GetUsers()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(users)
}

// func TestCreateUser(t *testing.T) {
// 	c := NewClient("http://localhost:8080", os.Getenv("GITBUCKE_TOKEN"))

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
