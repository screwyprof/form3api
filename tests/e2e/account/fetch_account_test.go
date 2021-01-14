//+build e2e

package account_test

import (
	"context"
	"testing"

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
