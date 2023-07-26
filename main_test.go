package main

import (
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func TestServe(t *testing.T) {
	Serve()
}
