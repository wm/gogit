package gogit

import (
	"syscall"
	"fmt"
	"code.google.com/p/goauth2/oauth"
	"github.com/google/go-github/github"
	"flag"
)

var t *oauth.Transport
var client *github.Client

func Run() {
	opts, ok := readOptions()

	if !ok {
		fmt.Printf("Please set your env variable\n")
		fmt.Printf("GOGIT_GH_TOKEN\n")
		return
	}

	t = &oauth.Transport{ Token: &oauth.Token{AccessToken: opts["TOKEN"]} }
	client = github.NewClient(t.Client())

	repos, _ := List(opts["OWNER"])

	for _, repo := range repos {
	    if repo.OpenIssuesCount > 0 {
			pulls, _ := repo.OpenPulls()
			printPulls(repo.Name, pulls)
		}
	}
}

func printPulls(repoName string, pulls []Pull) {
	pullCount := len(pulls)

	fmt.Printf("%v (%v)\n", repoName, pullCount)

	if pullCount > 0 {
		fmt.Printf("| Pull | Comments | Passing | :octocatted: |\n")
		for _, pull := range pulls {
			fmt.Printf("| %4d | %8d | %7s | %12v |\n",
			pull.State.Number,
			pull.State.CommentCount,
			pull.State.Status,
			pull.State.Octocatted)
		}
	}

	fmt.Println("")
}

func readOptions() (opts map[string]string, ok bool) {
	var owner string
	flag.StringVar(&owner, "owner", "wm", "The Owner (Org/user) of the repos")
	flag.Parse()

	accessToken, okToken := syscall.Getenv("GOGIT_GH_TOKEN")

	opts = map[string]string {
		"TOKEN": accessToken,
		"OWNER": owner,
	}
	return opts, okToken
}
