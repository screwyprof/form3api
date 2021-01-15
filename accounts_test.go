package form3api

import (
	"github.com/screwyprof/form3api/assert"
	"testing"
)

func TestLinksNextPageNum(t *testing.T) {
	t.Run("valid next link given, number returned", func(t *testing.T) {
		t.Parallel()

		// arrange
		links := Links{
			Next: "/v1/organisation/accounts?page%5Bnumber%5D=1&page%5Bsize%5D=10",
		}

		// act
		num, err := links.NextPageNum()

		// assert
		assert.Ok(t, err)
		assert.Equals(t, uint64(1), num)
	})

	t.Run("invalid link format given, error returned", func(t *testing.T) {
		t.Parallel()

		// arrange
		links := Links{
			Next: "{}://",
		}

		// act
		_, err := links.NextPageNum()

		// assert
		assert.NotNil(t, err)
	})

	t.Run("invalid page number given, error returned", func(t *testing.T) {
		t.Parallel()

		// arrange
		links := Links{
			Next: "/v1/organisation/accounts?page%5Bnumber%5D=invalid_number&page%5Bsize%5D=10",
		}

		// act
		_, err := links.NextPageNum()

		// assert
		assert.NotNil(t, err)
	})
}
