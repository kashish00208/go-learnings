package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Http server in golang")

	resp, err := http.Get("http://example.com/form")

	if err != nil {
		println(err)
	}
	fmt.Println(resp)

}
