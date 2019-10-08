package main

import (
	"errors"
	"os"
)

const (
	expenseURLKey = "EXPENSE_URL"
)

type Config struct {
	Client *Client
}

func (c *Config) LoadAndValidate() error {
	expenseURL := os.Getenv(expenseURLKey)
	if len(expenseURL) == 0 {
		return errors.New("define EXPENSE_URL environment variable")
	}
	c.Client = &Client{
		URL: expenseURL,
	}
	return nil
}
