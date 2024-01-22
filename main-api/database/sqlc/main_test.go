package db

import (
	"testing"
	"database/sql"
	"log"
	"os"
)

var testQueries *Queries

const {
	dbDriver = "postgres"
	dbSource = "postgresql://root:addison@172.18.48.243:5432/elo-system?sslmode=disable"
}

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource);
	if err != nil {
		log.Fatal("cannot connect to db: %w", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}