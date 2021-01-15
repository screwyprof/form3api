package form3api

import (
	"net/url"
	"strconv"
	"time"
)

// CreateAccount is a request used to create an account.
type CreateAccount struct {
	AccountData `json:"data"`
}

// FetchAccount is a request used to fetch an account.
type FetchAccount struct {
	AccountID string
}

// DeleteAccount is a request used to delete an account.
type DeleteAccount struct {
	AccountID string
	Version   uint64
}

// ListAccounts specifies pagination options to List Accounts.
type ListAccounts struct {
	Page Page
}

// Page contains optional page parameters.
type Page struct {
	Number uint64
	Size   uint64
}

// Account represents an Account resource.
type Account struct {
	AccountData AccountData `json:"data"`
	Links       Links       `json:"links"`
}

// Accounts represents a list of Accounts.
type Accounts struct {
	AccountData []AccountData `json:"data"`
	Links       Links         `json:"links"`
}

// AccountData are the resource data that is the subject of the API call.
type AccountData struct {
	ID             string             `json:"id"`
	OrganisationID string             `json:"organisation_id"`
	Type           string             `json:"type"`
	Version        uint64             `json:"version"`
	CreatedOn      *time.Time         `json:"created_on"`
	ModifiedOn     *time.Time         `json:"modified_on"`
	Attributes     *AccountAttributes `json:"attributes"`
}

// Links represent HATEOAS section of the API response.
type Links struct {
	Self  string `json:"self"`
	First string `json:"first,omitempty"`
	Last  string `json:"last,omitempty"`
	Prev  string `json:"prev,omitempty"`
	Next  string `json:"next,omitempty"`
}

// NextPageNum parses the Next link and returns the next page number.
func (l Links) NextPageNum() (uint64, error) {
	u, err := url.Parse(l.Next)
	if err != nil {
		return 0, err
	}

	num, err := strconv.Atoi(u.Query().Get("page[number]"))
	if err != nil {
		return 0, err
	}

	return uint64(num), nil
}

// AccountAttributes are attributes for account resource.
type AccountAttributes struct {
	Country       string `json:"country"`
	Currency      string `json:"base_currency"`
	BankID        string `json:"bank_id"`
	BankIDCode    string `json:"bank_id_code"`
	AccountNumber string `json:"account_number"`
	BIC           string `json:"bic"`
	IBAN          string `json:"iban"`
	CustomerID    string `json:"customer_id"`
	*ConfirmationOfPayee
}

// ConfirmationOfPayee is a subset of account attributes which encapsulates the confirmation of payee.
type ConfirmationOfPayee struct {
	Name                    []string `json:"name"`
	AlternativeNames        []string `json:"alternative_names"`
	AccountClassification   string   `json:"account_classification"`
	JointAccount            bool     `json:"joint_account"`
	AccountMatchingOptOut   bool     `json:"account_matching_opt_out"`
	SecondaryIdentification string   `json:"secondary_identification"`
	Switched                bool     `json:"switched"`
	Status                  string   `json:"status"` // TODO: convert to an enum type: "pending", "confirmed", "failed"
}
