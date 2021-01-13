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

	serializer *jsonSerializer
}

// NewRequest creates new instance of RequestBuilder.
func NewRequest() *RequestBuilder {
	return &RequestBuilder{client: &http.Client{}, serializer: &jsonSerializer{}}
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
	// prepare request
	req, err := rb.buildRequest(ctx, params)
	if err != nil {
		return err
	}

	// exec request
	resp, err := rb.client.Do(req)
	if err != nil {
		return err
	}

	// parse response
	return rb.bindResponse(resp, res)
}

func (rb *RequestBuilder) buildRequest(ctx context.Context, params interface{}) (*http.Request, error) {
	body, err := rb.serializer.Serialize(params)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, rb.requestMethod, rb.baseURL, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", rb.serializer.ContentType())

	return req, nil
}

func (rb *RequestBuilder) bindResponse(resp *http.Response, res interface{}) error {
	if err := rb.checkResponse(resp); err != nil {
		return err
	}

	if err := rb.serializer.Deserialize(resp.Body, res); err != nil {
		return err
	}

	return resp.Body.Close()
}

func (rb *RequestBuilder) checkResponse(resp *http.Response) error {
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest {
		var apiErr *APIError
		if err := rb.serializer.Deserialize(resp.Body, &apiErr); err != nil {
			return err
		}
		apiErr.StatusCode = resp.StatusCode
		return apiErr
	}
	return nil
}
