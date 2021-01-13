package form3api

import "fmt"

type APIError struct {
	StatusCode int
	Code       string `json:"code"`
	Msg        string `json:"message"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("API Call Error: status code: %d, [%s] - %s", e.StatusCode, e.Code, e.Msg)
}
