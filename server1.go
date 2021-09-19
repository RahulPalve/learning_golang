package main

import (
	"fmt"
	"net/http"
)

func handler(resp http.ResponseWriter, req *http.Request) {
	// actualresp = resp.(http.Response) //https://tour.golang.org/methods/15
	fmt.Printf("%T", resp) //This will print underlying interface which is http.response not http.Response
	resp.Write([]byte("Hello Web, from go"))
	// resp.Write([]byte("what will happen after this?"))
}
func main() {
	http.HandleFunc("/api", handler)
	http.ListenAndServe(":8000", nil)
}
