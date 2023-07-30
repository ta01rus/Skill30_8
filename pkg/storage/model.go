package storage

type Users struct {
	ID   uint
	Name string
}
type Tasks struct {
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

type TaskView struct {
	Tasks
	AuthorName   string
	AssignedName string
}
