package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5"
	"github.com/ta01rus/Skill30_8/pkg/storage"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type DbConfig struct {
	User   string `env:"DB_USER"`
	Passw  string `env:"DB_PASSW"`
	Host   string `env:"DB_HOST"`
	Port   string `env:"DB_PORT"`
	DbName string `env:"DB_NAME"`
	Shema  string `env:"DB_SHEMA"`

	DirMigrateTemp string `env:"DB_MIGRATE_TMP"`
}

func (dc *DbConfig) UrlPostgres() string {
	return fmt.Sprintf(`postgres://%s:%s@%s:%s/%s`, dc.User, dc.Passw, dc.Host, dc.Port, dc.DbName)
}

func NewConfigFromEnv() *DbConfig {
	c := DbConfig{}
	godotenv.Load()
	err := env.Parse(&c)
	if err != nil {
		log.Fatalf("unable to parse ennvironment variables: %e", err)
	}
	return &c
}

type Postgres struct {
	*sql.DB
	config *DbConfig
}

func NewEnv() (storage.DB, error) {

	c := NewConfigFromEnv()
	db, err := sql.Open("pgx", c.UrlPostgres())
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Postgres{
		DB:     db,
		config: c,
	}, nil
}
