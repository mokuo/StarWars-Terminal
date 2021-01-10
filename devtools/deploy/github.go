package main

import (
	"context"
	"log"
	"os"

	"github.com/google/go-github/v33/github"

	"golang.org/x/oauth2"
)

func repositoriesService(ctx context.Context) *github.RepositoriesService {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	return client.Repositories
}

func GetLatestRelease() *github.RepositoryRelease {
	ctx := context.Background()
	repo := repositoriesService(ctx)
	release, res, err := repo.GetLatestRelease(ctx, "mokuo", "starwars-terminal")
	if err != nil {
		log.Fatal(res.Body)
		log.Fatal(err)
	}

	return release
}

func CreateRelease(version, body string) *github.RepositoryRelease {
	ctx := context.Background()
	repo := repositoriesService(ctx)
	release := github.RepositoryRelease{
		TagName: &version,
		Name:    &version,
		Body:    &body,
	}
	createdRelease, res, err := repo.CreateRelease(ctx, "mokuo", "starwars-terminal", &release)
	if err != nil {
		log.Fatal(res.Body)
		log.Fatal(err)
	}

	return createdRelease
}
