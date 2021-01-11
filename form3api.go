package form3api

import (
	"context"
	"net/http"
)

// HTTPClient an interface to abstract the http client. Used for testing purposes.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client communicates with Fake Form 3 Account API
type Client struct {
	baseURL string
	client  HTTPClient
}

// NewClient creates a new Fake Form 3 Account API client.
//
// If httpClient is nil, then built-in http client will be used.
func NewClient(httpClient HTTPClient, baseURL string) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	return &Client{
		client:  httpClient,
		baseURL: baseURL,
	}
}

// CreateAccount creates an account.
//
// Form 3 API docs: https://api-docs.form3.tech/api.html?shell#organisation-accounts-create
func (c *Client) CreateAccount(ctx context.Context, r CreateAccount) (*Account, error) {
	req, err := c.postJSONReq(ctx, c.baseURL+"/organisation/accounts", r)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var acc *Account
	if err := c.bindJSONResp(resp, &acc); err != nil {
		return nil, err
	}

	return acc, nil
}
