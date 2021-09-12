## Defer keyword

- defer adds the fucntion defered to stack for ex.
```go

defer fmt.Println("hello")
fmt.Println("World")
return
```
- here the `fmt.Println("hello")` pushed onto stack and executed when the `main` function returns in our case.
- It it helpful in cleanup scenarios like closing a file we can defer `file.Close()` it will be executed after that function returns due to any scenario.
## Goroutines

- function starting with keyword `go`
## Refs:     
- https://stackoverflow.com/questions/24805673/declare-and-initialize-pointer-concisely-i-e-pointer-to-int