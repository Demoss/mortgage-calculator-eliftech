package command

type CreateBankRequest struct {
	Name           string `json:"name" binding:"required"`
	Rate           int    `json:"rate" binding:"required"`
	MaxLoan        int    `json:"maxLoan" binding:"required"`
	MinDownPayment int    `json:"minDownPayment" binding:"required"`
	LoanTerm       int    `json:"loanTerm" binding:"required"`
}
