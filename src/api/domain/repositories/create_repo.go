package repositories

import (
	"strings"

	"github.com/cmd-ctrl-q/industry-rest-microservices/src/api/utils/errors"
)

type CreateRepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (r *CreateRepoRequest) Validate() errors.APIError {
	r.Name = strings.TrimSpace(r.Name)
	if r.Name == "" {
		return errors.NewBadRequestError("invalid repository name")
	}
	return nil
}

type CreateRepoResponse struct {
	ID    int64  `json:"id"`
	Owner string `json:"owner"`
	Name  string `json:"name"`
}

type CreateReposResponse struct {
	StatusCode int `json:"status"`
	Results    []CreateRepositoriesResult
}

type CreateRepositoriesResult struct {
	Response *CreateRepoResponse `json:"repo"`  // response from github
	Error    errors.APIError     `json:"error"` // error from github
}
