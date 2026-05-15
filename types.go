package main

import (
	"math/rand"
	"time"
)

type TransferRequest struct {
	ToAccount int     `json:"toAccount"`
	Amount    float64 `json:"amount"`
}
type Account struct {
	ID         int       `json:"id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	BankNumber int64     `json:"bank_number"`
	Balance    float64   `json:"balance"`
	CreatedAt  time.Time `json:"created_at"`
}

type CreateAccountRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func NewAccount(firstname, lastname string) (*Account, error) {
	return &Account{
		FirstName:  firstname,
		LastName:   lastname,
		BankNumber: int64(rand.Intn(1000000)),
		CreatedAt:  time.Now().UTC(),
	}, nil
}
