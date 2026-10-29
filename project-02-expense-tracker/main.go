package main

import "fmt"

type Expense struct {
    Amount      float64
    Category    string
    Description string
}

func addExpense() {
    var amount float64
    var category string
    var description string

    fmt.Println()
    fmt.Println("--- Add Expense ---")

    fmt.Print("Enter amount: ")
    _, err := fmt.Scanln(&amount)
    if err != nil {
        fmt.Println("Invalid amount. Please enter a number.")
        return
    }

    fmt.Print("Enter category: ")
    fmt.Scanln(&category)

    fmt.Print("Enter description: ")
    fmt.Scanln(&description)

    expense := Expense{
        Amount:      amount,
        Category:    category,
        Description: description,
    }

    fmt.Println()
    fmt.Println("Expense recorded:")
    fmt.Println("Amount:", expense.Amount)
    fmt.Println("Category:", expense.Category)
    fmt.Println("Description:", expense.Description)
}

func main() {
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

    if choice == 1 {
        addExpense()
    } else if choice == 2 {
        fmt.Println("View expenses")
    } else if choice == 3 {
        fmt.Println("View spending summary")
    } else if choice == 4 {
        fmt.Println("Goodbye!")
    } else {
        fmt.Println("Invalid option")
    }
}