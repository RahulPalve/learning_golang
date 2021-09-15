## Channels
- Channel makes synchronization across goroutines easy
- Basically go channels simplify `waitGroup` 
```go
func someroutine(c chan){
    5->c //as c's work done main thread can starts it execution
}
c := make(chan int)
go someroutine(c chan)
d:=<-c // this will wait here unitl goroutine send data


```
compare above code with `goroutine_sync.go` for more clarity

-  there are two types of channels, Buffered and Unbuffered
above is unbuffered to create a buffered channel pass an extra capacity arg in `make(chan int, 2)` 
- in unbuffered channel if you are using a same goroutinr once a value is send it is bloacked cant use  to send, anothr data via same channel it will jsut wait there so add capacity which will be like no of types this channle can be used
