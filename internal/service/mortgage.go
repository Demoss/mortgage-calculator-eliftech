package service

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"math"
	"mortgage-calulator-eliftech/internal/command"
	"mortgage-calulator-eliftech/internal/repository"
)

type MortgageService struct {
	repo repository.Bank
}

func NewMortgageService(repo repository.Bank) *MortgageService {
	return &MortgageService{repo: repo}
}
func (s *MortgageService) GetSuitableBanks(ctx context.Context, mortgage command.CreateMortgage) (float64, error) {
	bank, err := s.repo.GetOne(ctx, mortgage.BankName)
	if err != nil {
		return 0, err
	}
	if bank.MaxLoan <= mortgage.InitialLoan {
		return 0, errors.New("bank has not enough money to borrow")
	}
	downPayment := float64(mortgage.InitialLoan) * float64(bank.MinDownPayment) / 100
	if float64(mortgage.DownPayment) < downPayment {
		return 0, errors.New(fmt.Sprintf("your down payment is not enough,you must pay %d%%(%v) of mortgage", bank.MinDownPayment, downPayment))
	}

	percentOfPaymentPeriod := math.Pow(1.0+(float64(bank.Rate)/100/12), float64(bank.LoanTerm))
	monthBorrow := float64(mortgage.InitialLoan) * float64(bank.Rate) / 100.0 / 12.0

	monthPayment := (monthBorrow * (1 + percentOfPaymentPeriod)) / ((1 + percentOfPaymentPeriod) - 1)
	return monthPayment, nil
}
