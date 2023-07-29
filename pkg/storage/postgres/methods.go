package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/pressly/goose"
	"github.com/ta01rus/Skill30_8/pkg/storage"
)

func (db *Postgres) Tasks(ctx context.Context, id, athID, asgID int, offset, limit int) ([]*storage.TaskView, error) {
	ret := []*storage.TaskView{}
	sqlt := `
			select id, title , author_name, assigned_name, "content", opened, closed 
			from task__view
			where ($1 = 0 or id = $1) and
				  ($2 = 0 or assigned_id = $2) and
				  ($3 = 0 or author_id = $3) 
			order by id
			offset $4 limit $5
		 `
	rows, err := db.QueryContext(ctx, sqlt, id, asgID, athID, offset, limit)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		task := new(storage.TaskView)

		err := rows.Scan(task.ID, task.Title, task.AuthorName,
			task.AssignedName, task.Content, task.Opened, task.Closed)
		if err != nil {
			return nil, err
		}
		ret = append(ret, task)
	}

	return ret, nil
}

func (db *Postgres) AddTasks(ctx context.Context, t *storage.Tasks) (*storage.Tasks, error) {
	return nil, nil
}

func (db *Postgres) DelTasks(ctx context.Context, id int) error {
	return nil
}
func (db *Postgres) Users(ctx context.Context, id, offset, limit int) ([]*storage.Users, error) {

	ret := []*storage.Users{}
	sqlt := `
			select id, name	from tasks
			where ($1 = 0 or id = $1) 
			order by id
			offset $2 limit $3
		 `
	rows, err := db.QueryContext(ctx, sqlt, id, offset, limit)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		user := new(storage.Users)

		err := rows.Scan(user.ID, user.Name)
		if err != nil {
			return nil, err
		}
		ret = append(ret, user)
	}
	return ret, nil

}

// добаавить пользователя
func (db *Postgres) AddUsers(ctx context.Context, u *storage.Users) (*storage.Users, error) {
	if !u.Valid() {
		err := fmt.Errorf("user is not valid")
		return nil, err
	}

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	sqlt := ` INSERT INTO users (name) VALUES ($1) RETURNING id	`
	stmt, err := tx.Prepare(sqlt)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, u.Name).Scan(&u.ID)
	if err != nil {
		return nil, err
	}
	tx.Commit()

	return u, nil
}

// удалить пользователя
func (db *Postgres) DelUsers(ctx context.Context, id int) error {
	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
	}
	stmt, err := tx.Prepare("delete from users where id = $1")

	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (db *Postgres) MigrateDB(cmd string, args ...string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()

	err := db.PingContext(ctx)
	if err != nil {
		return err
	}

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	err = goose.Run(cmd, db.Conn(), db.config.DirMigrateTemp, args...)
	if err != nil {
		log.Panicln(err)
	}

	return nil
}

func (db *Postgres) Conn() *sql.DB {
	return db.DB
}
