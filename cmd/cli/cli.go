package cli

import (
	"flag"
	handler "github-user-activity/cmd/cli/handler/github"
	githubuser "github-user-activity/internal/provider/GithubUser"
	client "github-user-activity/internal/provider/GithubUser/Client"
	service "github-user-activity/internal/service/github"
)

func Run() {

	//Settings

	Providerclient := client.NewClient("https://api.github.com")
	Provider := githubuser.NewProvider(Providerclient)
	service := service.NewService(Provider)
	Handler := handler.NewHandler(service)

	//Usage

	flag.Parse()

	args := flag.Args()

	if len(args) < 1 {
		println("Please provide a GitHub username as an argument.")
		return
	}

	username := args[0]

	println("Fetching activity for GitHub user:", username)

	err := Handler.GetActivity(username)
	if err != nil {
		panic(err)
	}

}
