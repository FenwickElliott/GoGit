package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
)

func main() {
	client := github.NewClient(nil)
	ctx := context.Background()
	orgs, _, err := client.Organizations.List(ctx, "willnorris", nil)
	check(err)

	fmt.Println(orgs)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
