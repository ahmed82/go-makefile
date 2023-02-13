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