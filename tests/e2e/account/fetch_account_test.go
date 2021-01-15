//+build e2e

package account_test

import (
	"context"
	"testing"

	"github.com/screwyprof/form3api"
	"github.com/screwyprof/form3api/assert"
)

func TestFetchAccount(t *testing.T) {
	// arrange
	want := createTestAccount(t, generateCreateAccountRequest())
	r := form3api.FetchAccount{
		AccountID: want.AccountData.ID,
	}

	// annihilate
	t.Cleanup(func() {
		deleteTestAccount(t, want.AccountData.ID)
	})

	// act
	acc, err := client.FetchAccount(context.Background(), r)

	// assert
	assert.Ok(t, err)
	assert.Equals(t, want, acc)
}
