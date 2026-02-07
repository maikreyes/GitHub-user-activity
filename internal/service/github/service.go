package github

import "github-user-activity/internal/ports"

type Service struct {
	Provider ports.GithubUserProvider
}

func NewService(provider ports.GithubUserProvider) *Service {
	return &Service{
		Provider: provider,
	}
}
