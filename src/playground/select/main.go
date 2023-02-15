package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mutex   sync.Mutex
	balance int
)

func main() {

	one := make(chan string)
	two := make(chan string)

	quit := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		one <- "Hey From Channel One"
	}()

	go func() {
		time.Sleep(3 * time.Second)
		two <- "Hello From Channel Two"
	}()

	select {
	case res1 := <-one:
		fmt.Println("I get from channel One!", res1)
	case res2 := <-two:
		fmt.Println("I get from channel Two!", res2)
		// if any value has been pushed to the quit channel then it is true for this case
		// amd it will quit the select by running the return
	case <-quit:
		return

	}

}
