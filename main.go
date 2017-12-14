package main

import (
	"context"
	"fmt"
	"io/ioutil"

	"golang.org/x/oauth2"

	"github.com/google/go-github/github"
)

func main() {
	// save api token locally as 'token'
	token, err := ioutil.ReadFile("token")
	check(err)

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: string(token)})
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	repos, _, err := client.Repositories.List(ctx, "", nil)
	check(err)

	fmt.Println(repos)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
