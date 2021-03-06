package db

import (
	db "andre/notesnotes-api/db/sqlc"
	"andre/notesnotes-api/util"
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
var testDb *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../..")
	if err != nil {
		log.Fatal("cannot load config")
	}

	// ENTRY POINT FOR ALL UNIT TEST
	testDb, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to database")
	}

	testQueries = db.New(testDb)

	os.Exit(m.Run())
}
