package main

import "fmt"

type Expense struct {
	Amount      float64
	Category    string
	Description string
}

func addExpense() Expense {

	var amount float64
	var category string
	var description string

	fmt.Println()
	fmt.Println("--- Add Expense ---")

	fmt.Print("Enter amount: ")
	_, err := fmt.Scanln(&amount)
	if err != nil {
		fmt.Println("Invalid amount. Please enter a number.")
		return Expense{}
	}

	fmt.Print("Enter category: ")
	fmt.Scanln(&category)

	fmt.Print("Enter description: ")
	fmt.Scanln(&description)

	expense1 := Expense{
		Amount:      amount,
		Category:    category,
		Description: description,
	}

	// expenses = append(expenses, expense1)

	// expense2 := Expense{
	// Amount:      500,
	// Category:    "Food",
	// Description: "Lunch",
	// }

	// expenses = append(expenses, expense2)

	fmt.Println()
	fmt.Println("Expense recorded:")

	return expense1
	// fmt.Println("Total expenses:", len(expenses))
	// fmt.Println("Expenses: ",expenses)
	// fmt.Println("Amount:", expenses.Amount)
	// fmt.Println("Category:", expenses.Category)
	// fmt.Println("Description:", expense.Description)
}

func main() {
	var expenses []Expense
	for {
		fmt.Println("==============================")
		fmt.Println("   PERSONAL EXPENSE TRACKER")
		fmt.Println("==============================")
		fmt.Println("1. Add expense")
		fmt.Println("2. View expenses")
		fmt.Println("3. View spending summary")
		fmt.Println("4. Exit")
		fmt.Println("==============================")

		var choice int

		fmt.Print("Choose an option: ")
		_, err := fmt.Scanln(&choice)

		if err != nil {
			fmt.Println("Please enter a valid number")
			return
		}

		switch choice {

		case 1:
			expense := addExpense()
			expenses = append(expenses, expense)

			fmt.Println("Total expenses:", len(expenses))
		case 2:
			if len(expenses) == 0 {
				fmt.Println("No expenses found.")
			} else {
				fmt.Println("\n--- All Expenses ---")

				for i, expense := range expenses {
					fmt.Println("Expense:", i+1)
					fmt.Println("Amount:", expense.Amount)
					fmt.Println("Category:", expense.Category)
					fmt.Println("Description:", expense.Description)
					fmt.Println("--------------------")
				}
			}
		case 3:
			fmt.Println("View spending summary")
		case 4:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}
