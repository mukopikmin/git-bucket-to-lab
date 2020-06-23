package gitlab

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestGetProjects(t *testing.T) {
	endpoint := "http://localhost/api/v4"
	token := os.Getenv("GITLAB_TOKEN")
	c := NewClient(endpoint, token)

	projects, err := c.GetProjects()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(projects)
}

// func TestCreateProject(t *testing.T) {
// 	endpoint := "http://localhost/api/v4"
// 	token := os.Getenv("GITLAB_TOKEN")
// 	c := NewClient(endpoint, token)

// 	project, err := c.CreateProject(faker.Word(), faker.Sentence())
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println(project)
// }
