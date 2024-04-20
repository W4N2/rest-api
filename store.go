package main

import "database/sql"

type Store interface {
	// Users
	CreateUser() error
	// Tasks
	CreateTask(t *Task) (*Task, error)
	GetTask(id string) (*Task, error)
	// Projects
	CreateProject(p *Project) (*Project, error)
	GetProject(id string) (*Project, error)
}

type Storage struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (s *Storage) CreateUser() error {
	return nil
}

func (s *Storage) CreateTask(t *Task) (*Task, error) {
	rows, err := s.db.Exec("INSERT INTO tasks (name, status, projectId, assignedToID) VALUES (?, ?, ?, ?)", t.Name, t.Status, t.ProjectID, t.assignedTo)

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

func (s *Storage) GetTask(id string) (*Task, error) {
	var t Task
	err := s.db.QueryRow("SELECT id, name, status, projectId, assignedToID, createdAt FROM tasks WHERE id = ?", id).Scan(&t.ID, &t.Name, &t.Status, &t.ProjectID, &t.assignedTo, &t.createdAt)
	return &t, err
}

func (s *Storage) CreateProject(p *Project) (*Project, error) {
	rows, err := s.db.Exec("INSERT INTO projects (name) VALUES (?)", p.Name)

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

func (s *Storage) GetProject(id string) (*Project, error) {
	var p Project
	err := s.db.QueryRow("SELECT id, name, createdAt FROM projects WHERE id = ?", id).Scan(&t.ID, &t.Name, &t.createdAt)
	return &p, err
}
