package service

type TaskSerializer struct {
	ID           uint
	Opened       uint
	Closed       uint
	AuthorName   string
	AssignedName string
	Title        string
	Content      string
}
