// This needs to be in its own folder since it has a different package
package main

import (
	"github.com/wm/gogit"
	"fmt"
)

func main() {
	repos, _ := gogit.List(gogit.Opts["OWNER"])

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
