package gogit

import (
	"syscall"
	"fmt"
	"code.google.com/p/goauth2/oauth"
	"github.com/google/go-github/github"
)

var t *oauth.Transport
var client *github.Client

func Run() {
	opts, ok := readOptions()

	if !ok {
		fmt.Printf("Please set your env variables\n")
		fmt.Printf("GOGIT_GH_TOKEN, GOGIT_GH_OWNER, GOGIT_GH_REPO\n")
		return
	}

	t = &oauth.Transport{ Token: &oauth.Token{AccessToken: opts["TOKEN"]} }
	client = github.NewClient(t.Client())

	repo := Repo{opts["OWNER"], opts["REPO"]}
	pulls, _ := repo.OpenPulls()

	fmt.Printf("| Pull | Comments | Passing | :octocatted: |\n")
	for _, pull := range pulls {
		fmt.Printf("| %4d | %8d | %7s | %12v |\n",
			pull.State.Number,
			pull.State.CommentCount,
			pull.State.Status,
			pull.State.Octocatted)
	}
}

func readOptions() (opts map[string]string, ok bool) {
	accessToken, okToken := syscall.Getenv("GOGIT_GH_TOKEN")
	owner, okOwner := syscall.Getenv("GOGIT_GH_OWNER")
	repoName, okRepo := syscall.Getenv("GOGIT_GH_REPO")

	opts = map[string]string {
		"TOKEN": accessToken,
		"OWNER": owner,
		"REPO": repoName,
	}
	return opts, (okToken && okOwner && okRepo)
}
