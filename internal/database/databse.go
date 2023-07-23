package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"github.com/pressly/goose"
)

type DB interface {
	Conn() *sql.DB
	MigrateDB(string, ...string) error
}

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
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("unable to parse ennvironment variables: %e", err)
	}

	err = env.Parse(&c)
	if err != nil {
		log.Fatalf("unable to parse ennvironment variables: %e", err)
	}
	return &c
}

type Postgres struct {
	config *DbConfig
	db     *sql.DB
}

func NewEnv() (DB, error) {

	c := NewConfigFromEnv()
	conn, err := sql.Open("pgx", c.UrlPostgres())
	if err != nil {
		return nil, err
	}

	return &Postgres{
		db:     conn,
		config: c,
	}, nil
}

func (p *Postgres) Conn() *sql.DB {
	return p.db
}

func (p *Postgres) MigrateDB(cmd string, args ...string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()

	err := p.db.PingContext(ctx)
	if err != nil {
		return err
	}

	// goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	err = goose.Run(cmd, p.db, p.config.DirMigrateTemp, args...)
	if err != nil {
		log.Panicln(err)
	}

	return nil
}
