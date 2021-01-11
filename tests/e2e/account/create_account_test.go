//+build e2e

package api_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"

	"github.com/screwyprof/form3api"
	"github.com/screwyprof/form3api/assert"
)

// TODO: make configurable via ENV variables
const baseURL = "http://localhost:8080/v1"

func TestCreateAccount(t *testing.T) {
	// arrange
	ID := gofakeit.UUID()
	organisationID := gofakeit.UUID()
	accountNumber := gofakeit.Numerify("#########")
	bankID := gofakeit.Numerify("######")

	want := &form3api.Account{
		AccountData: form3api.AccountData{
			ID: ID,
			OrganisationID: organisationID,
			Type: "accounts",
			Attributes: &form3api.AccountAttributes{
				AccountNumber: accountNumber,
				BankID: bankID,
				BankIDCode: "GBDSC",
				BIC: "NWBKGB42",
				Country: "GB",
				Currency: "GBP",
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
			ID: ID,
			OrganisationID: organisationID,
			Type: "accounts",
			Attributes: &form3api.AccountAttributes{
				AccountNumber: accountNumber,
				BankID: bankID,
				BankIDCode: "GBDSC",
				Country: "GB",
				Currency: "GBP",
				CustomerID: "234",
				IBAN: "GB28NWBK40030212764204",
				BIC: "NWBKGB42",
				ConfirmationOfPayee: &form3api.ConfirmationOfPayee{
					AccountClassification: "Personal",
				},
			},
		},
	}

	// act
	c := form3api.NewClient(nil, baseURL)
	acc, err := c.CreateAccount(context.Background(), r)

	// assert
	assert.Ok(t, err)
	assertAccount(t, want, acc)
}

func assertAccount(tb testing.TB, want, got *form3api.Account) {
	tb.Helper()

	assertIBAN(tb, got.AccountData.Attributes.IBAN)
	assertCustomer(tb, got.AccountData.Attributes.CustomerID)

	got.AccountData.CreatedOn = nil
	got.AccountData.ModifiedOn = nil
	got.AccountData.Attributes.CustomerID = ""
	got.AccountData.Attributes.IBAN = ""
	assert.Equals(tb, want, got)
}

// TODO: should check against the pattern
func assertIBAN(tb testing.TB, IBAN string) {
	tb.Helper()
	assert.True(tb, IBAN != "")
}

// TODO: should check against the pattern
func assertCustomer(tb testing.TB, customer string) {
	tb.Helper()
	assert.True(tb, customer != "")
}