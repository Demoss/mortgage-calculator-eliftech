package service

import (
	"context"
	"errors"
	"mortgage-calulator-eliftech/internal/command"
	"mortgage-calulator-eliftech/internal/domain"
	"mortgage-calulator-eliftech/internal/repository"
)

type BankService struct {
	repo repository.Bank
}

func NewBankService(repo repository.Bank) *BankService {
	return &BankService{repo: repo}
}

func (s *BankService) CreateBank(ctx context.Context, bank domain.Bank) error {
	if bank.Rate <= 0 {
		return errors.New("rate must be more than 0")
	}
	if bank.MinDownPayment > 100 {
		return errors.New("min down payment more than loan")
	}

	return s.repo.CreateBank(ctx, bank)
}

func (s *BankService) DeleteBank(ctx context.Context, name string) error {
	return s.repo.DeleteBank(ctx, name)
}

func (s *BankService) Update(ctx context.Context, name string, cmd command.UpdateBankRequest) error {
	return s.repo.Update(ctx, name, cmd)
}

func (s *BankService) GetAll(ctx context.Context) ([]domain.Bank, error) {
	banks, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return banks, nil
}

func (s *BankService) GetOne(ctx context.Context, name string) (domain.Bank, error) {
	bank, err := s.repo.GetOne(ctx, name)
	if err != nil {
		return domain.Bank{}, err
	}
	return bank, nil
}
