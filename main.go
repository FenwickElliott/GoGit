package main

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func main() {
	client := initialize()
	fmt.Println(client)
}

func initialize() *github.Client {
	// save api token locally
	token, err := ioutil.ReadFile("token")
	check(err)
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: string(token)})
	tc := oauth2.NewClient(ctx, ts)
	return github.NewClient(tc)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
