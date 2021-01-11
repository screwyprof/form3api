package form3api

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

func (c *Client) postJSONReq(ctx context.Context, endpoint string, r interface{}) (*http.Request, error) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	err := enc.Encode(r)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func (c *Client) bindJSONResp(resp *http.Response, obj interface{}) error {
	return json.NewDecoder(resp.Body).Decode(&obj)
}
