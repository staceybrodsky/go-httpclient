package gomime

import "testing"

func TestHeaders(t *testing.T) {
	if HeaderContentType != "Content-Type" {
		t.Error("invalid content type header")
	}

	if HeaderUserAgent != "User-Agent" {
		t.Error("invalid content type user agent")
	}

	if ContentTypeJson != "application/json" {
		t.Error("invalid content type json")
	}

	if ContentTypeXml != "application/xml" {
		t.Error("invalid content type xml")
	}

	if ContentTypeOctetStream != "application/octet-stream" {
		t.Error("invalid content type octet-stream")
	}
}
