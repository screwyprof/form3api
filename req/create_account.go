package req

import "time"

type CreateAccount struct {
	AccountData `json:"data"`
}

type AccountData struct {
	ID             string             `json:"id"`
	OrganisationID string             `json:"organisation_id"`
	Type           string             `json:"type"`
	Version        uint64             `json:"version"`
	CreatedOn      *time.Time         `json:"created_on"`
	ModifiedOn     *time.Time         `json:"modified_on"`
	Attributes     *AccountAttributes `json:"attributes"`
}

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

type ConfirmationOfPayee struct {
	Name                    []string `json:"name"`
	AlternativeNames        []string `json:"alternative_names"`
	AccountClassification   string   `json:"account_classification"`
	JointAccount            bool     `json:"joint_account"`
	AccountMatchingOptOut   bool     `json:"account_matching_opt_out"`
	SecondaryIdentification string   `json:"secondary_identification"`
	Switched                bool     `json:"switched"`
	Status                  string   `json:"status"` // TODO: convert to an enum type holding "pending", "confirmed", "failed"
}
