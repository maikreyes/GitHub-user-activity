package githubuser

import "github-user-activity/internal/domain"

func (p *Provider) FetchGithubActivity(username string) ([]domain.ApiResponse, error) {
	return p.Client.FetchGithubActivity(username)
}
