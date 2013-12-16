package gogit

import (
	"github.com/google/go-github/github"
)

type Repo struct {
	Organization string
	Name         string
	OpenIssuesCount int
}

func List(owner string) (result []Repo, err error){
	opts := github.RepositoryListByOrgOptions{"private", github.ListOptions{0, 100}}
	ghRepos, _, err := client.Repositories.ListByOrg(owner, &opts)
	if err != nil { return nil, err }

	repos := make([]Repo, len(ghRepos))

	for i, ghRepo := range ghRepos {
		repos[i].Organization = owner
		repos[i].Name         = *ghRepo.Name
		repos[i].OpenIssuesCount = *ghRepo.OpenIssuesCount
	}

	return repos, nil
}

func (repo *Repo) OpenPulls() (result []Pull, err error){
	ghPulls, _, err := client.PullRequests.List(repo.Organization, repo.Name, nil)
	if err != nil { return nil, err }

	return updatePulls(repo, &ghPulls), nil
}

func updatePulls(repo *Repo, ghPulls *[]github.PullRequest)([]Pull) {
	pulls := make([]Pull, len(*ghPulls))
	c := make(chan Pull, len(pulls))

	for _, ghPull := range *ghPulls {
		go updatePull(repo, ghPull, c)
	}

	for i := 0; i < cap(c); i++ {
		pulls[i] = <-c
	}

	return pulls
}

func updatePull(repo *Repo, ghPull github.PullRequest, c chan Pull) {
	pull := Pull{&ghPull, nil, repo}
	pull.Update()

	c <- pull
}
