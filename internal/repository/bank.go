package repository

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"mortgage-calulator-eliftech/internal/command"
	"mortgage-calulator-eliftech/internal/domain"
	"strings"
)

type BankPostgres struct {
	db *sqlx.DB
}

func NewBankPostgres(db *sqlx.DB) *BankPostgres {
	return &BankPostgres{db: db}
}

func (r *BankPostgres) CreateBank(ctx context.Context, bank domain.Bank) error {
	query := fmt.Sprintf("INSERT INTO %s (name,rate,max_loan,min_down_payment,loan_term) values($1,$2,$3,$4,$5)", banks)
	err := r.db.QueryRowContext(ctx, query, bank.Name, bank.Rate, bank.MaxLoan, bank.MinDownPayment, bank.LoanTerm)
	if err != nil {
		return err.Err()
	}

	return nil
}

func (r *BankPostgres) DeleteBank(ctx context.Context, name string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE name=$1", banks)
	_, err := r.db.ExecContext(ctx, query, name)
	return err
}

func (r *BankPostgres) Update(ctx context.Context, name string, cmd command.UpdateBankRequest) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	fmt.Println(cmd)
	fmt.Println(name)
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

	query := fmt.Sprintf("UPDATE %s SET %s WHERE name=$1", banks, setQuery)

	fmt.Println(query)
	args = append(args, name)
	_, err := r.db.ExecContext(ctx, query, args...)
	return err
}

func (r BankPostgres) GetAll(ctx context.Context) ([]domain.Bank, error) {
	var res []domain.Bank

	query := fmt.Sprintf("SELECT * FROM %s", banks)

	if err := r.db.SelectContext(ctx, &res, query); err != nil {
		return nil, err
	}
	fmt.Println(res)
	return res, nil
}

func (r *BankPostgres) GetOne(ctx context.Context, name string) (domain.Bank, error) {
	var res domain.Bank
	query := fmt.Sprintf("SELECT * FROM %s WHERE name=$1", banks)
	if err := r.db.GetContext(ctx, &res, query, name); err != nil {
		return domain.Bank{}, err
	}
	return res, nil
}
