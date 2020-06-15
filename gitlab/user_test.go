package gitlab

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestGetAuthorizedUser(t *testing.T) {
	endpoint := "http://localhost/api/v4"
	token := os.Getenv("GITLAB_TOKEN")
	c := NewClient(endpoint, token)

	user, err := c.GetAuthorizedUser()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(user)
}
