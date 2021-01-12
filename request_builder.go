package form3api

import (
	"context"
	"net/http"
)

// HTTPClient an interface to abstract the http client. Used for testing purposes.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// RequestBuilder builds and sends the API requests.
type RequestBuilder struct {
	client        HTTPClient
	baseURL       string
	requestMethod string
}

// NewRequest creates new instance of RequestBuilder.
func NewRequest() *RequestBuilder {
	return &RequestBuilder{client: &http.Client{}}
}

// WithClient sets an http client
func (rb *RequestBuilder) WithClient(client HTTPClient) *RequestBuilder {
	rb.client = client
	return rb
}

// WithBaseURL sets endpoint url.
func (rb *RequestBuilder) WithBaseURL(baseURL string) *RequestBuilder {
	rb.baseURL = baseURL
	return rb
}

// WithMethod sets request method.
func (rb *RequestBuilder) WithMethod(method string) *RequestBuilder {
	rb.requestMethod = method
	return rb
}

// Exec builds and sends the request to endpoint url with the given method and params.
// Binds response body to res on success, returns an error on failure.
func (rb *RequestBuilder) Exec(ctx context.Context, params interface{}, res interface{}) error {
	s := &jsonSerializer{}

	// prepare request
	body, err := s.Serialize(params)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, rb.requestMethod, rb.baseURL, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", s.ContentType())

	// exec request
	resp, err := rb.client.Do(req)
	if err != nil {
		return err
	}

	// parse response
	defer resp.Body.Close()

	// TODO: Check resp.StatusCode
	// TODO: Check resp.ContentType
	// TODO: Handle errors, use resp.StatusCode on error
	if err = s.Deserialize(resp.Body, res); err != nil {
		return err
	}

	return nil
}
