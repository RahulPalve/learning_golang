package main

import "fmt"
import "unsafe"

type Request struct{
	// Remember this is not var declaration dont use `var url string` here
	url string
	payload map[string] string
	timeout int
}

type Response struct{
	body string
}

func (r Request) get() Response{
	
	body := Response{body:"I am feeling lucky"}//it is not key-val so use var directly 
	return body
}

// To modify data in struct a pointer reciver is necessary, otherwise changes will be jst 
// in function scope
func (r *Request) changeUrl(url string ){
	(*r).url = url

}

func main(){
	var someApiCall Request //memory allocated here 

	payload := map[string]string{
		"Token": "skjdoh3ekjce2",
	}
	
	someApiCall = Request{url: "http://google.com", payload: payload}
	fmt.Println(someApiCall)
	fmt.Println(someApiCall.get())

	someApiCall.changeUrl("http://yahoo.com")
	fmt.Println(someApiCall)

	
	fmt.Printf("%p %d",&someApiCall, unsafe.Sizeof(someApiCall))


}