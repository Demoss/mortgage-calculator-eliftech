package command

type CreateMortgage struct {
	InitialLoan int    `json:"initialLoan"`
	DownPayment int    `json:"downPayment"`
	BankName    string `json:"bankName"`
}
