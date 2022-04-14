package service

import (
	"context"
	"mortgage-calulator-eliftech/internal/command"
	"mortgage-calulator-eliftech/internal/domain"
	"mortgage-calulator-eliftech/internal/repository"
)

type Bank interface {
	CreateBank(ctx context.Context, bank domain.Bank) error
	DeleteBank(ctx context.Context, name string) error
	Update(ctx context.Context, name string, cmd command.UpdateBankRequest) error
	GetAll(ctx context.Context) ([]domain.Bank, error)
	GetOne(ctx context.Context, name string) (domain.Bank, error)
}

type Mortgage interface {
	GetSuitableBanks(ctx context.Context, mortgage command.CreateMortgage) (float64, error)
}

type Service struct {
	Bank
	Mortgage
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Bank:     NewBankService(repos.Bank),
		Mortgage: NewMortgageService(repos.Bank),
	}
}
