//+build e2e

package account_test

import (
	"os"
	"testing"

	"github.com/screwyprof/form3api"
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
	form3api.Equals(tb, want, got)
}

// TODO: should check against the pattern
func assertIBAN(tb testing.TB, IBAN string) {
	tb.Helper()
	form3api.True(tb, IBAN != "")
}

// TODO: should check against the pattern
func assertCustomer(tb testing.TB, customer string) {
	tb.Helper()
	form3api.True(tb, customer != "")
}