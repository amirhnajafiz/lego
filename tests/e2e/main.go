package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
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
	uri := url.URL{
		Host: fmt.Sprintf("%s/v1/new", baseURL),
	}

	uri.Query().Add("key", ktest)
	uri.Query().Add("value", vtest)

	req, _ := http.NewRequest(http.MethodPost, uri.String(), nil)
	client := http.Client{}

	if resp, err := client.Do(req); err != nil {
		log.Println(err)

		return errResponse
	} else {
		if resp.StatusCode != 200 {
			return errServer
		}
	}

	return nil
}

func TestGet() error {
	uri := url.URL{
		Host: fmt.Sprintf("%s/v1/get", baseURL),
	}

	uri.Query().Add("key", ktest)

	req, _ := http.NewRequest(http.MethodGet, uri.String(), nil)
	client := http.Client{}

	if resp, err := client.Do(req); err != nil {
		log.Println(err)

		return errResponse
	} else {
		if resp.StatusCode != 200 {
			return errServer
		}

		data := make([]byte, 0)
		if _, err := resp.Body.Read(data); err == nil {
			if string(data) != vtest {
				return errResponse
			}
		}
	}

	return nil
}

func TestDelete() error {
	uri := url.URL{
		Host: fmt.Sprintf("%s/v1/del", baseURL),
	}

	uri.Query().Add("key", ktest)

	req, _ := http.NewRequest(http.MethodDelete, uri.String(), nil)
	client := http.Client{}

	if resp, err := client.Do(req); err != nil {
		log.Println(err)

		return errResponse
	} else {
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
}
