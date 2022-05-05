package repository

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"mortgage-calulator-eliftech/internal/command"
	"mortgage-calulator-eliftech/internal/domain"
	"strings"
)

type BankPostgres struct {
	db *gorm.DB
}

func NewBankPostgres(db *gorm.DB) *BankPostgres {
	return &BankPostgres{db: db}
}

func (r *BankPostgres) CreateBank(ctx context.Context, bank domain.Bank) error {
	if res := r.db.WithContext(ctx).Create(&bank); res.Error != nil {
		errors.New(fmt.Sprintf("%s", gorm.ErrInvalidData))
	}

	return nil
}

func (r *BankPostgres) DeleteBank(ctx context.Context, name string) error {
	if err := r.db.WithContext(ctx).Delete(&name); err != nil {
		errors.Is(err.Error, gorm.ErrRecordNotFound)
	}

	return nil
}

func (r *BankPostgres) Update(ctx context.Context, name string, cmd command.UpdateBankRequest) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if cmd.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *cmd.Name)
		argId++
	}
	if cmd.Rate != nil {
		setValues = append(setValues, fmt.Sprintf("rate=$%d", argId))
		args = append(args, *cmd.Rate)
		argId++
	}
	if cmd.MaxLoan != nil {
		setValues = append(setValues, fmt.Sprintf("max_loan=$%d", argId))
		args = append(args, *cmd.MaxLoan)
		argId++
	}
	if cmd.MinDownPayment != nil {
		setValues = append(setValues, fmt.Sprintf("min_down_payment=$%d", argId))
		args = append(args, *cmd.MinDownPayment)
		argId++
	}
	if cmd.LoanTerm != nil {
		setValues = append(setValues, fmt.Sprintf("loan_term=$%d", argId))
		args = append(args, *cmd.LoanTerm)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE name=$%d", banks, setQuery, argId)

	args = append(args, name)
	if res := r.db.WithContext(ctx).Exec(query, args...); res.Error != nil {
		return res.Error
	}

	return nil
}

func (r BankPostgres) GetAll(ctx context.Context) ([]domain.Bank, error) {
	var res []domain.Bank

	data := r.db.WithContext(ctx).Table("banks").Find(&res)
	if data.Error != nil {
		return nil, data.Error
	}

	return res, nil
}

func (r *BankPostgres) GetOne(ctx context.Context, name string) (domain.Bank, error) {
	var res domain.Bank

	r.db.WithContext(ctx).Where("name = ?", name).Find(&res)

	return res, nil

}
