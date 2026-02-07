package github

import githubuser "github-user-activity/internal/provider/GithubUser"

type Handler struct {
	Provider *githubuser.Provider
}

func NewHandler(provider *githubuser.Provider) *Handler {
	return &Handler{
		Provider: provider,
	}
}
