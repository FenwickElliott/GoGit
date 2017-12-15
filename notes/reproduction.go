package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	getRepo("GoGit", "FenwickElliott")
}

func getRepo(repo, owner string) {
	url := "https://api.github.com/repos/" + owner+ "/" + repo
	req, err := http.NewRequest("GET", url, nil)
	check(err)

	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	check(err)

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	check(err)

	fmt.Println(string(bodyBytes))
}

func check(err error) {
	if err != nil {
		fmt.Println("Error: ",err)
	}
}
