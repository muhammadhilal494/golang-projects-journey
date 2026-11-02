package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestAddExpense(t *testing.T) {

	input := "500\nFood\nLunch\n"

	reader := bufio.NewReader(strings.NewReader(input))

	expense, err := addExpense(reader)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if expense.Amount != 500 {
		t.Errorf("Expected amount 500, got %f", expense.Amount)
	}

	if expense.Category != "Food" {
		t.Errorf("Expected category Food, got %s", expense.Category)
	}

	if expense.Description != "Lunch" {
		t.Errorf("Expected description Lunch, got %s", expense.Description)
	}
}

func TestAddExpenseInvalidAmount(t *testing.T) {

	input := "abc\n"

	reader := bufio.NewReader(strings.NewReader(input))

	_, err := addExpense(reader)

	if err == nil {
		t.Fatal("Expected an error for invalid amount, got nil")
	}
}
