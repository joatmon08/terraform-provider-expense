package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	URL string
}

type Expense struct {
	ID       string  `json:"id,omitempty"`
	Name     string  `json:"name"`
	TripID   string  `json:"tripId"`
	Cost     float64 `json:"cost"`
	Currency string  `json:"currency"`
	Date     string  `json:"date"`
}

func (c *Client) GetExpenses() (*[]Expense, error) {
	url := fmt.Sprintf("%s/api/expense", c.URL)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var expenses []Expense
	if err := json.Unmarshal(data, &expenses); err != nil {
		return nil, err
	}
	return &expenses, nil
}

func (c *Client) GetExpenseByID(id string) (*Expense, error) {
	url := fmt.Sprintf("%s/api/expense/%s", c.URL, id)
	response, err := http.Get(url)
	if response.StatusCode == http.StatusNotFound {
		return nil, errors.New("Expense not found")
	}
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var expense Expense
	if err := json.Unmarshal(data, &expense); err != nil {
		return nil, err
	}
	return &expense, nil
}

func (c *Client) CreateExpense(expense *Expense) (string, error) {
	url := fmt.Sprintf("%s/api/expense", c.URL)
	body, err := json.Marshal(expense)
	if err != nil {
		return "", err
	}
	response, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	var created Expense
	if err := json.Unmarshal(data, &created); err != nil {
		return "", err
	}
	return created.ID, nil
}

func (c *Client) UpdateExpense(expense *Expense) error {
	url := fmt.Sprintf("%s/api/expense/%s", c.URL, expense.ID)
	body, err := json.Marshal(expense)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(body))
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return err
	}

	if _, err = http.DefaultClient.Do(req); err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteExpense(id string) error {
	url := fmt.Sprintf("%s/api/expense/%s", c.URL, id)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}
	if _, err = http.DefaultClient.Do(req); err != nil {
		return err
	}
	return nil
}
