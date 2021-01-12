package form3api

import (
	"context"
	"net/http"
)

// Exec builds and sends the request to url with the given method and params.
// Binds response body to res on success, returns an APIError on failure.
//
// TODO: Refactor to use a request builder to minimize the number of arguments.
// TODO: Consider moving to a separate package.
func (c *Client) Exec(ctx context.Context, method, url string, params interface{}, res interface{}) error {
	s := &jsonSerializer{}

	// prepare request
	body, err := s.Serialize(params)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", s.ContentType())

	// exec request
	resp, err := c.client.Do(req)
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
