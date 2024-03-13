package main

import (
	"database/sql"
	"log"

	"github.com/biubiupiuQAQ/bank/tree/master/api"
	db "github.com/biubiupiuQAQ/bank/tree/master/db/tutorial"
	"github.com/biubiupiuQAQ/bank/tree/master/db/util"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot Connect Database: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
