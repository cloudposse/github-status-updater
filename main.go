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

var (
	token = flag.String("token", os.Getenv("GITHUB_TOKEN"), "Github auth token")
	owner = flag.String("owner", os.Getenv("GITHUB_OWNER"), "Github repository owner")
	repo  = flag.String("repo", os.Getenv("GITHUB_REPO"), "Github repository name")
	sha   = flag.String("sha", os.Getenv("GITHUB_COMMIT_SHA"), "Github commit status SHA")
	state = flag.String("state", os.Getenv("GITHUB_COMMIT_STATE"), "Github state of commit")
	ctx   = flag.String("ctx", os.Getenv("GITHUB_COMMIT_CONTEXT"), "Github commit status context")
	desc  = flag.String("desc", os.Getenv("GITHUB_COMMIT_DESCRIPTION"), "Github commit status description")
	url   = flag.String("url", os.Getenv("GITHUB_COMMIT_TARGET_URL"), "Github commit status target URL")
)

func main() {
	flag.Parse()

	if *token == "" {
		flag.PrintDefaults()
		log.Fatal("-token required")
	}
	if *owner == "" {
		flag.PrintDefaults()
		log.Fatal("-owner required")
	}
	if *repo == "" {
		flag.PrintDefaults()
		log.Fatal("-repo required")
	}
	if *sha == "" {
		flag.PrintDefaults()
		log.Fatal("-sha required")
	}
	if *state == "" {
		flag.PrintDefaults()
		log.Fatal("-state required")
	}

	http.DefaultClient.Transport = myRoundTripper{*token}
	client := github.NewClient(http.DefaultClient)
	st := &github.RepoStatus{}
	st.State = state

	if *ctx != "" {
		st.Context = ctx
	}
	if *desc != "" {
		st.Description = desc
	}
	if *url != "" {
		st.TargetURL = url
	}

	st, _, err := client.Repositories.CreateStatus(context.Background(), *owner, *repo, *sha, st)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Created status", *st.ID)
}
