package store

import "github.com/w4n2/rest-api/src/internal/types"

// Mocks

type MockStore struct{}

func (m *MockStore) CreateTask(t *types.Task) (*types.Task, error) {
	return &types.Task{}, nil
}

func (m *MockStore) GetTask(id string) (*types.Task, error) {
	return &types.Task{}, nil
}

func (m *MockStore) CreateProject(p *types.Project) (*types.Project, error) {
	return &types.Project{}, nil
}

func (m *MockStore) GetProject(id string) (*types.Project, error) {
	return &types.Project{}, nil
}

func (m *MockStore) GetUserByID(id string) (*types.User, error) {
	return &types.User{}, nil
}

func (m *MockStore) CreateUser(u *types.User) (*types.User, error) {
	return &types.User{}, nil
}
