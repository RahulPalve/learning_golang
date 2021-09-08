# Interfaces

 - In order to understand, first understand what is interface in OOP, and then try to differentiate how go do it in diff manner.
- In my first understaning interfaces allows us to be a little bit type independent, for ex. we have `win10, linux, mac` types we have to implent a method `shutdown()`, we can create a interface OS and just do it like `var os1 OS = make(win10)` so now after that we dont care what is underlying type because we are aware of interface methods we can direcly call them. So interfaces provide a polymorphism. and abstraction to type.
- Naming convention or is suffix `Interface, er` or prefix `i`
- As empty interface does not have any method, so every type implements emty intreface
- 

### Other things
- go's special if syntax
```go
if err := json.Unmarshal([]byte(input), &val); err != nil {
        panic(err)
    
```
 ## References     
 - https://www.cs.utah.edu/~germain/PPS/Topics/interfaces.html
 - https://medium.com/@dotronglong/interface-naming-convention-in-golang-f53d9f471593
 - https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go
