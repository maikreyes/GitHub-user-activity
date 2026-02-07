package github

func (h *Handler) GetActivity(username string) error {

	return h.Service.FetchGithubActivity(username)
}
