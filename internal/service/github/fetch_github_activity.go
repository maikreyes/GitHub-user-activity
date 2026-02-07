package github

import (
	"fmt"
	"slices"
)

func (s *Service) FetchGithubActivity(username string) error {

	ApiResponse, err := s.Provider.FetchGithubActivity(username)

	if err != nil {
		return err
	}

	var push int = 0
	var pullRequest int = 0
	var issues int = 0
	var issueComment int = 0
	var create int = 0
	var delete int = 0
	var fork int = 0
	var watch int = 0

	var repos []string

	for _, event := range ApiResponse {
		switch event.Type {
		case "PushEvent":
			push++
		case "PullRequestEvent":
			pullRequest++
		case "IssuesEvent":
			issues++
		case "IssueCommentEvent":
			issueComment++
		case "CreateEvent":
			create++
		case "DeleteEvent":
			delete++
		case "ForkEvent":
			fork++
		case "WatchEvent":
			watch++
		default:
			continue
		}

		if !slices.Contains(repos, event.Repo.Name) {
			repos = append(repos, event.Repo.Name)
		}

	}

	if push != 0 {
		fmt.Printf("Push Events: %d\n", push)
	}

	if pullRequest != 0 {
		fmt.Printf("Pull Request Events: %d\n", pullRequest)
	}

	if issues != 0 {
		fmt.Printf("Issues Events: %d\n", issues)
	}

	if issueComment != 0 {
		fmt.Printf("Issue Comment Events: %d\n", issueComment)
	}

	if create != 0 {
		fmt.Printf("Create Events: %d\n", create)
	}

	if delete != 0 {
		fmt.Printf("Delete Events: %d\n", delete)
	}

	if fork != 0 {
		fmt.Printf("Fork Events: %d\n", fork)
	}

	if watch != 0 {
		fmt.Printf("Watch Events: %d\n", watch)
	}

	fmt.Printf("Repositories: %v\n", repos)

	return nil

}
