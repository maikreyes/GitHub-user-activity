package ports

import "github-user-activity/internal/domain"

type GithubUserProvider interface {
	FetchGithubActivity(username string) ([]domain.ApiResponse, error)
}
