## Understanding Golang's pass by value

To understand how go passes var to function we need to understand how really vars store data.
let's first see how C do that
```c
int a;
int &b=a; //reference
int *c=&a; //pointer
```
References does are just alias for var
this is how refs are stored
```
addr   | x0 |  x1  |  x2  |
memory |    |   2  |      |
var    |    |  a,b |      |
```
this is for pointer. 
```
addr   | x0 |  x1  |  x2  |
memory |    |   2  |  x1  |
var    |    |   a  |   b  |
```

Now lets see how data in passed in func in C
```c
void somefunc(int *a){
    *a=2; //changes a's value in main scope
    int c=4;
    &a = c; //callers val changes
}
.
.
main(){
    int *a=10;
    somefunc(&a)
}
```
In C while doing above Operation `&a` address of a is passed
```
addr   | x0 |  x1  |  x2  | x3
memory | x1 |  10  |      | 4
var    | *a |      |      | c

```
x1 is passed to the function
after &a = c; the a var points to data block of c val 4 for c
**adress of (\*a) in main = adress of (\*a) in func**
```go
func someFunc(x *int) {
    *x = 2 // Whatever variable caller passed in will now be 2
    y := 7
    x = &y // has no impact on the caller because we overwrote the pointer value!
}
```
In go the pointer is passed by value means copied to diff location so
**adress of (\*a) in main != adress of (\*a) in func**
Remember here we are talking about adress of pntr not address stored in pntr



**Refs:**
https://dave.cheney.net/2017/04/29/there-is-no-pass-by-reference-in-go
https://stackoverflow.com/questions/47296325/passing-by-reference-and-value-in-go-to-functions





