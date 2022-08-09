package gohttp

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/staceybrodsky/go-httpclient/gomime"
)

func TestGetRequestHeaders(t *testing.T) {
	// Initialization
	client := httpClient{}
	client.builder = &clientBuilder{}
	client.builder.headers = make(http.Header)

	commonHeaders := make(http.Header)
	commonHeaders.Set(gomime.HeaderContentType, gomime.ContentTypeJson)
	commonHeaders.Set(gomime.HeaderUserAgent, "cool-http-client")
	client.builder.headers = commonHeaders

	// Execution
	requestHeaders := make(http.Header)
	requestHeaders.Set("X-Request-Id", "ABC-123")

	finalHeaders := client.getRequestHeaders(requestHeaders)

	// Validation
	if len(finalHeaders) != 3 {
		t.Error("we expected 3 errors")
	}
}

func TestGetRequestBody(t *testing.T) {
	client := httpClient{}

	t.Run("with nil body", func(t *testing.T) {
		body, err := client.getRequestBody("", nil)

		if err != nil {
			t.Error("no error expected when passing a nil body")
		}
		if body != nil {
			t.Error("no body expected when passing a nil body")
		}
	})

	t.Run("with json", func(t *testing.T) {
		// json := []byte(`{
		// 	"FirstName": "John",
		// 	"LastName": "Doe"
		// }`)
		requestBody := []string{"one", "two"}
		body, err := client.getRequestBody("application/json", requestBody)

		// fmt.Println(string(json))
		fmt.Println(string(body))

		if body == nil {
			t.Error("expecting body not to be nil")
		}
		if err != nil {
			t.Error("expecting error to be nil")
		}
	})

	t.Run("with xml", func(t *testing.T) {

	})

	t.Run("with json as default", func(t *testing.T) {

	})
}
