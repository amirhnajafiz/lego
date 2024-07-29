package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

// create base variables
var (
	baseURL = "http://localhost:8080"

	ktest = "test-k"
	vtest = "test-v"

	errServer   = errors.New("server failed to respond")
	errResponse = errors.New("unexpected response")
)

// defining our testing types
type testingFunc func() error

// create testing methods
func TestPost() error {
	uri := fmt.Sprintf("%s/v1/new", baseURL)
	req, err := http.NewRequest(http.MethodPost, uri, nil)
	if err != nil {
		panic(err)
	}

	q := req.URL.Query()
	q.Add("key", ktest)
	q.Add("value", vtest)
	req.URL.RawQuery = q.Encode()

	client := http.Client{}

	if resp, err := client.Do(req); err != nil {
		log.Println(err)

		return errResponse
	} else {
		log.Println(uri, resp.StatusCode)
		if resp.StatusCode != 200 {
			return errServer
		}
	}

	return nil
}

func TestGet() error {
	uri := fmt.Sprintf("%s/v1/get", baseURL)
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		panic(err)
	}

	q := req.URL.Query()
	q.Add("key", ktest)
	req.URL.RawQuery = q.Encode()

	client := http.Client{}

	if resp, err := client.Do(req); err != nil {
		log.Println(err)

		return errResponse
	} else {
		log.Println(uri, resp.StatusCode)
		if resp.StatusCode != 200 {
			return errServer
		}
	}

	return nil
}

func TestDelete() error {
	uri := fmt.Sprintf("%s/v1/del", baseURL)
	req, err := http.NewRequest(http.MethodDelete, uri, nil)
	if err != nil {
		panic(err)
	}

	q := req.URL.Query()
	q.Add("key", ktest)
	req.URL.RawQuery = q.Encode()

	client := http.Client{}

	if resp, err := client.Do(req); err != nil {
		log.Println(err)

		return errResponse
	} else {
		log.Println(uri, resp.StatusCode)
		if resp.StatusCode != 200 {
			return errServer
		}
	}

	return nil
}

func main() {
	methods := []testingFunc{
		TestPost,
		TestGet,
		TestDelete,
	}

	for _, method := range methods {
		if err := method(); err != nil {
			panic(err)
		}
	}

	log.Println("all tests are passed.")
}
