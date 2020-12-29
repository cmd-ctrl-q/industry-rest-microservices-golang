package github

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRepoRequestAsJson(t *testing.T) {
	request := CreateRepoRequest{
		Name:        "golang-tut",
		Description: "golang rest api microservices tut",
		Homepage:    "https://github.com",
		Private:     true,
		HasIssues:   true,
		HasProjects: true,
		HasWiki:     true,
	}

	// Validation Step. (results are expected results)

	// Marshal takes an input interface and attempts to create a valid json string
	bytes, err := json.Marshal(request)
	assert.Nil(t, err)      // not expecting error when marshalling.
	assert.NotNil(t, bytes) //
	fmt.Println(string(bytes))

	// Unmarshal into struct
	var target CreateRepoRequest
	err = json.Unmarshal(bytes, &target)
	assert.Nil(t, err) // make sure you have Nil / No error after unmarshalling bytes into target.
	fmt.Println(target)
	// assert.EqualValues(t, `{"name":"golang-tut","description":"golang rest api microservices tut","homepage":"https://github.com","private":true,"has_issues":true,"has_projects":true,"has_wiki":true}`, string(bytes))

	// Test indivudal attributes
	assert.EqualValues(t, target.Name, request.Name)
	assert.EqualValues(t, target.Description, request.Description)
	assert.EqualValues(t, target.Homepage, request.Homepage)
	assert.EqualValues(t, target.Private, request.Private)
	assert.EqualValues(t, target.HasIssues, request.HasIssues)
	assert.EqualValues(t, target.HasProjects, request.HasProjects)
	assert.EqualValues(t, target.HasWiki, request.HasWiki)
}
