package command

type CreateBankRequest struct {
	Name           string `json:"name"`
	Rate           int    `json:"rate"`
	MaxLoan        int    `json:"maxLoan"`
	MinDownPayment int    `json:"minDownPayment"`
	LoanTerm       int    `json:"loanTerm"`
}
