package main

import (
	"context"
	"fmt"
	"os"

	"github.com/drone/drone-go/drone"
	"github.com/pubgo/xerror"
	"golang.org/x/oauth2"
)

var (
	token = os.Getenv("DRONE_TOKEN")
	host  = os.Getenv("DRONE_SERVER")
)

func main() {
	// create an http client with oauth authentication.
	config := new(oauth2.Config)

	// create the drone client with authenticator
	client := drone.NewClient(host, config.Client(
		context.Background(),
		&oauth2.Token{
			AccessToken: token,
		},
	))

	// gets the current user
	user, err := client.Self()
	xerror.Panic(err)
	fmt.Printf("%#v\n\n", user)

	// gets the named repository information
	repo, err := client.Repo(os.Getenv("Repo_namespace"), os.Getenv("Repo_name"))
	xerror.Panic(err)
	fmt.Printf("%#v\n\n", repo)

	lst, err := client.BuildList(os.Getenv("Repo_namespace"), os.Getenv("Repo_name"), drone.ListOptions{})
	xerror.Panic(err)
	for i := range lst {
		fmt.Printf("%#v\n\n", lst[i])
	}

}
