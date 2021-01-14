package main

import (
	"context"
	"fmt"
	"os"

	"github.com/screwyprof/form3api"
)

const defaultBaseURL = "http://localhost:8080/v1"

func main() {
	// create client instance
	c := form3api.NewClient(nil, defaultBaseURL)

	accountID := "51646a03-a52e-4e51-b405-cf2b8078c1a8"

	// create an account
	createdAccount, err := c.CreateAccount(context.Background(), form3api.CreateAccount{
		AccountData: form3api.AccountData{
			ID:             accountID,
			OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
			Type:           "accounts",
			Attributes: &form3api.AccountAttributes{
				AccountNumber: "10000004",
				BankID:        "400302",
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
	})
	failOnError(err)
	fmt.Println("Account Created:")
	printAccount(createdAccount)

	// fetch an account
	fetchedAccount, err := c.FetchAccount(context.Background(), form3api.FetchAccount{AccountID: accountID})
	failOnError(err)
	fmt.Println("Account Fetched:")
	printAccount(fetchedAccount)

	// Output:
	// Account Created:
	// Account ID: 51646a03-a52e-4e51-b405-cf2b8078c1a8
	// Organisation ID: eb0bd6f5-c3f5-44b2-b677-acd23cdde73c
	// Account Number: 10000004
	// IBAN: GB28NWBK40030212764204
	//
	// Account Fetched:
	// Account ID: 51646a03-a52e-4e51-b405-cf2b8078c1a8
	// Organisation ID: eb0bd6f5-c3f5-44b2-b677-acd23cdde73c
	// Account Number: 10000004
	// IBAN: GB28NWBK40030212764204
}

func printAccount(acc *form3api.Account) {
	fmt.Printf("Account ID: %s\n", acc.AccountData.ID)
	fmt.Printf("Organisation ID: %s\n", acc.AccountData.OrganisationID)
	fmt.Printf("Account Number: %s\n", acc.AccountData.Attributes.AccountNumber)
	fmt.Printf("IBAN: %s\n", acc.AccountData.Attributes.IBAN)
	fmt.Println()
}

func failOnError(err error) {
	if err != nil {
		fmt.Printf("An error occurred: %v\n", err)
		os.Exit(1)
	}
}
