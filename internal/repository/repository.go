package repository

import (
	"context"
	"gorm.io/gorm"
	"mortgage-calulator-eliftech/internal/command"
	"mortgage-calulator-eliftech/internal/domain"
)

const (
	banks = "banks"
)

type Bank interface {
	CreateBank(ctx context.Context, bank domain.Bank) error
	DeleteBank(ctx context.Context, name string) error
	Update(ctx context.Context, name string, cmd command.UpdateBankRequest) error
	GetAll(ctx context.Context) ([]domain.Bank, error)
	GetOne(ctx context.Context, name string) (domain.Bank, error)
}

type Repository struct {
	Bank
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Bank: NewBankPostgres(db),
	}
}
