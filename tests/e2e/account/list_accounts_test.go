//+build e2e

package account_test

import (
	"context"
	"testing"

	"github.com/screwyprof/form3api"
	"github.com/screwyprof/form3api/assert"
)

func TestListAccounts(t *testing.T) {
	// arrange
	accounts := createTestAccounts(t, 5)
	r := form3api.ListAccounts{
		Page: form3api.Page{Number:1, Size: 2},
	}

	want := &form3api.Accounts{
		AccountData: []form3api.AccountData{
			accounts[2].AccountData,
			accounts[3].AccountData,
		},
		Links: form3api.Links{
			Self: "/v1/organisation/accounts?page%5Bnumber%5D=1&page%5Bsize%5D=2",
			First: "/v1/organisation/accounts?page%5Bnumber%5D=first&page%5Bsize%5D=2",
			Last: "/v1/organisation/accounts?page%5Bnumber%5D=last&page%5Bsize%5D=2",
			Prev: "/v1/organisation/accounts?page%5Bnumber%5D=0&page%5Bsize%5D=2",
			Next: "/v1/organisation/accounts?page%5Bnumber%5D=2&page%5Bsize%5D=2",
		},
	}

	// annihilate
	t.Cleanup(func() {
		for i := range accounts {
			deleteTestAccount(t, accounts[i].AccountData.ID)
		}
	})

	// act
	got, err := client.ListAccounts(context.Background(), r)

	// assert
	assert.Ok(t, err)
	assert.Equals(t, want, got)
}

func TestListAllAccounts(t *testing.T) {
	// arrange
	accounts := createTestAccounts(t, 5)

	// annihilate
	t.Cleanup(func() {
		for i := range accounts {
			deleteTestAccount(t, accounts[i].AccountData.ID)
		}
	})

	r := form3api.ListAccounts{
		Page: form3api.Page{Size: 2},
	}

	// act
	var result []form3api.AccountData
	for {
		res, err := client.ListAccounts(context.Background(), r)
		assert.Ok(t, err)

		result = append(result, res.AccountData...)
		if res.Links.Next == "" {
			break
		}

		num, err := res.Links.NextPageNum()
		assert.Ok(t, err)

		r.Page.Number = num
	}

	// assert
	assert.Equals(t, len(accounts), len(result))
}
