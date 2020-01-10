package repositories

// CreateRepoRequest describes repo to be created
type CreateRepoRequest struct {
	Name string `json:"name"`
}

// CreateRepoResponse describes created repo
type CreateRepoResponse struct {
	ID    int64  `json:"id"`
	Owner string `json:"owner"`
	Name  string `json:"string"`
}
