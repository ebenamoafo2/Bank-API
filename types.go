package main

import (
	"math/rand"
	"time"
)

type Account struct {
	ID         int       `json:"id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	BankNumber int       `json:"bank_number"`
	Balance    int64     `json:"balance"`
	CreatedAt  time.Time `json:"createdAt"`
}

type CreateAccountRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func NewAccount(firstname, lastname string) *Account {
	return &Account{
		FirstName:  firstname,
		LastName:   lastname,
		BankNumber: rand.Intn(10),
		Balance:    0,
		CreatedAt:  time.Now().UTC(),
	}
}
