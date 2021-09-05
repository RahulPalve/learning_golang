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



























)

