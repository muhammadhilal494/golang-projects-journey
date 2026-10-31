package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Expense struct {
	Amount      float64
	Category    string
	Description string
}

func readInput(reader *bufio.Reader, prompt string) (string, error) {
	fmt.Print(prompt)

	input, err := reader.ReadString('\n')

	if err != nil && err != io.EOF {
		return "", err
	}

	if err == io.EOF && input == "" {
		return "", io.EOF
	}

	input = strings.TrimSpace(input)

	return input, nil
}

func addExpense(reader *bufio.Reader) (Expense, error) {

	amountInput, err := readInput(reader, "Enter amount: ")
	if err != nil {
		return Expense{}, err
	}

	amount, err := strconv.ParseFloat(amountInput, 64)
	if err != nil {
		return Expense{}, errors.New("invalid amount: please enter a number")
	}

	if amount <= 0 {
		return Expense{}, errors.New("amount must be greater than zero")
	}

	category, err := readInput(reader, "Enter category: ")
	if err != nil {
		return Expense{}, err
	}

	if category == "" {
		return Expense{}, errors.New("category cannot be empty")
	}

	description, err := readInput(reader, "Enter description: ")
	if err != nil {
		return Expense{}, err
	}

	if description == "" {
		return Expense{}, errors.New("description cannot be empty")
	}

	expense := Expense{
		Amount:      amount,
		Category:    category,
		Description: description,
	}

	return expense, nil
}

func viewExpenses(expenses []Expense) {

	if len(expenses) == 0 {
		fmt.Println("\nNo expenses found.")
		return
	}

	fmt.Println("\n========== All Expenses ==========")

	for i, expense := range expenses {
		fmt.Printf(
			"%d. Amount: %.2f | Category: %s | Description: %s\n",
			i+1,
			expense.Amount,
			expense.Category,
			expense.Description,
		)
	}
}

func viewSummary(expenses []Expense) {

	if len(expenses) == 0 {
		fmt.Println("\nNo expenses available for summary.")
		return
	}

	var total float64

	categoryTotals := make(map[string]float64)
	categoryNames := make(map[string]string)

	for _, expense := range expenses {

		total += expense.Amount

		// Normalize category for calculations
		key := strings.ToLower(strings.TrimSpace(expense.Category))

		categoryTotals[key] += expense.Amount

		// Preserve the first original category name
		if _, exists := categoryNames[key]; !exists {
			categoryNames[key] = expense.Category
		}
	}

	fmt.Println("\n========== Spending Summary ==========")

	fmt.Printf("Total Expenses: %.2f\n", total)

	average := total / float64(len(expenses))

	fmt.Printf("Average Expense: %.2f\n", average)

	for key, amount := range categoryTotals {
		fmt.Printf("%s: %.2f\n", categoryNames[key], amount)
	}

}

func main() {

	reader := bufio.NewReader(os.Stdin)

	var expenses []Expense

	for {

		fmt.Println("\n========== Expense Tracker ==========")
		fmt.Println("1. Add Expense")
		fmt.Println("2. View Expenses")
		fmt.Println("3. View Spending Summary")
		fmt.Println("4. Exit")

		choiceInput, err := readInput(reader, "Enter your choice: ")

		if errors.Is(err, io.EOF) {
			fmt.Println("\nGoodbye!")
			return
		}

		if err != nil {
			fmt.Println("Input error:", err)
			return
		}

		choice, err := strconv.Atoi(choiceInput)

		if err != nil {
			fmt.Println("Invalid choice. Please enter a number.")
			continue
		}

		switch choice {

		case 1:
			expense, err := addExpense(reader)

			if err != nil {
				if errors.Is(err, io.EOF) {
					fmt.Println("\nInput ended. Goodbye!")
					return
				}

				fmt.Println("Error:", err)
				fmt.Println("Expense was not saved.")
				continue
			}

			expenses = append(expenses, expense)

			fmt.Println("\nExpense added successfully!")

		case 2:
			viewExpenses(expenses)

		case 3:
			viewSummary(expenses)

		case 4:
			fmt.Println("\nThank you for using Expense Tracker!")
			return

		default:
			fmt.Println("Invalid choice. Please select 1-4.")
		}
	}
}
