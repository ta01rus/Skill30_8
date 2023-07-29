package storage

import "context"

type DB interface {

	// получение заданий
	Tasks(context.Context, int, int, int) ([]*Tasks, error)

	AddTasks(context.Context, *Tasks) (*Tasks, error)

	DelTasks(context.Context, int) error

	// пользователи
	Users(context.Context, int) ([]*Users, error)

	// добаавить пользователя
	AddUsers(context.Context, *Users) (*Users, error)

	// удалить пользователя
	DelUsers(context.Context, int) error

	// контроль миграций
	MigrateDB(string, ...string) error
}
