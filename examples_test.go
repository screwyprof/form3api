package form3api_test

import (
	"context"
	"fmt"

	"github.com/screwyprof/form3api"
)

func ExampleClient_CreateAccount() {
	c := form3api.NewClient(nil, "http://localhost:8080/v1")

	accountID := "51646a03-a52e-4e51-b405-cf2b8078c1a8"
	acc, err := c.CreateAccount(context.Background(), form3api.CreateAccount{
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
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Account created: %v\n", acc)
}

func ExampleClient_FetchAccount() {
	c := form3api.NewClient(nil, "http://localhost:8080/v1")

	r := form3api.FetchAccount{
		AccountID: "51646a03-a52e-4e51-b405-cf2b8078c1a8",
	}
	acc, err := c.FetchAccount(context.Background(), r)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Account fetched: %v\n", acc)
}

func ExampleClient_DeleteAccount() {
	c := form3api.NewClient(nil, "http://localhost:8080/v1")

	r := form3api.DeleteAccount{
		AccountID: "51646a03-a52e-4e51-b405-cf2b8078c1a8",
	}
	err := c.DeleteAccount(context.Background(), r)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Account deleted")
}

func ExampleClient_ListAccounts() {
	c := form3api.NewClient(nil, "http://localhost:8080/v1")

	r := form3api.ListAccounts{
		Page: form3api.Page{Size: 2},
	}

	var accounts []form3api.AccountData
	for {
		res, err := c.ListAccounts(context.Background(), r)
		if err != nil {
			return
		}

		accounts = append(accounts, res.AccountData...)
		if res.Links.Next == "" {
			break
		}

		num, err := res.Links.NextPageNum()
		if err != nil {
			return
		}

		r.Page.Number = num
	}

	fmt.Printf("Accounts: %v\n", accounts)
}
