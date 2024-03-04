package main

import (
	"database/sql"
	"log"

	"github.com/biubiupiuQAQ/bank/tree/master/api"
	db "github.com/biubiupiuQAQ/bank/tree/master/db/tutorial"
	_ "github.com/go-sql-driver/mysql"
)

const (
	dbDriver      = "mysql"
	dbSource      = "root:123456@tcp(localhost:3307)/bank?parseTime=true"
	serverAddress = "localhost:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot Connect Database: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
