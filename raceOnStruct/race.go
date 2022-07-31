package main

import (
	"fmt"
	"time"
)

type Account struct {
	balance int
}

func (a *Account) deposit(amount int) {
	a.balance += amount
}

func (a *Account) withdraw(amount int) {
	if amount > a.balance {
		fmt.Printf("Not enough balance: %d\n", a.balance)
	} else {
		time.Sleep(10 * time.Millisecond)
		a.balance -= amount
	}
}

func main() {
	var account *Account = &Account{500}
	fmt.Printf("Balance at the beginning: %d\n", account.balance)
	for i := 1; i <= 5; i++ {
		go doShopping(i, account, 800)
		go transferTo(i, account, 400)
	}
	time.Sleep(1 * time.Second)
	fmt.Printf("Balance at the end: %d\n", account.balance)
}

func doShopping(count int, a *Account, expense int) {
	fmt.Printf("%d: Balance before shopping: %d\n", count, a.balance)
	a.withdraw(expense)
	fmt.Printf("%d: Balance after shopping: %d\n", count, a.balance)
}

func transferTo(count int, a *Account, amount int) {
	a.deposit(amount)
	fmt.Printf("%d: Balance after transfer: %d\n", count, a.balance)
}
