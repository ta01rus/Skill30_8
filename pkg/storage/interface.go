package storage

import "context"

type DB interface {
	Tasks(ctx context.Context, id, athID, asgID int, offset, limit int) ([]*Tasks, error)

	InsTasks(context.Context, *Tasks) (*Tasks, error)

	UpdTasks(context.Context, *Tasks) (*Tasks, error)

	DelTasks(context.Context, int) error

	// контроль миграций
	MigrateDB(string, ...string) error
}
