//+build e2e

package api_test

import (
	"testing"

	"github.com/screwyprof/form3api"
	"github.com/screwyprof/form3api/assert"
	"github.com/screwyprof/form3api/req"
	"github.com/screwyprof/form3api/resp"
)

// TODO: make configurable via ENV variables
const baseURL = "http://localhost:8080/v1"

func TestCreateAccount(t *testing.T) {
	t.Skip("WIP")

	// arrange
	want := &resp.Account{
		AccountData: resp.AccountData{
			ID: "0fa6dd82-4e7e-4e82-9049-fc86cf52d434",
			OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
			Attributes: &resp.AccountAttributes{
				AccountNumber: "10000004",
				BankID: "400302",
				BankIDCode: "GBDSC",
				BIC: "NWBKGB42",
				Country: "GB",
				Currency: "GBP",
				ConfirmationOfPayee: &resp.ConfirmationOfPayee{
					AccountClassification: "Personal",
				},
			},
		},
		Links: resp.Links{
			Self: "/v1/organisation/accounts/0fa6dd82-4e7e-4e82-9049-fc86cf52d434",
		},
	}

	r := req.CreateAccount{
				AccountData: req.AccountData{
					ID: "0fa6dd82-4e7e-4e82-9049-fc86cf52d434",
					OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
					Type: "accounts",
					Attributes: &req.AccountAttributes{
						Country: "GB",
						Currency: "GBP",
						BankID: "400302",
						BankIDCode: "GBDSC",
						AccountNumber: "10000004",
						CustomerID: "234",
						IBAN: "GB28NWBK40030212764204",
						BIC: "NWBKGB42",
						ConfirmationOfPayee: &req.ConfirmationOfPayee{
							AccountClassification: "Personal",
						},
					},
				},
			}

	// act
	c := form3api.NewClient(nil)
	acc, err := c.CreateAccount(r)

	// assert
	assert.Ok(t, err)
	assert.Equals(t, want, acc)
}
