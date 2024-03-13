package tutorial

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/biubiupiuQAQ/bank/tree/master/db/util"
	_ "github.com/go-sql-driver/mysql"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testDB, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot Connect Database: ", err)
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
}
