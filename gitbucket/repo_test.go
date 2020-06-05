package gitbucket

import (
	"os"
	"testing"

	"github.com/bxcodec/faker/v3"
)

func TestCreateRepo(t *testing.T) {
	c := NewClient("http://localhost:8080", os.Getenv("GITBUCKE_TOKEN"))

	name := faker.Word()
	description := "descriptipn"
	private := false

	repo, err := c.CreateRepo(name, description, private)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	if repo.Name != name || repo.Description != description || repo.Private != private {
		t.Log(repo)
		t.Log(name)
		t.Log(description)
		t.Log(private)
		t.Fatal("assertion error")
	}
}
