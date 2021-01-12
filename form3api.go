package form3api

import (
	"context"
	"net/http"
)

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
	var acc *Account
	err := NewRequest().
		WithClient(c.client).
		WithBaseURL(c.baseURL+"/organisation/accounts").
		WithMethod(http.MethodPost).
		Exec(ctx, r, &acc)
	if err != nil {
		return nil, err
	}
	return acc, nil
}
