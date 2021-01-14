package form3api

import (
	"fmt"
	"net/http"
)

// APIError represent response API error.
type APIError struct {
	Response *http.Response `json:"-"`
	Code     string         `json:"code"`
	Msg      string         `json:"error_message"`
}

func (r *APIError) Error() string {
	return fmt.Sprintf("API Call Error: %s %s: %d %s %s",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.Msg, r.Code)
}
