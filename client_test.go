package form3api_test

import (
	"bytes"
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/screwyprof/form3api"
)

func TestClientExec(t *testing.T) {
	t.Run("invalid request params given, error returned", func(t *testing.T) {
		t.Parallel()

		// arrange
		params := make(chan struct{})

		// act
		c := form3api.NewClient(nil, "")
		err := c.Exec(context.Background(), http.MethodPost, "/some_url", params, nil)

		// assert
		form3api.NotNil(t, err)
	})

	t.Run("invalid request url given, error returned", func(t *testing.T) {
		t.Parallel()

		// arrange
		c := form3api.NewClient(nil, "")

		// act
		err := c.Exec(context.Background(), "", "unknown_scheme://_invalid__url", nil, nil)

		// assert
		form3api.NotNil(t, err)
	})

	t.Run("transport error occurred, error returned", func(t *testing.T) {
		t.Parallel()

		// arrange
		client := &httpClientMock{ExpectedError: errors.New("some error")}
		c := form3api.NewClient(client, "")

		// act
		err := c.Exec(context.Background(), "", "/some_url", nil, nil)

		// assert
		form3api.NotNil(t, err)
	})

	t.Run("body deserialization error occurred, error returned", func(t *testing.T) {
		t.Parallel()

		// arrange
		client := &httpClientMock{}
		client.HandlerFunc = func(req *http.Request) (*http.Response, error) {
			resp := &http.Response{
				Body: ioutil.NopCloser(bytes.NewReader([]byte("invalid json body"))),
			}
			return resp, nil
		}

		c := form3api.NewClient(client, "")

		// act
		err := c.Exec(context.Background(), "", "/some_url", nil, nil)

		// assert
		form3api.NotNil(t, err)
	})

	t.Run("valid request, valid response", func(t *testing.T) {
		t.Parallel()

		// arrange
		req := &testRequest{SomeValue: 42}
		want := &testResponse{SomeValue: 777}

		client := &httpClientMock{
			TB:                t,
			ExpectedReqMethod: http.MethodPost,
			ExpectedReqBody:   req,
			ResponseBody:      want,
		}
		c := form3api.NewClient(client, "")

		// act
		var got *testResponse
		err := c.Exec(context.Background(), http.MethodPost, "/some_url", req, &got)

		// assert
		form3api.Ok(t, err)
		form3api.Equals(t, want, got)
	})
}

type testRequest struct {
	SomeValue int
}

type testResponse struct {
	SomeValue int
}
