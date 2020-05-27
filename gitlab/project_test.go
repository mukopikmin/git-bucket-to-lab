package gitlab

import (
	"fmt"
	"log"
	"testing"

	"github.com/bxcodec/faker/v3"
)

func TestGetProjects(t *testing.T) {
	endpoint := "http://localhost/api/v4"
	token := "8vJG_YxuJ5K1xTt5xeM-"
	c := NewClient(endpoint, token)

	projects, err := c.GetProjects()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(projects)
}

func TestCreateProject(t *testing.T) {
	endpoint := "http://localhost/api/v4"
	token := "8vJG_YxuJ5K1xTt5xeM-"
	c := NewClient(endpoint, token)

	project, err := c.CreateProject(faker.Word(), faker.Sentence())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(project)
}
