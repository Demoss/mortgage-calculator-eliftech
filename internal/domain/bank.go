package domain

type Bank struct {
	Name           string `json:"name" db:"name"`
	Rate           int    `json:"rate" db:"rate"`
	MaxLoan        int    `json:"maxLoan" db:"max_loan"`
	MinDownPayment int    `json:"minDownPayment" db:"min_down_payment"`
	LoanTerm       int    `json:"loanTerm" db:"loan_term"`
}
