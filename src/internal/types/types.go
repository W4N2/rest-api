package main

import "time"

type Task struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Status     string    `json:"status"`
	ProjectID  int64     `json:"projectID"`
	assignedTo int64     `json:"assignedTo"`
	createdAt  time.Time `json:"createdAt"`
}

type Project struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	createdAt time.Time `json:"createdAt"`
}

type User struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Password  string    `json:"password"`
	createdAt time.Time `json:"createdAt"`
}
