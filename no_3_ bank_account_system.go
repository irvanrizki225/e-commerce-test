package main

import (
	"fmt"
	"sync"
)

type BankAccount struct {
	balance int
	mu      sync.Mutex
}

func (acc *BankAccount) Deposit(amount int) {
	acc.mu.Lock()
	defer acc.mu.Unlock()
	acc.balance += amount
}

func (acc *BankAccount) Withdraw(amount int) bool {
	acc.mu.Lock()
	defer acc.mu.Unlock()
	if acc.balance >= amount {
		acc.balance -= amount
		return true
	}
	return false
}

func runBankAccountSystem() {
	account := BankAccount{balance: 1000}
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		account.Deposit(500)
		wg.Done()
	}()

	go func() {
		account.Withdraw(300)
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Final Balance:", account.balance)
}
