package main

import (
	 "fmt"
	 "github.com/SebOra-WSEI/auth/start"
	 "github.com/gin-gonic/gin"
	 "github.com/joho/godotenv"
	 "log"
	 "net/http"
)

const Port = ":8081"

func main() {
	 if err := godotenv.Load(); err != nil {
		  fmt.Println(err.Error())
		  log.Fatal("Error loading .env file")
	 }

	 start.Database()

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
