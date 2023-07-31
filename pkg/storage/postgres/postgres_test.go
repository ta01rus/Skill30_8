package postgres

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/ta01rus/Skill30_8/pkg/storage"
)

var (
	Db storage.DB
)

func TestMain(m *testing.M) {
	var err error

	Db, err = New(&DbConfig{
		User:   "foo",
		Passw:  "1234567",
		DbName: "skill_db",
		Host:   "127.0.0.1",
		Port:   "5444",
	})
	if err != nil {
		log.Fatal(err)
	}

	m.Run()
}

func Test_InsTask(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ret, err := Db.InsTasks(ctx, &storage.Tasks{
		ID:         1000,
		Opened:     0,
		Closed:     0,
		AuthorID:   100,
		AssignedID: 101,
		Title:      "A",
		Content:    "AAA",
	})
	if err != nil {
		t.Error(err)
	}
	if ret.ID == 0 {
		err := fmt.Errorf("%s", "ошибка создания задачи")
		t.Error(err)
	}

}

func Test_DelTask(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := Db.DelTasks(ctx, 1000)
	if err != nil {
		t.Error(err)
	}
}

func Test_UpdTask(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ret, err := Db.UpdTasks(ctx, &storage.Tasks{
		ID:         1001,
		Opened:     0,
		Closed:     0,
		AuthorID:   101,
		AssignedID: 100,
		Title:      "B",
		Content:    "BBBB",
	})
	if err != nil {
		t.Error(err)

	}

	if ret != nil && ret.ID == 0 {
		err := fmt.Errorf("%s", "ошибка обновления задачи")
		t.Error(err)
	}
}

func Test_TaskOnID(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ret, err := Db.Tasks(ctx, 1001, 0, 0, 0, 100)
	if err != nil {
		t.Error(err)
	}
	if len(ret) == 0 {
		err := fmt.Errorf("%s", "задача не найдена")
		t.Error(err)

	}

}

func Test_TaskOnAuthorID(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ret, err := Db.Tasks(ctx, 0, 100, 0, 0, 100)
	if err != nil {
		t.Error(err)
	}
	if len(ret) == 0 {
		err := fmt.Errorf("%s", "задачи не найдены")
		t.Error(err)

	}

}

func Test_TaskOnAssignedID(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ret, err := Db.Tasks(ctx, 0, 0, 101, 0, 100)
	if err != nil {
		t.Error(err)
	}
	if len(ret) == 0 {
		err := fmt.Errorf("%s", "задачи не найдены")
		t.Error(err)

	}
}
