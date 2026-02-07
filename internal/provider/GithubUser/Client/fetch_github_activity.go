package client

import (
	"encoding/json"
	"github-user-activity/internal/domain"
)

func (c *Client) FetchGithubActivity(username string) ([]domain.ApiResponse, error) {

	resp, err := c.httpClient.Get(c.ApiUrl + "/users/" + username + "/events")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var apiResponses []domain.ApiResponse
	err = json.NewDecoder(resp.Body).Decode(&apiResponses)
	if err != nil {
		return nil, err
	}

	return apiResponses, nil

}
