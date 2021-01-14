//+build e2e

package account_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"

	"github.com/screwyprof/form3api"
)

func TestFetchAccount(t *testing.T) {
	// arrange
	want, err := createTestAccount(generateCreateAccountRequest())
	form3api.Ok(t, err)

	r := form3api.FetchAccount{
		AccountID: want.AccountData.ID,
	}

	// act
	acc, err := client.FetchAccount(context.Background(), r)

	// assert
	form3api.Ok(t, err)
	form3api.Equals(t, want, acc)
}

func createTestAccount(r form3api.CreateAccount) (*form3api.Account, error) {
	return client.CreateAccount(context.Background(), r)
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


