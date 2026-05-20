package main

import (
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type LoginResponse struct {
	Number int64  `json:"number"`
	Token  string `json:"token"`
}

type LoginRequest struct {
	BankNumber int64  `json:"bank_number"`
	Password   string `json:"password"`
}
type TransferRequest struct {
	ToAccount int     `json:"toAccount"`
	Amount    float64 `json:"amount"`
}
type Account struct {
	ID           int       `json:"id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	BankNumber   int64     `json:"bank_number"`
	Passwordhash string    `json:"-"`
	Balance      float64   `json:"balance"`
	CreatedAt    time.Time `json:"created_at"`
}

type CreateAccountRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}

func (a *Account) ValidPassword(pw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(a.Passwordhash), []byte(pw)) == nil
}

func NewAccount(firstname, lastname, Password string) (*Account, error) {
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &Account{
		FirstName:    firstname,
		LastName:     lastname,
		Passwordhash: string(hashBytes),
		BankNumber:   int64(rand.Intn(1000000)),
		CreatedAt:    time.Now().UTC(),
	}, nil
}
