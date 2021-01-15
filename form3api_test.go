package form3api_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"github.com/brianvoe/gofakeit/v6"

	"github.com/screwyprof/form3api"
)

func TestNewClient(t *testing.T) {
	c := form3api.NewClient(nil, "")
	form3api.NotNil(t, c)
}

func TestClientCreateAccount(t *testing.T) {
	t.Run("valid request given, account created", func(t *testing.T) {
		// arrange
		ID := gofakeit.UUID()
		organisationID := gofakeit.UUID()
		accountNumber := gofakeit.Numerify("#########")
		bankID := gofakeit.Numerify("######")

		want := &form3api.Account{
			AccountData: form3api.AccountData{
				ID:             ID,
				OrganisationID: organisationID,
				Type:           "accounts",
				Attributes: &form3api.AccountAttributes{
					AccountNumber: accountNumber,
					BankID:        bankID,
					BankIDCode:    "GBDSC",
					BIC:           "NWBKGB42",
					Country:       "GB",
					Currency:      "GBP",
					ConfirmationOfPayee: &form3api.ConfirmationOfPayee{
						AccountClassification: "Personal",
					},
				},
			},
			Links: form3api.Links{
				Self: "/v1/organisation/accounts/" + ID,
			},
		}

		r := form3api.CreateAccount{
			AccountData: form3api.AccountData{
				ID:             ID,
				OrganisationID: organisationID,
				Type:           "accounts",
				Attributes: &form3api.AccountAttributes{
					AccountNumber: accountNumber,
					BankID:        bankID,
					BankIDCode:    "GBDSC",
					Country:       "GB",
					Currency:      "GBP",
					CustomerID:    "234",
					IBAN:          "GB28NWBK40030212764204",
					BIC:           "NWBKGB42",
					ConfirmationOfPayee: &form3api.ConfirmationOfPayee{
						AccountClassification: "Personal",
					},
				},
			},
		}

		client := &httpClientMock{
			TB:                t,
			ExpectedReqMethod: http.MethodPost,
			ExpectedReqBody:   r,
			StatusCode:        http.StatusCreated,
			ResponseBody:      want,
		}

		// act
		c := form3api.NewClient(client, "")
		got, err := c.CreateAccount(context.Background(), r)

		// assert
		form3api.Ok(t, err)
		form3api.Equals(t, want, got)
	})

	t.Run("an error occurred, error returned", func(t *testing.T) {
		// arrange
		client := &httpClientMock{ExpectedError: errors.New("some error")}
		c := form3api.NewClient(client, "")

		// act
		_, err := c.CreateAccount(context.Background(), form3api.CreateAccount{})

		// assert
		form3api.NotNil(t, err)
	})
}

func TestClientFetchAccount(t *testing.T) {
	t.Run("valid request given, account created", func(t *testing.T) {
		// arrange
		ID := gofakeit.UUID()
		want := generateTestAccount(ID)

		r := form3api.FetchAccount{
			AccountID: ID,
		}

		client := &httpClientMock{
			TB:                t,
			ExpectedReqMethod: http.MethodGet,
			StatusCode:        http.StatusOK,
			ResponseBody:      want,
		}

		// act
		c := form3api.NewClient(client, "")
		got, err := c.FetchAccount(context.Background(), r)

		// assert
		form3api.Ok(t, err)
		form3api.Equals(t, want, got)
	})

	t.Run("an error occurred, error returned", func(t *testing.T) {
		// arrange
		client := &httpClientMock{ExpectedError: errors.New("some error")}
		c := form3api.NewClient(client, "")

		// act
		_, err := c.FetchAccount(context.Background(), form3api.FetchAccount{})

		// assert
		form3api.NotNil(t, err)
	})
}

func TestClientDeleteAccount(t *testing.T) {
	t.Run("valid request given, account created", func(t *testing.T) {
		// arrange
		r := form3api.DeleteAccount{
			AccountID: gofakeit.UUID(),
		}

		client := &httpClientMock{
			TB:                t,
			ExpectedReqMethod: http.MethodDelete,
			StatusCode:        http.StatusNoContent,
		}

		// act
		c := form3api.NewClient(client, "")
		err := c.DeleteAccount(context.Background(), r)

		// assert
		form3api.Ok(t, err)
	})

	t.Run("an error occurred, error returned", func(t *testing.T) {
		// arrange
		client := &httpClientMock{ExpectedError: errors.New("some error")}
		c := form3api.NewClient(client, "")

		// act
		err := c.DeleteAccount(context.Background(), form3api.DeleteAccount{})

		// assert
		form3api.NotNil(t, err)
	})
}

func TestClientListAccounts(t *testing.T) {
	t.Run("valid request given, account created", func(t *testing.T) {
		// arrange
		want := &form3api.Accounts{
			AccountData: []form3api.AccountData{
				generateTestAccount(gofakeit.UUID()).AccountData,
				generateTestAccount(gofakeit.UUID()).AccountData,
			},
			Links: form3api.Links{
				Self:  "/v1/organisation/accounts?page%5Bnumber%5D=1&page%5Bsize%5D=2",
				First: "/v1/organisation/accounts?page%5Bnumber%5D=first&page%5Bsize%5D=2",
				Last:  "/v1/organisation/accounts?page%5Bnumber%5D=last&page%5Bsize%5D=2",
				Prev:  "/v1/organisation/accounts?page%5Bnumber%5D=0&page%5Bsize%5D=2",
				Next:  "/v1/organisation/accounts?page%5Bnumber%5D=2&page%5Bsize%5D=2",
			},
		}

		r := form3api.ListAccounts{
			Page: form3api.Page{Number: 1, Size: 2},
		}

		client := &httpClientMock{
			TB:                t,
			ExpectedReqMethod: http.MethodGet,
			StatusCode:        http.StatusOK,
			ResponseBody:      want,
		}

		// act
		c := form3api.NewClient(client, "")
		got, err := c.ListAccounts(context.Background(), r)

		// assert
		form3api.Ok(t, err)
		form3api.Equals(t, want, got)
	})

	t.Run("an error occurred, error returned", func(t *testing.T) {
		// arrange
		client := &httpClientMock{ExpectedError: errors.New("some error")}
		c := form3api.NewClient(client, "")

		// act
		_, err := c.ListAccounts(context.Background(), form3api.ListAccounts{})

		// assert
		form3api.NotNil(t, err)
	})
}

func generateTestAccount(ID string) *form3api.Account {
	return &form3api.Account{
		AccountData: form3api.AccountData{
			ID:             ID,
			OrganisationID: gofakeit.UUID(),
			Type:           "accounts",
			Attributes: &form3api.AccountAttributes{
				AccountNumber: gofakeit.Numerify("#########"),
				BankID:        gofakeit.Numerify("######"),
				BankIDCode:    "GBDSC",
				BIC:           "NWBKGB42",
				Country:       "GB",
				Currency:      "GBP",
				ConfirmationOfPayee: &form3api.ConfirmationOfPayee{
					AccountClassification: "Personal",
				},
			},
		},
		Links: form3api.Links{
			Self: "/v1/organisation/accounts/" + ID,
		},
	}
}

type httpClientMock struct {
	TB testing.TB

	ExpectedError error

	ExpectedReqMethod string
	ExpectedReqBody   interface{}

	StatusCode   int
	ResponseBody interface{}
	HandlerFunc  func(req *http.Request) (*http.Response, error)
}

func (c *httpClientMock) Do(req *http.Request) (*http.Response, error) {
	if c.HandlerFunc == nil {
		c.HandlerFunc = c.defaultHandler
	}
	return c.HandlerFunc(req)
}

func (c *httpClientMock) defaultHandler(req *http.Request) (*http.Response, error) {
	if c.ExpectedError != nil {
		return nil, c.ExpectedError
	}

	assertRequestMethod(c.TB, c.ExpectedReqMethod, req)
	assertRequestBody(c.TB, c.ExpectedReqBody, req)

	return &http.Response{
		StatusCode: c.StatusCode,
		Body:       ioutil.NopCloser(bytes.NewReader(toJSONBytes(c.TB, c.ResponseBody))),
	}, nil
}

func assertRequestMethod(tb testing.TB, want string, r *http.Request) {
	tb.Helper()
	form3api.Equals(tb, want, r.Method)
}

func assertRequestBody(tb testing.TB, want interface{}, r *http.Request) {
	tb.Helper()
	if want == nil {
		return
	}
	wantType := reflect.TypeOf(want)

	got := reflect.New(wantType).Interface()
	form3api.Ok(tb, json.NewDecoder(r.Body).Decode(&got))

	wantPtr := reflect.New(wantType)
	wantPtr.Elem().Set(reflect.ValueOf(want))

	form3api.Equals(tb, wantPtr.Interface(), got)
}

func toJSONBytes(tb testing.TB, object interface{}) []byte {
	tb.Helper()
	jsonBytes, err := json.Marshal(object)
	form3api.Ok(tb, err)
	return jsonBytes
}
