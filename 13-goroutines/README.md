1. Creating goroutines
    - Use `go` keyword in front of function call
    - When using anonymous functions, pass data as local variables
2. Synchronization
    - Use `sync.WaitGroup` to wait for groups of goroutines to complete
    - Use `Sync.Mutex` and `Sync.RWMutex` to protect data access
3. Parallelism
    - By default, Go will use CPU threads to equal to available cores
    - Change with runtime.GOMAXPROCS
    - More threads can increase performance, but too many can slow it down
4. Best Practices
    - Don't create goroutines in libraries
        - Let consumer control concurrency
    - When creating a goroutine, know how it will end
        - Avoids subtle memory leaks
    - Check for race conditions at compile time
        - go run -race src/path/to/file.go
