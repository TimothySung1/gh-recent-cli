package cmd

import (
	"fmt"
	"os"

	"github.com/TimothySung1/gh-recent-cli/api"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ghr",
	Short: "gh-recent is a cli tool for retrieving information on recent changes to a GitHub repository",
	Long:  "gh-recent is a cli tool for retrieving information on recent changes to a GitHub repository, including issues, commits, pull requests, and more",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error will executing gh-recent '%s'\n", err)
		os.Exit(1)
	}

	fmt.Println(api.HttpGetRequest("https://api.github.com/issues")) // /repos/{owner}/{repo}/issues for specific repository's issues
}
