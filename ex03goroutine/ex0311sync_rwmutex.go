package main

import (
	"fmt"
	"github.com/forseti/gowork/ex03goroutine/helper"
	"sync"
	"time"
)

type BankAccount struct {
	sync.RWMutex
	Balance int
}

func (a *BankAccount) AddBalance(amount int) {
	a.Lock()
	a.Balance += amount
	a.Unlock()
}

func (a *BankAccount) GetBalance() int {
	a.RLock()
	balance := a.Balance
	a.RUnlock()
	return balance
}

func main() {
	account := BankAccount{Balance: 0}

	helper.AsyncLoop(100, func(index int) {
		for j := 0; j < 100; j++ {
			account.AddBalance(1)
			fmt.Println(account.GetBalance())
		}
	})

	time.Sleep(5 * time.Second)
	fmt.Println("Total balance:", account.GetBalance())
}
