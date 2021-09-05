# Notes


## JSON, IO and private public

- IOutil provide high level intrefcae than os module, ioutil.ReadFile(str) reads entire file, while os.Open() returns Fpointer which can then used for read write and need to be closed file.Read([]bytes) read only chunk as []bytes size. func OpenFile(name string, flag int, perm FileMode) (*File, error).
- lowercase varibles define private member while first Capital public hence
```
type Contact1 struct{
	Name string
	Number int
}

type Contact2 struct{
	name string
	number int
}


var c1 Contact1
var c2 Contact2

```

c2.name cannot accessed returns 0 vlue

- json.Unmarshal([]bytes, &to_var) reads from bytes and store data to address pointed by second arg



- range syntax
```for i, v := range(contact_list){
		//contact_list[i] == v
	}
```

- make and new allocates memory, new returns pointer to type passed, while make return exact type
- or maps, slices can be initialised to avoid make or new

- struct can have method by 
```func (c Contact) DisplayPrivate(){
		fmt.Println(c.Name)
		fmt.Println(c.Number)
}
```
whre the arg before function name is called **receiver argument**

- confusing is that iven Public method of stuct cant access its private members
## Literals
- Literals ""Programming languages use the word "Literal" when referring to syntactic ways to construct some data structure. It means it's not constructed by creating an empty one and adding or subtracting as you go.""
In go simply assigning values to types speicifically to composite type have some different syntax which is called as Literals,
  Ideally map should be initiated like 
  ```
  var hashtable map[string] string = {
		"rahul":"palve"}
		```
The reason this doesnt work is usual `{}` doest make sense to go as map, when we `var x int` it is just declaration not mem allocation, below syntax allocation mem dusring initialization
  ```
  var hashtable =  map[string] string{
		"name":"rahul"} //map literal
  q := []int{2, 3, 5, 7, 11, 13} //slice literal

  ```
It is better to use `make` or 	`new` instead this little confusing syntax

- Clearning doubts regardings Literal
`var x int = 10` is valid
`var x = int 10` is invalid
but `var q = []int{2, 3, 5, 7, 11, 13}` is valid as its a literal syntax
`var q []int = {1,2,3}` is invalid
Literal syntax is valid for **map, slice, array, struct**

literal syntax usua;;y is like
`var var_name = type{data_according_to_type}`
general syntax is `var var_name type`
to allocate mem use `make` or `new`

## Other

- 



























