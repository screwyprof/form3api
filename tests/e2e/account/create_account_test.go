//+build e2e

package account_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"

	"github.com/screwyprof/form3api"
)

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
	acc, err := client.CreateAccount(context.Background(), r)

	// assert
	form3api.Ok(t, err)
	assertAccountCreated(t, want, acc)
}
