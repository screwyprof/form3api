package form3api_test

import (
	"net/http"
	"testing"

	"github.com/screwyprof/form3api"
)

func TestAPIError_Error(t *testing.T) {
	err := &form3api.APIError{
		StatusCode: http.StatusNotFound,
		Code:       "PAGE_NOT_FOUND",
		Msg:        "Page not found",
	}
	want := "API Call Error: status code: 404, [PAGE_NOT_FOUND] - Page not found"
	form3api.Equals(t, want, err.Error())
}
