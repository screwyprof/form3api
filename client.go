package form3api

import (
	"net/http"

	"github.com/screwyprof/form3api/req"
	"github.com/screwyprof/form3api/resp"
)

const defaultBaseURL = "http://localhost:8080/v1"

type Client struct {
	BaseURL string
	client  *http.Client
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	return &Client{
		client:  httpClient,
		BaseURL: defaultBaseURL,
	}
}

func (c *Client) CreateAccount(r req.CreateAccount) (*resp.Account, error) {
	return nil, nil
}
