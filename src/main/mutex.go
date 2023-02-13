package main

import (
	"fmt"
	"sync"
)

var (
	mutex   sync.Mutex
	balance int
)

func withdraw(value int, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Printf("Withdrawing %d to account with balance %d\n ", value, balance)
	balance -= value
	mutex.Unlock()
	wg.Done()

}

func deposit(value int, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Printf("Depositing %d to account with balance %d\n ", value, balance)
	balance += value
	mutex.Unlock()
	wg.Done()
}

func mainmutix() {

	fmt.Println("Mutex Example")
	balance = 1000

	var wg sync.WaitGroup
	wg.Add(2)
	go withdraw(700, &wg)
	go deposit(500, &wg)
	wg.Wait()

	fmt.Printf(" New balance %d\n", balance)
}
