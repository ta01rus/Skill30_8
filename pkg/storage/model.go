package storage

import (
	"errors"
	"fmt"
)

type Users struct {
	ID   uint   `form:"user_id"`
	Name string `form:"user"`
}

func (u *Users) Valid() bool {
	return u.Name != ""
}

type Tasks struct {
	ID         uint
	Opened     uint
	Closed     uint
	AuthorID   uint
	AssignedID uint
	Title      string `form:"title"`
	Content    string `form:"content"`
}

type Labels struct {
	ID   uint
	Name string
}

type TaskLabels struct {
	ID      uint
	LabelID uint
	TaskID  uint
}

type TaskView struct {
	Tasks
	AuthorName   string `form:"author"`
	AssignedName string `form:"assigned"`
}

func (t *TaskView) Check() error {
	var err error

	if t.AuthorName == "" {
		errors.Join(err, fmt.Errorf("AuthorName is empty"))
	}

	if t.AssignedName == "" {
		errors.Join(err, fmt.Errorf("AssignedName is empty"))
	}

	if t.Title == "" {
		errors.Join(err, fmt.Errorf("Title is empty"))
	}
	if t.Content == "" {
		errors.Join(err, fmt.Errorf("Content is empty"))
	}
	return err
}
