package main

import (
	"bufio"
	"fmt"
	"net/http"
)

var home = "http://127.0.0.1:8080"
var bind = "http://127.0.0.1:8080/bind?name="

func main() {
	resp, err := http.Get(home)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan(); i++ {
		fmt.Println(scanner.Text())
	}

	resp, err = http.Get(bind + "Amir")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	scanner = bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan(); i++ {
		fmt.Println(scanner.Text())
	}
}
