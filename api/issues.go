package api

import "fmt"

// wraps the http get request corresponding with issues.
// gives a concise message for new changes to issues in the given repo

func GetIssues(repo repoInfo) {
	uri := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues", repo.Owner, repo.Name)
	res := HttpGetRequest(uri)
	fmt.Println(res)
}
