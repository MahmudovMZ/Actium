package models

import "time"

type Task struct {
	Id          int       `json:"-"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Creator_Id  int       `json:"creator_id"`
	CreatedAt   time.Time `json:"created_at"`
	Deadline    string    `json:"deadline"`
}

func (t *Task) AddTask(title, description, status, deadline string, creator_Id int) error {

	t.Title = title
	t.Description = description
	t.Status = status
	t.Deadline = deadline
	t.Creator_Id = creator_Id

	return nil
}

var ValidStatuses = map[string]bool{
	"New":         true,
	"In progress": true,
	"Completed":   true,
	"Canceled":    true,
}
