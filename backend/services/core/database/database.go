package database

import (
	"database/sql"
	"fmt"
	"github.com/SebOra-WSEI/Destination_spot/shared/env"
	"github.com/XSAM/otelsql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"log"
)

var Db *gorm.DB
var DbSQL *sql.DB

func Start() {
	var err error

	connStr, err := env.GetEnvVariableByName(env.ConnectionString)

	DbSQL, err = otelsql.Open(
		"mysql",
		connStr,
		otelsql.WithSQLCommenter(true),
		otelsql.WithAttributes(semconv.DBSystemMySQL),
	)

	if err != nil {
		log.Fatalf("Error connecting to the database:", "error", err.Error())
	}

	fmt.Println("*** Core Service successfully connected to the database ***")
}
