package storage

type Users struct {
	ID   uint
	Name string
}

func (u *Users) Valid() bool {
	return u.Name != ""
}

type Task struct {
	ID         uint
	Opened     uint
	Closed     uint
	AuthorID   uint
	AssignedID uint
	Title      string
	Content    string
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
