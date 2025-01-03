package api

import (
	"fmt"
	"io"
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
	httpHeader http.Header = http.Header{}
)

// get access token from auth
func init() {
	path := filepath.FromSlash(fmt.Sprintf("./api/auth/%s", accessTokenFileName))
	token, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Access token file not found, check /api/auth/README.md: ", err)
		os.Exit(1)
	}
	accessToken := string(token)

	fmt.Println("Access token found:", accessToken)

	httpHeader.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	httpHeader.Set("X-GitHub-Api-Version", gitHubApiVersion)
	httpHeader.Set("Accept", accept)
}

func HttpGetRequest(uri string) string {
	client := http.Client{}
	request, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "client: could not create request: %s\n", err)
		os.Exit(1)
	}
	request.Header = httpHeader.Clone()
	response, err := client.Do(request)
	if err != nil {
		fmt.Fprintf(os.Stderr, "client: error making http request: %s\n", err)
		os.Exit(1)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "client: could not read response body: %s\n", err)
		os.Exit(1)
	}

	return string(body)
}
