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
		printUsage()
		return
	}

	t = &oauth.Transport{ Token: &oauth.Token{AccessToken: opts["TOKEN"]} }
	client = github.NewClient(t.Client())

	repos, _ := List(opts["OWNER"])

	c := make(chan []Pull, len(repos))

	for _, repo := range repos {
	    if repo.OpenIssuesCount > 0 {
			go getReposOpenPulls(repo, c)
		}
	}

	for _, repo := range repos {
	    if repo.OpenIssuesCount > 0 {
			printRepoPulls(c)
		}
	}
}

func getReposOpenPulls(repo Repo, c chan []Pull)  {
	pulls, _ := repo.OpenPulls()
	c <- pulls
}

func printRepoPulls(c chan []Pull) {
	pulls := <-c
	pullCount := len(pulls)

	if pullCount > 0 {
		repoName := pulls[0].Repo.Name
		fmt.Printf("%v (%v)\n", repoName, pullCount)
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
	accessToken, okToken := syscall.Getenv("GOGIT_GH_TOKEN")
	owner, okOwner       := syscall.Getenv("GOGIT_OWNER")

	if !okToken { accessToken = "" }
	if !okOwner { owner = "" }

	ownerFlag := flag.String("owner", owner, "The Owner (Org/user) of the repos")
	tokenFlag := flag.String("token", accessToken, "The github token")

	flag.Parse()

	opts = map[string]string {
		"TOKEN": *tokenFlag,
		"OWNER": *ownerFlag,
	}

	if (opts["TOKEN"] == "") || (opts["OWNER"] == "") {
		return nil, false
	}

	return opts, true
}

func printUsage() {
	fmt.Printf("Usage: gogit -token 'MY_GH_TOKEN' -owner 'MyOrganization'")
	fmt.Printf("\n")
	fmt.Printf("Alternatively you can set the GOGIT_GH_TOKEN and GOGIT_OWNER")
	fmt.Printf(" env variables.\n")
}
