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
		return nil, err
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
	values ($1, $2, $3, $4, $5)`

	_, err := s.db.Exec(
		query,
		acc.FirstName,
		acc.LastName,
		acc.BankNumber,
		acc.Balance,
		acc.CreatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStore) GetAccounts() ([]*Account, error) {
	rows, err := s.db.Query("SELECT * FROM accounts")
	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			s.logger.Error("error closing rows", "error", err)
		}
	}(rows)

	var accounts []*Account
	for rows.Next() {
		var acc Account
		if err := rows.Scan(&acc.ID, &acc.FirstName, &acc.LastName, &acc.BankNumber, &acc.Balance, &acc.CreatedAt); err != nil {
			return nil, err
		}
		accounts = append(accounts, &acc)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return accounts, nil
}

func (s *PostgresStore) DeleteAccount(id int) error              { return nil }
func (s *PostgresStore) UpdateAccount(*Account) error            { return nil }
func (s *PostgresStore) GetAccountByID(id int) (*Account, error) { return nil, nil }
