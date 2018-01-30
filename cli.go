package main

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
)

type options struct {
	Commit      string `short:"c" long:"commit" description:"Commit (SHA) to update status on" required:"true"`
	Context     string `long:"context" description:"Context to include in status update"`
	Description string `short:"d" long:"description" description:"Text to include with the status update"`
	Repo        string `short:"r" long:"repo" description:"Github repo name where the commit exists" required:"true"`
	State       string `short:"s" long:"state" description:"State. Must be one of 'pending', 'success', 'error', or 'failure'" required:"true"`
	TargetUrl   string `short:"t" long:"target-url" description:"URL to include in status update"`
	User        string `short:"u" long:"user" description:"Github user that owns the repo" required:"true"`
	Version     func() `short:"v" long:"version" description:"Display the version github-commit-status"`
}

func parseCliArgs() *options {
	opts := &options{}

	opts.Version = func() {
		fmt.Println(version)
		os.Exit(0)
	}

	parser := flags.NewParser(opts, flags.Default)

	args, err := parser.Parse()
	if err != nil {
		helpDisplayed := false

		for _, i := range args {
			if i == "-h" || i == "--help" {
				helpDisplayed = true
				break
			}
		}

		if !helpDisplayed {
			parser.WriteHelp(os.Stderr)
		}
		os.Exit(1)
	}

	if !validState(opts.State) {
		os.Stderr.Write([]byte(fmt.Sprintf("'%s' is an invalid state\n", opts.State)))
		parser.WriteHelp(os.Stderr)
		os.Exit(1)
	}

	return opts
}

func states() []string {
	return []string{"error", "failure", "pending", "success"}
}

func validState(state string) bool {
	for _, s := range states() {
		if state == s {
			return true
		}
	}
	return false
}
