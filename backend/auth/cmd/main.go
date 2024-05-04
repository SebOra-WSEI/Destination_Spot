package main

import (
	 "fmt"
	 "github.com/gin-gonic/gin"
	 "github.com/jinzhu/gorm"
	 _ "github.com/jinzhu/gorm/dialects/mysql"
	 "log"
	 "net/http"
)

const Port = ":8081"

var Db *gorm.DB

func main() {
	 var err error

	 Db, err = gorm.Open("mysql", "user:0JcusWWFmfZyXrq8VomQ@tcp(database:3306)/destination_spot")
	 if err != nil {
		  fmt.Println("Error while connecting to the database:", err.Error())
		  panic("Failed to connect to the database")
	 }

	 fmt.Println("*** Auth Service successfully connected to the database ***")

	 r := gin.Default()

	 r.GET(
		  "/", func(c *gin.Context) {
				c.JSON(
					 http.StatusOK, gin.H{
						  "Service": "Auth",
					 },
				)
		  },
	 )

	 log.Fatal(r.Run(Port))
}
