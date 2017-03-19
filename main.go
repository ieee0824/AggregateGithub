package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

const (
	token = ""
	user  = ""
)

func main() {

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	list := []string{}

	for i := 1; i <= 3; i++ {
		options := &github.RepositoryListOptions{}
		options.Page = i

		repos, _, err := client.Repositories.List(ctx, user, options)
		if err != nil {
			log.Fatalln(err)
		}

		for _, repo := range repos {
			list = append(list, *repo.Name)
		}
	}

	result := make([]int, 24)
	for _, v := range list {
		punchcards, _, err := client.Repositories.ListPunchCard(ctx, user, v)
		if err != nil {
			log.Fatalln(err)
		}

		for _, punchcard := range punchcards {
			result[*punchcard.Hour] += *punchcard.Commits
		}

	}
	for k, v := range result {
		fmt.Println(k, v)
	}

	all := 0
	for _, v := range result {
		all += v
	}
	fmt.Println(all)
}
