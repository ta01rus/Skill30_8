package postgres

import (
	"log"
	"testing"

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

}

func Test_DelTask(t *testing.T) {

}

func Test_UpdTask(t *testing.T) {

}

func Test_TaskOnID(t *testing.T) {

}

func Test_TaskOnAuthorID(t *testing.T) {

}

func Test_TaskOnAssignedID(t *testing.T) {

}
