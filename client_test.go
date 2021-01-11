package form3api_test

import (
	"context"
	"encoding/json"
	"github.com/brianvoe/gofakeit/v6"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/screwyprof/form3api"
	"github.com/screwyprof/form3api/assert"
)

func TestNewClient(t *testing.T) {
	c := form3api.NewClient(nil, "")
	assert.NotNil(t, c)
}

func TestClientCreateAccount(t *testing.T) {
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

	h := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assertRequestMethod(t, "POST", req)

		acc := form3api.CreateAccount{}
		assert.Ok(t, json.NewDecoder(req.Body).Decode(&acc))
		assert.Equals(t, acc, r)

		_, _ = w.Write([]byte(toJSON(t, want)))
	})

	s := httptest.NewServer(h)
	defer s.Close()

	// act
	c := form3api.NewClient(s.Client(), s.URL)
	resp, err := c.CreateAccount(context.Background(), r)

	// assert
	assert.Ok(t, err)
	assert.Equals(t, want, resp)
}

func toJSON(tb testing.TB, object interface{}) string {
	tb.Helper()
	bytes, err := json.Marshal(object)
	assert.Ok(tb, err)
	return string(bytes)
}

func assertRequestMethod(tb testing.TB, want string, r *http.Request) {
	tb.Helper()
	assert.Equals(tb, want, r.Method)
}
