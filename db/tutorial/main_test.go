package tutorial

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbDriver = "mysql"
	dbSource = "root:123456@tcp(localhost:3307)/bank?parseTime=true"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot Connect Database: ", err)
	}

	testQueries = New(conn)
	os.Exit(m.Run())
}
