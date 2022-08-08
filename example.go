package main

import (
	"fmt"

	"github.com/staceybrodsky/go-httpclient/gohttp"
)

var (
	githubHttpClient = getGithubClient()
)

func getGithubClient() gohttp.Client {
	client := gohttp.NewBuilder().
		DisableTimeouts(true).
		SetMaxIdleConnections(5).
		Build()

	return client
}

func main() {
	getUrls()
}

type Data struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Login     string `json:"login"`
	Name      string `json:"name"`
}

func getUrls() {
	response, err := githubHttpClient.Get("http://api.github.com/users/staceybrodsky", nil)
	if err != nil {
		panic(err)
	}

	var data Data
	if err := response.UnmarshalJson(&data); err != nil {
		panic(err)
	}
	fmt.Println(data.Login)
	fmt.Println(response.StatusCode())
	fmt.Println(response.String())
}
