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
	ref         = flag.String("ref", os.Getenv("GITHUB_REF"), "Ref can be a SHA, a branch name, or a tag name")
	state       = flag.String("state", os.Getenv("GITHUB_STATE"), "State of the commit, branch or tag. Possible values are pending, success, error, or failure")
	ctx         = flag.String("context", os.Getenv("GITHUB_CONTEXT"), "Label to differentiate this status from the statuses of other systems")
	description = flag.String("description", os.Getenv("GITHUB_DESCRIPTION"), "Short high level summary of the status")
	url         = flag.String("url", os.Getenv("GITHUB_TARGET_URL"), "URL of the page representing the status")
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
	if *ref == "" {
		flag.PrintDefaults()
		log.Fatal("-ref or GITHUB_REF required")
	}
	if *state == "" {
		flag.PrintDefaults()
		log.Fatal("-state or GITHUB_STATE required")
	}
	if !isValidState(*state) {
		flag.PrintDefaults()
		log.Fatal("-state or GITHUB_STATE must be one of 'error', 'failure', 'pending', 'success'")
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
	repoStatus, _, err := githubClient.Repositories.CreateStatus(context.Background(), *owner, *repo, *ref, repoStatus)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Updated status", *repoStatus.ID)
}
