package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex
var balance int = 600

func deposit(count int, amount int) {
	fmt.Printf("%d: Balance before deposit: %d\n", count, balance)
	mu.Lock()
	balance = balance + amount
	fmt.Printf("%d: Balance after deposit: %d\n", count, balance)
	mu.Unlock()
}

func withdraw(count int, amount int) {
	fmt.Printf("%d: Balance before withdraw: %d\n", count, balance)
	mu.Lock()
	defer mu.Unlock()
	if amount > balance {
		fmt.Printf("%d: Not enough balance: %d\n", count, balance)
	} else {
		time.Sleep(1 * time.Millisecond)
		balance = balance - amount
		fmt.Printf("%d: Balance after withdraw: %d\n", count, balance)
	}
}

func main() {
	fmt.Printf("Balance at the beginning: %d\n", balance)
	for i := 1; i <= 5; i++ {
		go withdraw(i, 500)
		go deposit(i, 200)
	}
	time.Sleep(1 * time.Second)
	fmt.Printf("Balance at the end: %d\n", balance)
}
