package examples

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/staceybrodsky/go-httpclient/gohttp"
)

func TestCreateRepo(t *testing.T) {
	t.Run("timeout from GitHub", func(t *testing.T) {
		gohttp.FlushMocks()
		gohttp.AddMock(gohttp.Mock{
			Method:      http.MethodPost,
			Url:         "https://api.github.com/user/repos",
			RequestBody: `{"name":"test-repo","description":"","private":true}`,
			Error:       errors.New("timeout from github"),
		})

		repository := Repository{
			Name:    "test-repo",
			Private: true,
		}
		repo, err := CreateRepo(repository)

		if repo != nil {
			t.Error("no repo expected when we get a timeout from github")
		}

		if err == nil {
			t.Error("an error is expected when we get a timeout from github")
		}

		if err.Error() != "timeout from github" {
			fmt.Println(err.Error())
			t.Error("invalid error message")
		}
	})

	t.Run("noError", func(t *testing.T) {
		gohttp.FlushMocks()
		gohttp.AddMock(gohttp.Mock{
			Method:      http.MethodPost,
			Url:         "https://api.github.com/user/repos",
			RequestBody: `{"name":"test-repo","description":"","private":true}`,

			ResponseStatusCode: http.StatusCreated,
			ResponseBody:       `{"id":123,"name": "test-repo"}`,
		})

		repository := Repository{
			Name:    "test-repo",
			Private: true,
		}
		repo, err := CreateRepo(repository)

		if err != nil {
			t.Error("no error expected when we get a valid response from github")
		}

		if repo == nil {
			t.Error("a valid repo was expected at this point")
		}

		if repo.Name != repository.Name {
			t.Error("invalid repository name")
		}
	})
}