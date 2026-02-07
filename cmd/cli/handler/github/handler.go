package github

import (
	"github-user-activity/internal/service/github"
)

type Handler struct {
	Service *github.Service
}

func NewHandler(service *github.Service) *Handler {
	return &Handler{
		Service: service,
	}
}
