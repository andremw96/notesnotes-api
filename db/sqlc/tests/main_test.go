package db

import (
	db "andre/notesnotes-api/db/sqlc"
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:quipper123@localhost:5432/notesnotes?sslmode=disable"
)

var testQueries *db.Queries

func TestMain(m *testing.M) {
	// ENTRY POINT FOR ALL UNIT TEST
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to database")
	}

	testQueries = db.New(conn)

	os.Exit(m.Run())
}
