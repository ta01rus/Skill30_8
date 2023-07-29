package storage

import "context"

type DB interface {

	// получение заданий
	Tasks(ctx context.Context, id, athID, asgID int, offset, limit int) ([]*TaskView, error)

	AddTasks(context.Context, *TaskView) (*TaskView, error)

	DelTasks(context.Context, int) error

	// пользователи
	Users(ctx context.Context, id int, offset, limit int) ([]*Users, error)

	// добаавить пользователя
	AddUsers(context.Context, *Users) (*Users, error)

	// удалить пользователя
	DelUsers(context.Context, int) error

	// контроль миграций
	MigrateDB(string, ...string) error
}
