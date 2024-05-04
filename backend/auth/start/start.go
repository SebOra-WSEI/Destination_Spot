package start

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
)

const ConnectionStringEnvName string = "CONNECTION_STRING"

var Db *gorm.DB

func Database() {
	var err error

	connStr := os.Getenv(ConnectionStringEnvName)
	if connStr == "" {
		log.Fatalf("Environment variable %v is empty", connStr)
	}

	Db, err = gorm.Open("mysql", connStr)
	if err != nil {
		fmt.Println("Error while connecting to the database:", err.Error())
		panic("Failed to connect to the database")
	}

	fmt.Println("*** Auth Service successfully connected to the database ***")
}
