package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAccount(t *testing.T) {
	_, err := NewAccount("Kylian", "Mbappe", "hunter")
	assert.Nil(t, err)
}
