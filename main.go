package main

import (
	"fmt"
	"time"
)

type Transaction struct {
	ID       int
	Amount   float64
	Category string
	Date     time.Time
	Type     string // "income" or "expense"
}

type Account struct {
	Name         string
	Balance      float64
	Transactions []Transaction
	nextID       int
}

func NewAccount(name string, initialBalance float64) *Account {
	return &Account{
		Name:         name,
		Balance:      initialBalance,
		Transactions: make([]Transaction, 0),
		nextID:       1,
	}
}

func (a *Account) AddIncome(amount float64, category string) {
	if amount <= 0 {
		fmt.Println("Amount must be positive")
		return
	}

	a.Balance += amount
	a.Transactions = append(a.Transactions, Transaction{
		ID:       a.nextID,
		Amount:   amount,
		Category: category,
		Date:     time.Now(),
		Type:     "income",
	})
	a.nextID++
}

func (a *Account) AddExpense(amount float64, category string) error {
	if amount <= 0 {
		return fmt.Errorf("amount must be positive")
	}
	if amount > a.Balance {
		return fmt.Errorf("insufficient balance")
	}

	a.Balance -= amount
	a.Transactions = append(a.Transactions, Transaction{
		ID:       a.nextID,
		Amount:   amount,
		Category: category,
		Date:     time.Now(),
		Type:     "expense",
	})
	a.nextID++
	return nil
}

func (a *Account) GetBalance() float64 {
	return a.Balance
}

func (a *Account) GetExpensesByCategory(category string) float64 {
	total := 0.0
	for _, t := range a.Transactions {
		if t.Type == "expense" && t.Category == category {
			total += t.Amount
		}
	}
	return total
}

func (a *Account) PrintStatement() {
	fmt.Printf("=== Account: %s ===\n", a.Name)
	fmt.Printf("Balance: $%.2f\n\n", a.Balance)
	fmt.Println("Transactions:")
	for _, t := range a.Transactions {
		fmt.Printf("[%d] %s - %s: $%.2f (%s)\n",
			t.ID, t.Date.Format("2006-01-02"), t.Category, t.Amount, t.Type)
	}
}

func main() {
	account := NewAccount("Savings", 5000)

	account.AddIncome(2000, "salary")
	account.AddExpense(500, "food")
	account.AddExpense(200, "entertainment")
	account.AddExpense(100, "food")

	account.PrintStatement()

	foodSpend := account.GetExpensesByCategory("food")
	fmt.Printf("\nTotal food expenses: $%.2f\n", foodSpend)
}
