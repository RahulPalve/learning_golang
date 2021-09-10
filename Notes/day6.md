# Stringers
- If you are familiar with Python Stringer is go's way of doing `__str__()` of python.

- Basically a type need to have `func (var_name type) String() string` which returns the string.
- We can say that `Stringer` is a interface having 
```go
type Stringer interface{
    String() string
}
```
- so any type haveing `String()` will implement Stringer interface.
- So if that type is used with any method in `fmt` package ex `Println` it will call `String()` method.
- used in `contactlist.go