# Makefile
```shell
make -h
```

```go
go mod init github.com/ahmed82/go-makefile
```

## Manual run without `make` , `Makefile`
```go
go run main.go mutex.go
or 
go run .
```

# Mutex example
```go
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

func main() {

	fmt.Println("Mutext")
	balance = 1000

	var wg sync.WaitGroup
	wg.Add(2)
	go withdraw(700, &wg)
	go deposit(500, &wg)
	wg.Wait()

	fmt.Printf(" New balance %d\n", balance)
}
```

# Go Channel's Select operetor
```go
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
```
