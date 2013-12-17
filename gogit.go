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
var ok bool
var Opts map[string]string

func init() {
	Opts, ok = readOptions()

	if !ok {
		printUsage()
		return
	}

	t = &oauth.Transport{ Token: &oauth.Token{AccessToken: Opts["TOKEN"]} }
	client = github.NewClient(t.Client())
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
