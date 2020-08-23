package main

import (
	"github.com/go-git/go-git/v5"
	"log"
	"fmt"
	"flag"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/go-git/go-git/v5/config"
)

var (
	repo = flag.String("repo", "", "GitHub repository")
)

func main() {
	flag.Parse()

	// First, a git repository needs to be initialized if it does not already exist.
	// Storage will be a storage base on memory.
	r, _ := git.Init(memory.NewStorage(), nil)

	// Create remote with name of remote and URL of repository
	_, err := r.CreateRemote(&config.RemoteConfig{
		Name: "website",
		URLs: []string{"https://github.com/" + *repo},
	})
	if err != nil {
		log.Fatal(err)
	}

	// Get a list of all remotes
	list, err := r.Remotes()
	if err != nil {
		log.Fatal(err)
	}

	for _, r := range list {
		fmt.Println(r)
	}
}
