package storage

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
	ID           uint
	Opened       uint
	Closed       uint
	AuthorName   string
	AssignedName string
	Title        string
	Content      string
}
