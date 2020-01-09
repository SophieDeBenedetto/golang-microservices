package github

// CreateRepoRequest describes the repo
type CreateRepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
	Private     bool   `json:"private"`
	HasIssues   bool   `json:"has_issues"`
	HasProjects bool   `json:"has_projects"`
	HasWiki     bool   `json:"has_wiki"`
}

// CreateRepoResponse from Github API
type CreateRepoResponse struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	FullName string    `json:"full_name"`
	Owner    RepoOwner `json:"owner"`
	// OR, if you don't need to ever work with RepoOwner on its own
	// Owner struct {

	// }
}

// RepoOwner describes the repo owner
type RepoOwner struct {
	ID      int64  `json:"id"`
	Login   string `json:"login"`
	URL     string `json:"url"`
	HTMLURL string `json:"html_url"`
}
