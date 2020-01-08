package github

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRepoRequestAsJson(t *testing.T) {
	request := CreateRepoRequest{
		Name:        "Test Name",
		Description: "Test Description",
		Homepage:    "Test Homepage",
		Private:     false,
		HasIssues:   true,
		HasProjects: true,
		HasWiki:     true,
	}

	bytes, err := json.Marshal(request)
	// assert we can go from request to bytearray and back
	var target CreateRepoRequest
	unmarshalErr := json.Unmarshal(bytes, &target)
	assert.Nil(t, unmarshalErr)
	assert.EqualValues(t, target.Name, request.Name)
	assert.Nil(t, err)
	assert.NotNil(t, bytes)
	// when you know all field values, i.e. none are dynamically generated, like time.
	assert.EqualValues(t, `{"name":"Test Name","description":"Test Description","homepage":"Test Homepage","private":false,"has_issues":true,"has_projects":true,"has_wiki":true}`, string(bytes))
}
