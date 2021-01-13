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

func TestRequestBuilder(t *testing.T) {
	t.Run("invalid request params given, error returned", func(t *testing.T) {
		t.Parallel()

		// arrange
		params := make(chan struct{})
		rb := form3api.NewRequest()

		// act
		err := rb.Exec(context.Background(), params, nil)

		// assert
		form3api.NotNil(t, err)
	})

	t.Run("invalid request url given, error returned", func(t *testing.T) {
		t.Parallel()

		// arrange
		rb := form3api.NewRequest().
			WithBaseURL("unknown_scheme://_invalid__url")

		// act
		err := rb.Exec(context.Background(), nil, nil)

		// assert
		form3api.NotNil(t, err)
	})

	t.Run("transport error occurred, error returned", func(t *testing.T) {
		t.Parallel()

		// arrange
		client := &httpClientMock{ExpectedError: errors.New("some error")}
		rb := form3api.NewRequest().
			WithClient(client).
			WithBaseURL("/some_url")

		// act
		err := rb.Exec(context.Background(), nil, nil)

		// assert
		form3api.NotNil(t, err)
	})

	t.Run("known API error occurred, error returned", func(t *testing.T) {
		t.Parallel()

		// arrange
		want := &form3api.APIError{
			StatusCode: http.StatusInternalServerError,
			Code:       "SOME_ERROR",
			Msg:        "some error",
		}

		client := &httpClientMock{
			TB:                t,
			ExpectedReqMethod: http.MethodGet,
			StatusCode:        want.StatusCode,
			ResponseBody:      want,
		}
		rb := form3api.NewRequest().
			WithClient(client).
			WithBaseURL("/some_url")

		// act
		err := rb.Exec(context.Background(), nil, nil)

		// assert
		form3api.NotNil(t, err)
	})

	t.Run("unknown API error occurred, error returned", func(t *testing.T) {
		t.Parallel()

		// arrange
		client := &httpClientMock{}
		client.HandlerFunc = func(req *http.Request) (*http.Response, error) {
			resp := &http.Response{
				StatusCode: http.StatusInternalServerError,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("invalid json body"))),
			}
			return resp, nil
		}
		rb := form3api.NewRequest().
			WithClient(client).
			WithBaseURL("/some_url")

		// act
		err := rb.Exec(context.Background(), nil, nil)

		// assert
		form3api.NotNil(t, err)
	})

	t.Run("body deserialization error occurred, error returned", func(t *testing.T) {
		t.Parallel()

		// arrange
		client := &httpClientMock{}
		client.HandlerFunc = func(req *http.Request) (*http.Response, error) {
			resp := &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("invalid json body"))),
			}
			return resp, nil
		}

		rb := form3api.NewRequest().
			WithClient(client).
			WithBaseURL("/some_url")

		// act
		err := rb.Exec(context.Background(), nil, nil)

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
			StatusCode:        http.StatusCreated,
		}
		rb := form3api.NewRequest().
			WithClient(client).
			WithBaseURL("/some_url").
			WithMethod(http.MethodPost)

		// act
		var got *testResponse
		err := rb.Exec(context.Background(), req, &got)

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
