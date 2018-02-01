package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"strings"
)

func addHeaders(c *config, req *http.Request) {
	req.Header.Add("Authorization", "token "+c.apiToken)
	req.Header.Add("Content-Type", "application/json")
	agent := "github-commit-status/" + version + " (" + runtime.GOOS + ")"
	req.Header.Add("User-Agent", agent)
}

type config struct {
	apiUrl   string
	apiToken string
}

func newConfig() *config {
	c := &config{
		apiToken: os.Getenv("GITHUB_TOKEN"),
		apiUrl:   getApiUrl(),
	}

	if c.apiUrl == "" {
		fmt.Printf("Error: Invalid API URL specified '%s'", c.apiUrl)
		os.Exit(1)
	}

	if c.apiToken == "" {
		fmt.Printf("Error: GITHUB_TOKEN environment variable not specified")
		os.Exit(1)
	}

	return c
}

func getApiUrl() string {
	if apiUrl := os.Getenv("GITHUB_API"); apiUrl != "" {
		return strings.TrimSuffix(apiUrl, "/")
	}
	return "https://api.github.com"
}

type statusBody struct {
	Context     string `json:"context,omitempty"`
	Description string `json:"description,omitempty"`
	State       string `json:"state"`
	TargetUrl   string `json:"target_url,omitempty"`
}

func statusRequestBody(o *options) ([]byte, error) {
	body := &statusBody{
		Context:     o.Context,
		Description: o.Description,
		State:       o.State,
		TargetUrl:   o.TargetUrl,
	}

	b, err := json.Marshal(body)
	if err != nil {
		return []byte{}, err
	}

	return b, nil
}

func statusUrl(o *options, c *config) string {
	return fmt.Sprintf("%s/repos/%s/%s/statuses/%s", c.apiUrl, o.User, o.Repo, o.Commit)
}

func updateStatus(opts *options, conf *config) error {
	url := statusUrl(opts, conf)

	body, err := statusRequestBody(opts)
	if err != nil {
		msg := fmt.Sprintf("Error creating API request body: %s", err.Error())
		return errors.New(msg)
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	addHeaders(conf, req)

	resp, err := client.Do(req)
	if err != nil {
		msg := fmt.Sprintf("Error making API call: %s", err.Error())
		return errors.New(msg)
	}

	return verifyResponse(resp)
}

func verifyResponse(resp *http.Response) error {
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		msg := fmt.Sprintf("Did not receive a successful API response (received '%s')\n", resp.Status)

		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			msg = msg + fmt.Sprintf("Error reading response body: %s", err)
			return errors.New(msg)
		}
		msg = msg + fmt.Sprintf("Response Body:\n%s", contents)
		return errors.New(msg)
	}

	return nil
}

func main() {
	options := parseCliArgs()
	config := newConfig()

	fmt.Printf("Updating status...\n")

	if err := updateStatus(options, config); err != nil {
		os.Stderr.Write([]byte(err.Error()))
		os.Exit(1)
	}
	fmt.Printf("Status update complete")
}
