package githubuser

import client "github-user-activity/internal/provider/GithubUser/Client"

type Provider struct {
	Client *client.Client
}

func NewProvider(client *client.Client) *Provider {
	return &Provider{
		Client: client,
	}
}
