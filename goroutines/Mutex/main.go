package main

import (
	"fmt"
	"sync"
	"time"
)

type BankAccount struct {
	balance int
	mutex   sync.Mutex
}

func (b *BankAccount) Deposit(amount int) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	b.balance += amount
	fmt.Println("Deposited", amount)
}

func (b *BankAccount) Withdraw(amount int) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if amount > b.balance {
		fmt.Println("Amount cannnot be withdrawn")
		return
	}

	b.balance -= amount
	fmt.Println("Amount withdrwan:", amount)
}

func (b *BankAccount) Balance() int {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	return b.balance
}

func main() {
	var wg sync.WaitGroup
	var account = &BankAccount{
		balance: 100,
	}

	for i := 1; i < 10; i++ {
		wg.Add(1)
		go func(amount int) {
			defer wg.Done()
			time.Sleep(time.Duration(amount) * time.Millisecond)
			account.Deposit(amount)
		}(i + 1)
	}

	for i := 1; i < 5; i++ {
		wg.Add(1)
		go func(amount int) {
			defer wg.Done()
			time.Sleep(time.Duration(amount) * time.Millisecond)
			account.Withdraw(amount)
		}(i + 1)
	}

	wg.Wait()
	fmt.Println(account.balance)

}
