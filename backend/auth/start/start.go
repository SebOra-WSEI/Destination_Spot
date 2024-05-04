package start

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

var Db *gorm.DB

func Database() {
	var err error

	conn := os.Getenv("CONNECTION_STRING")

	Db, err = gorm.Open("mysql", conn)
	if err != nil {
		fmt.Println("Error while connecting to the database:", err.Error())
		panic("Failed to connect to the database")
	}

	fmt.Println("*** Auth Service successfully connected to the database ***")
}
