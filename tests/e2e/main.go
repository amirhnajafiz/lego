package main

import (
	"bufio"
	"fmt"
	"net/http"
)

var (
	baseURL     = "http://localhost:8080"
	healthPath  = "/healthz"
	metricsPath = "/metrics"
)

func main() {
	resp, err := http.Get(baseURL + healthPath)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan(); i++ {
		fmt.Println(scanner.Text())
	}

	resp, err = http.Get(baseURL + metricsPath)
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
