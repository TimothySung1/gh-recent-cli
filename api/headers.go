package api

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

const (
	gitHubApiVersion    = "2022-11-28"
	accept              = "application/vnd.github+json"
	accessTokenFileName = "access-token.txt"
)

var (
	accessToken string
	httpHeader  string
)

func httpAccept() string {
	return fmt.Sprintf("Accept: %s", accept)
}

func httpXGitHubApiVersion() string {
	return fmt.Sprintf("X-GitHub-Api-Version: %s", gitHubApiVersion)
}

// get access token from auth
func init() {
	path := filepath.FromSlash(fmt.Sprintf("/api/auth/%s", accessTokenFileName))
	token, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Access token file not found, check /api/auth/README.md")
	}
	accessToken = string(token)
}

func httpRequest() {
	client := http.Client{}
	fmt.Println(client)
}
