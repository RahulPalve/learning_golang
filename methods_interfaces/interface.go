package main

import "fmt"

type OSInterface interface {
	shutdown() bool
}

type Win10 struct {
	Logo string
}

type Linux struct {
	Logo string
}

func (w Win10) shutdown() bool {
	return false
}

func main() {
	var os OSInterface
	os = Win10{Logo: "Windows"}
	fmt.Println(os.shutdown())
	//fmt.Println(os.Logo) //type OSInterface has no field or method Logo

	win10 := os.(Win10) //to get underlying value of type
	fmt.Println(win10, win10.Logo)
}
