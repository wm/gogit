package gogit

import (
	"fmt"
	"code.google.com/p/goauth2/oauth"
	"github.com/google/go-github/github"
)

var t = &oauth.Transport{
	Token: &oauth.Token{AccessToken: "7c5f06367ffea77071c84e32f02a505304248097"},
}
var client = github.NewClient(t.Client())

func Run() {
  repo := Repo{"IoraHealth", "IoraHealth"}
  pulls, _ := repo.OpenPulls()

  for _, pull := range pulls {
	  fmt.Printf("[number: %d, comments: %d, status: %s, octocatted: %v]\n",
	             pull.State.Number,
	             pull.State.CommentCount,
	             pull.State.Status,
	             pull.State.Octocatted)
  }
}
