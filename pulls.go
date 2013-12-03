package gogit

import (
	"strings"
	"github.com/google/go-github/github"
)

type Repo struct {
	Organization string
	Name         string
}

type PullState struct {
	Number       int
	CommentCount int
	Status       string
	Octocatted   bool
}

type Pull struct {
	Data  *github.PullRequest
	State *PullState
	Repo  *Repo
}

func (repo *Repo) OpenPulls() (result []Pull, err error){
	ghPulls, _, err := client.PullRequests.List(repo.Organization, repo.Name, nil)
	pulls := make([]Pull, len(ghPulls))

	for i, githubPull := range ghPulls {
		pull := Pull{&githubPull, nil, repo}
		pull.Update()

		pulls[i] = pull
	}

	return pulls, nil
}

func (pull *Pull) Update() {
	pull.Data, _, _ = client.PullRequests.Get(
		pull.Repo.Organization, pull.Repo.Name, *pull.Data.Number)

	pull.State  = &PullState{
		*pull.Data.Number,
		*pull.Data.Comments,
		status(pull.Repo, pull.Data),
		octocatted(pull.Repo, pull.Data),
	}
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
