package main

import (
	"fmt"

	"github.com/parnurzeal/gorequest"
)

func printStatus(resp gorequest.Response, body string, errs []error) {
	fmt.Println(resp.Status)
}

func main() {
	request := gorequest.New()

	_, body, errs := request.Get(
		"https://test161.ops-class.org/api-v1/submissions/123").
		Send(`{"ID":1, "name":"name"}`).
		End(printStatus)
	if errs != nil {
		panic(errs)
	}

	fmt.Println(body)

	_, body, errs = request.Post(
		"https://test161.ops-class.org/api-v1/submissions").
		Send(`{"ID":1, "name":"name"}`).
		End(printStatus)

	if errs != nil {
		panic(errs)
	}

	fmt.Println(body)
}
