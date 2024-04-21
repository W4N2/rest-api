package store

import (
	"database/sql"

	"github.com/w4n2/rest-api/src/internal/types"
)

type Store interface {
	// Users
	CreateUser(u *types.User) (*types.User, error)
	GetUserByID(id string) (*types.User, error)
	// Tasks
	CreateTask(t *types.Task) (*types.Task, error)
	GetTask(id string) (*types.Task, error)
	// Projects
	CreateProject(p *types.Project) (*types.Project, error)
	GetProject(id string) (*types.Project, error)
}

type Storage struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (s *Storage) CreateTask(t *types.Task) (*types.Task, error) {
	rows, err := s.db.Exec("INSERT INTO tasks (name, status, projectId, assignedToID) VALUES (?, ?, ?, ?)", t.Name, t.Status, t.ProjectID, t.AssignedTo)

	if err != nil {
		return nil, err
	}

	id, err := rows.LastInsertId()
	if err != nil {
		return nil, err
	}

	t.ID = id
	return t, nil
}

func (s *Storage) GetTask(id string) (*types.Task, error) {
	var t types.Task
	err := s.db.QueryRow("SELECT id, name, status, projectId, assignedToID, createdAt FROM tasks WHERE id = ?", id).Scan(&t.ID, &t.Name, &t.Status, &t.ProjectID, &t.AssignedTo, &t.CreatedAt)
	return &t, err
}

func (s *Storage) CreateProject(p *types.Project) (*types.Project, error) {
	rows, err := s.db.Exec("INSERT INTO projects (name, description) VALUES (?, ?)", p.Name, p.Description)

	if err != nil {
		return nil, err
	}

	id, err := rows.LastInsertId()
	if err != nil {
		return nil, err
	}

	p.ID = id
	return p, nil
}

func (s *Storage) GetProject(id string) (*types.Project, error) {
	var p types.Project
	err := s.db.QueryRow("SELECT id, name, description, createdAt FROM projects WHERE id = ?", id).Scan(&p.ID, &p.Name, &p.Description, &p.CreatedAt)
	return &p, err
}

func (s *Storage) CreateUser(u *types.User) (*types.User, error) {
	rows, err := s.db.Exec("INSERT INTO users (email, firstName, lastName, password) VALUES (?,?,?,?)", u.Email, u.FirstName, u.LastName, u.Password)

	if err != nil {
		return nil, err
	}

	id, err := rows.LastInsertId()
	if err != nil {
		return nil, err
	}

	u.ID = id
	return u, nil
}

func (s *Storage) GetUserByID(id string) (*types.User, error) {
	var u types.User
	err := s.db.QueryRow("SELECT id, email, firstName, lastName, password, createdAt FROM users WHERE id = ?", id).Scan(&u.ID, &u.Email, &u.FirstName, &u.LastName, &u.Password, &u.CreatedAt)
	return &u, err
}
