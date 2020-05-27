package gitlab

import (
	"fmt"
	"log"
	"testing"
)

func TestGetSelf(t *testing.T) {
	endpoint := "http://localhost/api/v4"
	token := "8vJG_YxuJ5K1xTt5xeM-"
	c := NewClient(endpoint, token)

	user, err := c.GetSelf()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(user)
}
