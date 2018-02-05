package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/net/context"
)

type myRoundTripper struct {
	accessToken string
}

func (rt myRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Set("Authorization", fmt.Sprintf("token %s", rt.accessToken))
	return http.DefaultTransport.RoundTrip(r)
}

func isValidState(state string) bool {
	validStates := [4]string{"error", "failure", "pending", "success"}
	for _, s := range validStates {
		if state == s {
			return true
		}
	}
	return false
}

var (
	token       = flag.String("token", os.Getenv("GITHUB_TOKEN"), "Github auth token")
	owner       = flag.String("owner", os.Getenv("GITHUB_OWNER"), "Github repository owner")
	repo        = flag.String("repo", os.Getenv("GITHUB_REPO"), "Github repository name")
	sha         = flag.String("sha", os.Getenv("GITHUB_COMMIT_SHA"), "Github commit SHA")
	state       = flag.String("state", os.Getenv("GITHUB_COMMIT_STATE"), "Github current state of the repository (pending, success, error, or failure)")
	ctx         = flag.String("context", os.Getenv("GITHUB_COMMIT_CONTEXT"), "Github commit status context")
	description = flag.String("description", os.Getenv("GITHUB_COMMIT_DESCRIPTION"), "Github commit status description")
	url         = flag.String("url", os.Getenv("GITHUB_COMMIT_TARGET_URL"), "Github commit status target URL")
)

func main() {
	flag.Parse()

	if *token == "" {
		flag.PrintDefaults()
		log.Fatal("-token or GITHUB_TOKEN required")
	}
	if *owner == "" {
		flag.PrintDefaults()
		log.Fatal("-owner or GITHUB_OWNER required")
	}
	if *repo == "" {
		flag.PrintDefaults()
		log.Fatal("-repo or GITHUB_REPO required")
	}
	if *sha == "" {
		flag.PrintDefaults()
		log.Fatal("-sha or GITHUB_COMMIT_SHA required")
	}
	if *state == "" {
		flag.PrintDefaults()
		log.Fatal("-state or GITHUB_COMMIT_STATE required")
	}
	if !isValidState(*state) {
		flag.PrintDefaults()
		log.Fatal("-state or GITHUB_COMMIT_STATE must be one of 'error', 'failure', 'pending', 'success'")
	}

	repoStatus := &github.RepoStatus{}
	repoStatus.State = state

	if *ctx != "" {
		repoStatus.Context = ctx
	}
	if *description != "" {
		repoStatus.Description = description
	}
	if *url != "" {
		repoStatus.TargetURL = url
	}

	http.DefaultClient.Transport = myRoundTripper{*token}
	githubClient := github.NewClient(http.DefaultClient)
	repoStatus, _, err := githubClient.Repositories.CreateStatus(context.Background(), *owner, *repo, *sha, repoStatus)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Updated status", *repoStatus.ID)
}
