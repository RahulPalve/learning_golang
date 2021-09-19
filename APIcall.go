package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	arr := []int{1, 2, 3}
	for _, v := range arr {
		resp, _ := http.Get("https://reqres.in/api/users/" + strconv.Itoa(v))

		fmt.Println(resp.Body)
	}

}
