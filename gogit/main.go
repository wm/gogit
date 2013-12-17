// This needs to be in its own folder since it has a different package
package main

import (
	"github.com/wm/gogit"
	"fmt"
	"syscall"
	"flag"
)

var opts map[string]string
var ok bool

func main() {
	opts, ok = readOptions()

	if !ok {
		printUsage()
		return
	}

	gogit.SetGithubToken(opts["TOKEN"])
	repos, _ := gogit.ListRepos(opts["OWNER"])

	c := make(chan []gogit.Pull, len(repos))

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

func getReposOpenPulls(repo gogit.Repo, c chan []gogit.Pull)  {
	pulls, _ := repo.OpenPulls()
	c <- pulls
}

func printRepoPulls(c chan []gogit.Pull) {
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
