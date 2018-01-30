package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"runtime"
	"testing"
)

var opts = &options{
	Commit:      "1234",
	Context:     "ci",
	Description: "it passed",
	Repo:        "repo1",
	State:       "success",
	TargetUrl:   "http://example.com",
	User:        "user1",
}

var addHeaderTCs = []struct {
	header string
	value  []string
}{
	{header: "Authorization", value: []string{"token 1234"}},
	{header: "Content-Type", value: []string{"application/json"}},
	{header: "User-Agent", value: []string{"github-commit-status/" + version + " (" + runtime.GOOS + ")"}},
}

func TestAddHeaders(t *testing.T) {
	c := &config{apiToken: "1234"}
	req, _ := http.NewRequest("GET", "http://example.com", nil)
	addHeaders(c, req)
	for _, tc := range addHeaderTCs {
		result := req.Header[tc.header]
		if !reflect.DeepEqual(result, tc.value) {
			t.Errorf("Error: expected '%s' to be\n'%s' got\n'%s'", tc.header, tc.value, result)
		}
	}
}

var apiUrlTcs = []struct {
	envVarValue string
	expected    string
}{
	{"", "https://api.github.com"},
	{"https://github.example.com", "https://github.example.com"},
	{"https://github.example.com/", "https://github.example.com"},
}

func TestGetApiUrl(t *testing.T) {
	for _, tc := range apiUrlTcs {
		err := os.Setenv("GITHUB_API", tc.envVarValue)
		if err != nil {
			t.Errorf("Unable to set env var 'GITHUB_API'")
		}

		result := getApiUrl()
		if result != tc.expected {
			t.Errorf("Error: expected\n'%s' got\n'%s'", tc.expected, result)
		}
		os.Setenv("GITHUB_API", "")
	}
}

func TestStatusRequestBody(t *testing.T) {
	expected := `{"context":"ci","description":"it passed","state":"success","target_url":"http://example.com"}`

	result, err := statusRequestBody(opts)
	if err != nil {
		t.Errorf("Unable to generate status request body")
	}

	if string(result) != expected {
		t.Errorf("Error: expected\n'%s' got\n'%s'", expected, result)
	}

}

func TestStatusUrl(t *testing.T) {
	os.Setenv("GITHUB_API", "http://api.example.com")
	os.Setenv("GITHUB_TOKEN", "foobar")
	c := newConfig()
	expected := os.Getenv("GITHUB_API") + "/repos/user1/repo1/statuses/1234"
	result := statusUrl(opts, c)

	if result != expected {
		t.Errorf("Error: expected\n'%s' got \n'%s", expected, result)
	}
	os.Setenv("GITHUB_API", "")
	os.Setenv("GITHUB_TOKEN", "")
}

var verificationResponseTCs = []struct {
	responseCode int
	expectedErr  bool
}{
	{http.StatusInternalServerError, true},
	{http.StatusBadRequest, true},
	{http.StatusCreated, false},
}

func TestVerifyResponse(t *testing.T) {
	for _, tc := range verificationResponseTCs {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "something failed", tc.responseCode)
		}))

		resp, err := http.Get(ts.URL)
		if err != nil {
			t.Errorf("Error create test request")
		}

		result := verifyResponse(resp)

		if tc.expectedErr {
			if result == nil {
				t.Errorf("Error: expected a http '%d' status code to return an error got '%v'", tc.responseCode, err)
			}
		} else {
			if result != nil {
				t.Errorf("Error: expected a http '%d' status code to not return an error got '%v'", tc.responseCode, err)
			}
		}

		ts.Close()
	}
}
