package main

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var client *github.Client

func main() {
	initialize()

	getRepos()
	getOrgs()
	getWithoutAuth()
}

func getRepos() {
	ctx := context.Background()
	repos, _, err := client.Repositories.List(ctx, "", nil)
	check(err)
	fmt.Println(repos)
}

func getOrgs() {
	ctx := context.Background()
	orgs, _, _ := client.Organizations.List(ctx, "", nil)
	fmt.Println(orgs)
}

func getWithoutAuth() {
	ctx := context.Background()
	client := github.NewClient(nil)
	repos, _, _ := client.Repositories.List(ctx, "fenwickelliott", nil)
	fmt.Println(repos)
}

func initialize() {
	// save api token locally
	token, err := ioutil.ReadFile("token")
	check(err)
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: string(token)})
	tc := oauth2.NewClient(ctx, ts)
	client = github.NewClient(tc)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
