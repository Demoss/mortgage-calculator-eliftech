package command

type UpdateBankRequest struct {
	Name           *string `json:"name"`
	Rate           *int    `json:"rate"`
	MaxLoan        *int    `json:"maxLoan"`
	MinDownPayment *int    `json:"minDownPayment"`
	LoanTerm       *int    `json:"loanTerm"`
}
