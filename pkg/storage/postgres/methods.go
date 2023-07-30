package postgres

import (
	"context"
	"log"
	"time"

	"github.com/pressly/goose"
	"github.com/ta01rus/Skill30_8/pkg/storage"
)

func (db *Postgres) Tasks(ctx context.Context, id, athID, asgID int, offset, limit int) ([]*storage.Tasks, error) {
	ret := []*storage.Tasks{}
	sqlt := `
			select id, title , author_id, assigned_id, "content", opened, closed 
			from tasks
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
		task := new(storage.Tasks)

		err := rows.Scan(&task.ID, &task.Title, &task.AuthorID,
			&task.AssignedID, &task.Content, &task.Opened, &task.Closed)
		if err != nil {
			return nil, err
		}
		ret = append(ret, task)
	}

	return ret, nil
}

// добавление
func (db *Postgres) InsTasks(ctx context.Context, t *storage.Tasks) (*storage.Tasks, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	sqlt := `INSERT INTO tasks (TITLE, AUTHOR_ID, ASSIGNED_ID, CONTENT, opened, closed) 
			 VALUES ($1, $2, $3, $4, 0, 0) 	
			 RETURNING id`

	stmt, err := tx.Prepare(sqlt)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, t.Title, t.AuthorID, t.AssignedID, t.Content).Scan(&t.ID)

	if err != nil {
		return nil, err
	}
	tx.Commit()

	return t, nil
}

func (db *Postgres) UpdTasks(ctx context.Context, t *storage.Tasks) (*storage.Tasks, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	sqlt := `UPDATE TASKS 
				SET  
	 			TITLE = $2,
				AUTHOR_ID = $3,
				ASSIGNED_ID = $4,
				CONTENT = $5,
				OPENED = 0,
				CLOSED = 0 
			 WHERE ID = $1`

	stmt, err := tx.Prepare(sqlt)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, t.ID, t.Title, t.AuthorID, t.AssignedID, t.Content)
	if err != nil {
		return nil, err
	}
	tx.Commit()

	return t, nil
}

// удаление зааадания
func (db *Postgres) DelTasks(ctx context.Context, id int) error {
	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
	}
	stmt, err := tx.Prepare("delete from tasks where id = $1")

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

// управление мигрантором
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

	err = goose.Run(cmd, db.DB, db.config.DirMigrateTemp, args...)
	if err != nil {
		log.Panicln(err)
	}

	return nil
}
