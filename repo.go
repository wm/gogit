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
