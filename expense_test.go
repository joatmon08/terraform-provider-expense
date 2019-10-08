// +build unit

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var id = "testing123"

func TestGetExpenses(t *testing.T) {
	client := Client{
		URL: "http://localhost:5001",
	}
	expenses, err := client.GetExpenses()
	assert.Nil(t, err, "Err should be nil")
	assert.Equal(t, 0, len(*expenses), "there should no expenses")
}
func TestCreateExpense(t *testing.T) {
	client := Client{
		URL: "http://localhost:5001",
	}
	expense := &Expense{
		ID:       id,
		Name:     "hotel",
		TripID:   "d7fd4bf6-aeb9-45a0-b671-85dfc4d09544",
		Cost:     1003.34,
		Currency: "US",
		Date:     "2019-08-20",
	}
	expenseID, err := client.CreateExpense(expense)
	assert.Nil(t, err, "Err should be nil")
	assert.NotEmpty(t, expenseID, "Expense ID should be populated")
}

func TestGetExpenseByID(t *testing.T) {
	client := Client{
		URL: "http://localhost:5001",
	}
	expense, err := client.GetExpenseByID(id)
	assert.Nil(t, err, "Err should be nil")
	assert.Equal(t, id, expense.ID, "Expense ID should match")
	assert.Equal(t, "hotel", expense.Name, "Expense Name should match")
}

func TestUpdateExpense(t *testing.T) {
	client := Client{
		URL: "http://localhost:5001",
	}
	expense := &Expense{
		ID:       id,
		Name:     "lunch",
		TripID:   "d7fd4bf6-aeb9-45a0-b671-85dfc4d09544",
		Cost:     1003.34,
		Currency: "US",
		Date:     "2019-08-20",
	}
	err := client.UpdateExpense(expense)
	assert.Nil(t, err, "Err should be nil")
	check, err := client.GetExpenseByID(id)
	assert.Nil(t, err, "Err should be nil")
	assert.Equal(t, "lunch", check.Name, "Expense Name should match")
}

func TestDeleteExpense(t *testing.T) {
	client := Client{
		URL: "http://localhost:5001",
	}
	err := client.DeleteExpense(id)
	assert.Nil(t, err, "Err should be nil")
}
