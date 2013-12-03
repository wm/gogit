package gogit

import (
	"strings"
	"code.google.com/p/goauth2/oauth"
	"github.com/google/go-github/github"
)

type Repo struct {
	Organization string
	Name string
}

type PullState struct {
	Number int
	CommentCount int
	Status string
	Octocatted bool
}

var t = &oauth.Transport{
	Token: &oauth.Token{AccessToken: "7c5f06367ffea77071c84e32f02a505304248097"},
}

var client = github.NewClient(t.Client())

func (repo *Repo) Open() (result []PullState, err error){
	pulls, _, err := client.PullRequests.List(repo.Organization, repo.Name, nil)
	pullStates := make([]PullState, len(pulls))

	for i, pull := range pulls {
		pull, _, _ := client.PullRequests.Get(repo.Organization, repo.Name, *pull.Number)
		pullState := PullState{*pull.Number, *pull.Comments, status(repo, pull), octocatted(repo, pull)}
		pullStates[i] = pullState
	}

	return pullStates, nil
}

func status(repo *Repo, pull *github.PullRequest) (status string) {
	sha := "09fe86ee37b1ec355c1ae55b50b37682f630cca3"
	statuses, _, _ := client.Repositories.ListStatuses(repo.Organization, repo.Name, sha)
	if len(statuses) > 0 {
		return *statuses[0].State
	}

	return ""
}

func octocatted(repo *Repo, pull *github.PullRequest) (bool) {
	comments, _, _ := client.Issues.ListComments(repo.Organization, repo.Name, *pull.Number, nil)
	for _, comment := range comments {
		if hasOctocat(comment) {
			return true
		}
	}
	return false
}

func hasOctocat(comment github.IssueComment)(bool) {
	return strings.ContainsAny(*comment.Body, ":octocat:")
}
