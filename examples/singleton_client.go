package examples

import (
	"net/http"
	"time"

	"github.com/staceybrodsky/go-httpclient/gohttp"
	"github.com/staceybrodsky/go-httpclient/gomime"
)

var (
	httpClient = getHttpClient()
)

func getHttpClient() gohttp.Client {
	headers := make(http.Header)
	headers.Set(gomime.HeaderContentType, gomime.ContentTypeJson)
	client := gohttp.NewBuilder().
		SetHeaders(headers).
		SetConnectionTimeout(2 * time.Second).
		SetResponseTimeout(3 * time.Second).
		SetUserAgent("Staceys-Computer").
		Build()
	return client
}
