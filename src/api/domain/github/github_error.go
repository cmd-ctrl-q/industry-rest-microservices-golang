package github

/*
the reason were adding the status code to the response body is because the
only place where we have the status code is after calling the api.
without it, we wouldn't know if this was an internal server error, bad request, etc.
currently it is: 422 Unprocessable Entity
*/

type GithubErrorResponse struct {
	StatusCode       int           `json:"status_code"`
	Message          string        `json:"message"`
	Errors           []GithubError `json:"errors"`
	DocumentationURL string        `json:"documentation_url"`
}

type GithubError struct {
	Resource string `json:"resource"`
	Code     string `json:"code"`
	Field    string `json:"field"`
	Message  string `json:"message"`
}

/*
"errors": [
        {
            "resource": "Repository",
            "code": "custom",
            "field": "name",
            "message": "name already exists on this account"
        }
	],
*/
