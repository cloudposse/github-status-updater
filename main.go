package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/google/go-github/github"
	"golang.org/x/net/context"
)

type roundTripper struct {
	accessToken string
	insecure    bool
}

func (rt roundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Set("Authorization", fmt.Sprintf("token %s", rt.accessToken))
	transport := http.Transport {TLSClientConfig: &tls.Config{InsecureSkipVerify: rt.insecure}}
	return transport.RoundTrip(r)
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
	action      = flag.String("action", os.Getenv("GITHUB_ACTION"), "Action to perform: 'update_state' or 'update_branch_protection'")
	token       = flag.String("token", os.Getenv("GITHUB_TOKEN"), "Github access token")
	owner       = flag.String("owner", os.Getenv("GITHUB_OWNER"), "Github repository owner")
	repo        = flag.String("repo", os.Getenv("GITHUB_REPO"), "Github repository name")
	ref         = flag.String("ref", os.Getenv("GITHUB_REF"), "Commit SHA, branch name or tag")
	state       = flag.String("state", os.Getenv("GITHUB_STATE"), "Commit state. Possible values are 'pending', 'success', 'error' or 'failure'")
	ctx         = flag.String("context", os.Getenv("GITHUB_CONTEXT"), "Status label. Could be the name of a CI environment")
	description = flag.String("description", os.Getenv("GITHUB_DESCRIPTION"), "Short high level summary of the status")
	url         = flag.String("url", os.Getenv("GITHUB_TARGET_URL"), "URL of the page representing the status")
	baseURL     = flag.String("baseURL", os.Getenv("GITHUB_BASE_URL"), "Base URL of github enterprise")
	uploadURL   = flag.String("uploadURL", os.Getenv("GITHUB_UPLOAD_URL"), "Upload URL of github enterprise")
	insecure    = flag.Bool("insecure", strings.ToLower(os.Getenv("GITHUB_INSECURE")) == "true", "Ignore SSL certificate check")
)

func getUserLogins(users []*github.User) []string {
	res := []string{}
	if users != nil && len(users) > 0 {
		for _, user := range users {
			if user != nil {
				res = append(res, user.GetLogin())
			}
		}
	}

	return res
}

func getTeamSlugs(teams []*github.Team) []string {
	res := []string{}
	if teams != nil && len(teams) > 0 {
		for _, team := range teams {
			if team != nil {
				res = append(res, team.GetSlug())
			}
		}
	}

	return res
}

func main() {
	flag.Parse()

	if *action == "" {
		flag.PrintDefaults()
		log.Fatal("-action or GITHUB_ACTION required")
	}
	if *action != "update_state" && *action != "update_branch_protection" {
		flag.PrintDefaults()
		log.Fatal("-action or GITHUB_ACTION must be 'update_state' or 'update_branch_protection'")
	}
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

	http.DefaultClient.Transport = roundTripper{*token, *insecure}
	var githubClient *github.Client
	if *baseURL != "" || *uploadURL != "" {
		if *baseURL == "" {
			flag.PrintDefaults()
			log.Fatal("-baseURL or GITHUB_BASE_URL required when using -uploadURL or GITHUB_UPLOAD_URL")
		}
		if *uploadURL == "" {
			flag.PrintDefaults()
			log.Fatal("-uploadURL or GITHUB_UPLOAD_URL required when using -baseURL or GITHUB_BASE_URL")
		}
		githubClient, _ = github.NewEnterpriseClient(*baseURL, *uploadURL, http.DefaultClient)
	} else {
		githubClient = github.NewClient(http.DefaultClient)
	}

	// Update status of a commit
	if *action == "update_state" {
		if *ref == "" {
			flag.PrintDefaults()
			log.Fatal("-ref or GITHUB_REF is required and must be a commit SHA")
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

		repoStatus, _, err := githubClient.Repositories.CreateStatus(context.Background(), *owner, *repo, *ref, repoStatus)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("github-status-updater: Updated status", *repoStatus.ID)

	} else if *action == "update_branch_protection" {
		if *ref == "" {
			flag.PrintDefaults()
			log.Fatal("-ref or GITHUB_REF is required and must be a branch name")
		}
		if *ctx == "" {
			flag.PrintDefaults()
			log.Fatal("-context or GITHUB_CONTEXT required")
		}

		// https://godoc.org/github.com/google/go-github/github#RepositoriesService.GetBranchProtection
		// Get the existing branch protection and copy all the fields (to not override them), add the provided context to the existing contexts
		protection, _, err := githubClient.Repositories.GetBranchProtection(context.Background(), *owner, *repo, *ref)
		if err != nil {
			log.Fatal(err)
		}

		protectionRequest := &github.ProtectionRequest{}

		var requiredStatusChecks *github.RequiredStatusChecks
		var requiredPullRequestReviews *github.PullRequestReviewsEnforcement
		var enforceAdmins *github.AdminEnforcement
		var restrictions *github.BranchRestrictions

		if protection != nil {
			requiredStatusChecks = protection.GetRequiredStatusChecks()
			requiredPullRequestReviews = protection.GetRequiredPullRequestReviews()
			enforceAdmins = protection.GetEnforceAdmins()
			restrictions = protection.GetRestrictions()
		}

		if requiredStatusChecks == nil {
			requiredStatusChecks = &github.RequiredStatusChecks{Strict: true, Contexts: []string{}}
		}

		protectionRequest.RequiredStatusChecks = &github.RequiredStatusChecks{Strict: requiredStatusChecks.Strict, Contexts: append(requiredStatusChecks.Contexts, *ctx)}

		if requiredPullRequestReviews != nil {
			pullRequestReviewsEnforcementRequest := &github.PullRequestReviewsEnforcementRequest{}

			users := getUserLogins(requiredPullRequestReviews.DismissalRestrictions.Users)
			teams := getTeamSlugs(requiredPullRequestReviews.DismissalRestrictions.Teams)

			dismissalRestrictionsRequest := &github.DismissalRestrictionsRequest{
				Users: &users,
				Teams: &teams,
			}

			pullRequestReviewsEnforcementRequest.DismissalRestrictionsRequest = dismissalRestrictionsRequest
			pullRequestReviewsEnforcementRequest.DismissStaleReviews = requiredPullRequestReviews.DismissStaleReviews
			pullRequestReviewsEnforcementRequest.RequireCodeOwnerReviews = requiredPullRequestReviews.RequireCodeOwnerReviews
			pullRequestReviewsEnforcementRequest.RequiredApprovingReviewCount = requiredPullRequestReviews.RequiredApprovingReviewCount
			protectionRequest.RequiredPullRequestReviews = pullRequestReviewsEnforcementRequest
		}

		if enforceAdmins != nil {
			protectionRequest.EnforceAdmins = enforceAdmins.Enabled
		}

		if restrictions != nil {
			branchRestrictionsRequest := &github.BranchRestrictionsRequest{Users: getUserLogins(restrictions.Users), Teams: getTeamSlugs(restrictions.Teams)}
			protectionRequest.Restrictions = branchRestrictionsRequest
		}

		// https://godoc.org/github.com/google/go-github/github#RepositoriesService.UpdateBranchProtection
		_, _, err = githubClient.Repositories.UpdateBranchProtection(context.Background(), *owner, *repo, *ref, protectionRequest)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("github-status-updater: Updated branch protection")
	}
}
