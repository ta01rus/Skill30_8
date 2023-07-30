package main

import (
	"log"
	"os"
	"runtime"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"github.com/ta01rus/Skill30_8/pkg/storage/postgres"

	"github.com/urfave/cli"
)

func init() {
	var err error
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	app := &cli.App{
		Name:  "pgx generator db",
		Usage: "PGX",
		Commands: []cli.Command{
			{
				Name:  "down",
				Usage: "",
				Action: func(*cli.Context) error {
					Migrate("down")
					return nil
				},
			},
			{
				Name:  "up",
				Usage: "",
				Action: func(*cli.Context) error {
					Migrate("up")
					return nil
				},
			},
			{
				Name:  "redo",
				Usage: "",
				Action: func(*cli.Context) error {
					Migrate("redo")
					return nil
				},
			},
			{
				Name:  "reset",
				Usage: "",
				Action: func(*cli.Context) error {
					Migrate("reset")
					return nil
				},
			},
			{
				Name:  "status",
				Usage: "",
				Action: func(*cli.Context) error {
					Migrate("status")
					return nil
				},
			},
			{
				Name:  "version",
				Usage: "",
				Action: func(*cli.Context) error {
					Migrate("version")
					return nil
				},
			},
			{
				Name:  "create",
				Usage: "",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name: "name, n",
						// Value: "up",
						Usage: "set name",
					},
					&cli.BoolFlag{
						Name: "sql",
						// Value: "up",
						Usage: "set type sql",
					},
					&cli.BoolFlag{
						Name:  "go",
						Usage: "set type go",
					},
				},
				Action: func(c *cli.Context) error {
					switch {
					case c.Bool("down"):
						Migrate("create", c.String("name"), "go")
					default:
						Migrate("create", c.String("name"), "sql")
					}
					return nil
				},
			},
		},
	}
	app.Run(os.Args)
}

func Migrate(cmd string, args ...string) {
	db, err := postgres.NewEnv()
	if err != nil {
		log.Fatal(err)
	}
	err = db.MigrateDB(cmd, args...)
	if err != nil {
		log.Fatal(err)
	}
}
