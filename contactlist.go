package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type Contact struct{
	Name string
	Number int
}
	
func (c Contact) DisplayPrivate(){
		fmt.Println(c.Name)
		fmt.Println(c.Number)
}

func main(){

	var contact_list  []Contact

	fmt.Println("\tCONTACT LIST")
	fmt.Println("1 Add Contact")
	fmt.Println("2 Find Contact")

	barr, _ := ioutil.ReadFile("contacts.json")

	json.Unmarshal(barr, &contact_list)

	for _, v := range(contact_list){
		v.DisplayPrivate()
	}

}