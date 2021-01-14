package form3api_test

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/screwyprof/form3api"
)

func TestAPIError_Error(t *testing.T) {
	// arrange
	want := "API Call Error: GET https://example.org/path?foo=bar: 404 Page not found PAGE_NOT_FOUND"

	baseURL, err := url.Parse("https://example.org/path?foo=bar")
	form3api.Ok(t, err)

	apiErr := &form3api.APIError{
		Response: &http.Response{
			Request: &http.Request{
				Method: http.MethodGet,
				URL:    baseURL,
			},
			StatusCode: http.StatusNotFound,
		},
		Code: "PAGE_NOT_FOUND",
		Msg:  "Page not found",
	}

	// act
	got := apiErr.Error()

	// assert
	form3api.Equals(t, want, got)
}
