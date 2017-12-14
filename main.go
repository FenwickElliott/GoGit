package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"reflect"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var client *github.Client
var ctx context.Context

func main() {
	// getWithoutAuth()
	initialize()

	// getRepos()
	// getOrgs()
	// createRepo("NewRepo", "Created by GoGit")
	deleteRepo("NewRepo")
}

func initialize() {
	// save api token locally
	token, err := ioutil.ReadFile("token")
	check(err)
	ctx = context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: string(token)})
	tc := oauth2.NewClient(ctx, ts)
	client = github.NewClient(tc)
}

func getRepos() {
	repos, _, err := client.Repositories.List(ctx, "", nil)
	check(err)
	fmt.Println(repos)
}

func getOrgs() {
	orgs, _, _ := client.Organizations.List(ctx, "", nil)
	fmt.Println(orgs)
}

func getWithoutAuth() {
	ctx := context.Background()
	client := github.NewClient(nil)
	repos, _, _ := client.Repositories.List(ctx, "fenwickelliott", nil)
	fmt.Println(repos)
}

func createRepo(name, description string) {
	repo := &github.Repository{
		Name:        github.String(name),
		Description: github.String(description),
	}
	repo, resp, err := client.Repositories.Create(ctx, "", repo)
	if err != nil {
		fmt.Println("Error: ", err)
		fmt.Println("Resp: ", resp)
		fmt.Println("Repo: ", reflect.TypeOf(repo))
	} else {
		fmt.Println("successfully created ", name)
	}
}

func deleteRepo(repoName string) {
	me, _, err := client.Users.Get(ctx, "")
	check(err)
	_, err = client.Repositories.Delete(ctx, *me.Login, repoName)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("successfully deleted", repoName)
	}
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
