package projects

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/w4n2/rest-api/src/internal/types"
)

func TestProjectsService_RegisterRoutes(t *testing.T) {
	actual := &ProjectsService{}
	assert.NotNil(t, actual)
}

func TestProjectsService_handleCreateProject(t *testing.T) {
	t.Skip("TODO")
}

func TestProjectsService_handleGetProject(t *testing.T) {
	t.Skip("TODO")
}

func TestNewProjectsService(t *testing.T) {
	t.Skip("TODO")
}

func Test_validateProjectPayload(t *testing.T) {
	actual := validateProjectPayload(&types.Project{})
	assert.Equal(t, errProjectNameRequired, actual)

	actual = validateProjectPayload(&types.Project{Name: "test"})
	assert.Nil(t, actual)

}
