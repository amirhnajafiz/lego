package main

import (
	"bufio"
	"fmt"
	"net/http"
)

var (
	homeUrl = "http://127.0.0.1:8080"
	bindUrl = "http://127.0.0.1:8080/bind?name="
)

func main() {
	resp, err := http.Get(homeUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan(); i++ {
		fmt.Println(scanner.Text())
	}

	resp, err = http.Get(bindUrl + "Amir")
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
