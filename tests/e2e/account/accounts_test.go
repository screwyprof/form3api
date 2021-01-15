//+build e2e

package account_test

import (
	"context"
	"os"
	"testing"

	"github.com/brianvoe/gofakeit/v6"

	"github.com/screwyprof/form3api"
	"github.com/screwyprof/form3api/assert"
)

const defaultTestBaseURL = "http://localhost:8080/v1"

var client *form3api.Client

func init() {
	baseURL := os.Getenv("TEST_BASE_URL")
	if baseURL == "" {
		baseURL = defaultTestBaseURL
	}

	client = form3api.NewClient(nil, baseURL)
}

func assertAccountCreated(tb testing.TB, want, got *form3api.Account) {
	tb.Helper()

	assertIBAN(tb, got.AccountData.Attributes.IBAN)
	assertCustomer(tb, got.AccountData.Attributes.CustomerID)

	// there are a few options to deal with non-deterministic responses
	// the simplest - is to ignore them for now.
	// Ideally the response should be checked against a pre-defined schema.
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

func createTestAccount(tb testing.TB, r form3api.CreateAccount) *form3api.Account {
	tb.Helper()
	acc, err := client.CreateAccount(context.Background(), r)
	assert.Ok(tb, err)
	return acc
}

func createTestAccounts(tb testing.TB, number int) []*form3api.Account {
	var res []*form3api.Account
	for i := 0; i < number; i++ {
		acc := createTestAccount(tb, generateCreateAccountRequest())
		res = append(res, acc)
	}
	return res
}

func generateCreateAccountRequest() form3api.CreateAccount {
	r := form3api.CreateAccount{
		AccountData: form3api.AccountData{
			ID: gofakeit.UUID(),
			OrganisationID: gofakeit.UUID(),
			Type: "accounts",
			Attributes: &form3api.AccountAttributes{
				AccountNumber: gofakeit.Numerify("#########"),
				BankID: gofakeit.Numerify("######"),
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
	return r
}

func deleteTestAccount(tb testing.TB, ID string) {
	tb.Helper()
	err := client.DeleteAccount(context.Background(), form3api.DeleteAccount{AccountID: ID})
	assert.Ok(tb, err)
}
