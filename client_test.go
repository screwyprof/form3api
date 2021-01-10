package form3api_test

import (
	"testing"

	"github.com/screwyprof/form3api"
	"github.com/screwyprof/form3api/assert"
)

func TestNewClient(t *testing.T) {
	c := form3api.NewClient(nil)
	assert.NotNil(t, c)
}
