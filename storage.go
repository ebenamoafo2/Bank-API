package main

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type storage interface {
	CreateAccount(acc *Account) error
	DeleteAccount(int) error
	UpdateAccount(account *Account) error
	GetAccounts() ([]*Account, error)
	GetAccountByID(int) (*Account, error)
}

type PostgresStore struct {
	db     *sql.DB
	logger *slog.Logger
}

func NewPostgresStore() (*PostgresStore, error) {
	if err := godotenv.Load(); err != nil {

		slog.Default().Warn("No .env file found, using environment variables")
	}

	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		return nil, fmt.Errorf("DATABASE_URL not set")
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		pingErr := err // save the original error
		if closeErr := db.Close(); closeErr != nil {
			slog.Error("failed to close db after failed ping", "error", closeErr)
		}
		return nil, pingErr // always return the ping error
	}

	slog.Info("database connection established")
	return &PostgresStore{db: db,
		logger: slog.Default(),
	}, nil
}

func (s *PostgresStore) Init() error {
	return s.CreateAccountTable()
}

func (s *PostgresStore) CreateAccountTable() error {
	query := `CREATE TABLE IF NOT EXISTS accounts (
        id          SERIAL PRIMARY KEY,
        first_name  VARCHAR(50)    NOT NULL,
        last_name   VARCHAR(50)    NOT NULL,
        bank_number SERIAL         UNIQUE,
        balance     DECIMAL(19, 4) NOT NULL DEFAULT 0,
        created_at  TIMESTAMP      DEFAULT NOW()
    )`
	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) CreateAccount(acc *Account) error {
	query := `insert into accounts 
	(first_name, last_name, bank_number, balance, created_at)
	values ($1, $2, $3, $4, $5) returning id`

	return s.db.QueryRow(
		query,
		acc.FirstName,
		acc.LastName,
		acc.BankNumber,
		acc.Balance,
		acc.CreatedAt).Scan(&acc.ID)

}

func (s *PostgresStore) GetAccounts() ([]*Account, error) {
	const query = `SELECT * FROM accounts`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			s.logger.Error("failed to close database rows after querying accounts", "error", err)
		}
	}()

	var accounts []*Account
	for rows.Next() {
		account, err := scanIntoAccount(rows)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return accounts, nil
}

func (s *PostgresStore) DeleteAccount(id int) error {
	const query = `DELETE FROM accounts WHERE id = $1`
	_, err := s.db.Exec(query, id)
	return err
}

func (s *PostgresStore) UpdateAccount(*Account) error { return nil }

func (s *PostgresStore) GetAccountByID(id int) (*Account, error) {
	const query = `SELECT * FROM accounts WHERE id = $1`
	rows, err := s.db.Query(query, id)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			s.logger.Error("failed to close database rows after querying account", "error", err)
		}
	}()
	for rows.Next() {
		account, err := scanIntoAccount(rows)
		if err != nil {
			return nil, err
		}
		return account, nil
	}
	return nil, fmt.Errorf("account not found")
}

func scanIntoAccount(rows *sql.Rows) (*Account, error) {
	account := new(Account)
	err := rows.Scan(&account.ID, &account.FirstName, &account.LastName, &account.BankNumber, &account.Balance, &account.CreatedAt)
	return account, err
}
