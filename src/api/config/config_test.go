package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstants(t *testing.T) {
	assert.EqualValues(t, "SECRET_GITHUB_ACCESS_TOKEN", secretGithubAccessToken)
}

func TestGetGithubAccessToken(t *testing.T) {
	assert.EqualValues(t, "6ff35b08075e998de98916b50b156dc8e5b010ea", GetGithubAccessToken())
}
