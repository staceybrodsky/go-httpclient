package examples

import (
	"errors"
	"net/http"
	"strings"
	"testing"

	"github.com/staceybrodsky/go-httpclient/gohttp"
)

func TestGet(t *testing.T) {
	gohttp.StartMockServer()

	GetEndpoints()

	t.Run("test error fetching from github", func(t *testing.T) {
		gohttp.AddMock(gohttp.Mock{
			Method: http.MethodGet,
			Url:    "https://api.github.com",
			Error:  errors.New("timeout getting github endpoints"),
		})

		endpoints, err := GetEndpoints()

		if endpoints != nil {
			t.Error("no endpoints expected")
		}

		if err == nil {
			t.Error("expected an error")
		}

		if err != nil && err.Error() != "timeout getting github endpoints" {
			t.Error("unexpected error")
		}
	})

	t.Run("test error unmarshal json response body", func(t *testing.T) {
		gohttp.AddMock(gohttp.Mock{
			Method:             http.MethodGet,
			Url:                "https://api.github.com",
			ResponseStatusCode: http.StatusOK,
			ResponseBody:       `{"current_user_url": 123}`,
		})

		endpoints, err := GetEndpoints()

		if endpoints != nil {
			t.Error("no endpoints expected")
		}

		if err == nil {
			t.Error("expected an error")
		}

		if err != nil && !strings.Contains(err.Error(), "cannot unmarshal number into Go struct field") {
			t.Error("unexpected error")
		}
	})

	t.Run("test no error", func(t *testing.T) {
		gohttp.AddMock(gohttp.Mock{
			Method:             http.MethodGet,
			Url:                "https://api.github.com",
			ResponseStatusCode: http.StatusOK,
			ResponseBody:       `{"current_user_url": "https://api.github.com/user"}`,
		})

		endpoints, err := GetEndpoints()

		if err != nil {
			t.Error("no error expected")
		}

		if endpoints == nil {
			t.Error("endpoints were expected and we got nil")
		}

		if endpoints.CurrentUserUrl != "https://api.github.com/user" {
			t.Error("unexpected error")
		}
	})

	gohttp.StopMockServer()
}
