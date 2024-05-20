package database

import (
	"fmt"
	"github.com/SebOra-WSEI/Destination_spot/internal/env"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

func Start() {
	var err error

	connStr := env.GetEnvVariableByName(env.ConnectionString)

	Db, err = gorm.Open("mysql", connStr)
	if err != nil {
		fmt.Println("Error while connecting to the database:", err.Error())
		panic("Failed to connect to the database")
	}

	fmt.Println("*** Auth Service successfully connected to the database ***")
}
