package form3api_test

import (
	"testing"

	"github.com/screwyprof/form3api"
)

func TestClient(t *testing.T) {
	t.Run("Can create new client", func(t *testing.T) {
		t.Parallel()
		c := form3api.NewClient()
		if c == nil {
			t.Fatal("Cannot create client")
		}
	})
}
