package gitlab

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestGetSelf(t *testing.T) {
	endpoint := "http://localhost/api/v4"
	token := os.Getenv("GITLAB_TOKEN")
	c := NewClient(endpoint, token)

	user, err := c.GetSelf()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(user)
}
