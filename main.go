package main

import (
	"andre/notesnotes-api/api"
	db "andre/notesnotes-api/db/sqlc"
	"andre/notesnotes-api/util"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// load config
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// create server
	// connect to database
	// create Store object
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to database:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create the server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start the server:", err)
	}
}
