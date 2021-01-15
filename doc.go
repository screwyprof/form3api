/*
This package implements a client to access Form3 APIs.
Usage:

	import "github.com/screwyprof/form3api"

	// Create a new Client instance, then call the corresponding methods to get what you want, for example:
	c := form3api.NewClient(nil, "http://localhost:8080/v1")

	accountID := "51646a03-a52e-4e51-b405-cf2b8078c1a8"
	acc, err := c.FetchAccount(context.Background(), form3api.FetchAccount{AccountID: accountID})
*/
package form3api
