package gogit

import (
	"github.com/google/go-github/github"
)

type Repo struct {
	Organization string
	Name         string
}

func (repo *Repo) OpenPulls() (result []Pull, err error){
	ghPulls, _, err := client.PullRequests.List(repo.Organization, repo.Name, nil)
	if err != nil { return nil, err }

	return updatePulls(repo, &ghPulls), nil
}

func updatePulls(repo *Repo, ghPulls *[]github.PullRequest)([]Pull) {
	pulls := make([]Pull, len(*ghPulls))

	for i, ghPull := range *ghPulls {
		pulls[i] = updatePull(repo, ghPull)
	}

	return pulls
}

func updatePull(repo *Repo, ghPull github.PullRequest)(Pull) {
	pull := Pull{&ghPull, nil, repo}
	pull.Update()

    return pull
}
