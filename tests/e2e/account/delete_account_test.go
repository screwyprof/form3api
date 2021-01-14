//+build e2e

package account_test

import (
	"context"
	"testing"

	"github.com/screwyprof/form3api"
)

func TestDeleteAccount(t *testing.T) {
	// arrange
	want := createTestAccount(t, generateCreateAccountRequest())

	r := form3api.DeleteAccount{
		AccountID: want.AccountData.ID,
	}

	// act
	err := client.DeleteAccount(context.Background(), r)

	// assert
	form3api.Ok(t, err)
}
