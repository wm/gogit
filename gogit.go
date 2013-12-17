package gogit

import (
	"code.google.com/p/goauth2/oauth"
	"github.com/google/go-github/github"
)

var t *oauth.Transport
var client *github.Client

func SetGithubToken(token string) {
	t = &oauth.Transport{ Token: &oauth.Token{AccessToken: token} }
	client = github.NewClient(t.Client())
}

