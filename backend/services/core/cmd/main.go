package main

import (
	 "github.com/gin-gonic/gin"
	 "log"
	 "net/http"
)

const Port = ":8080"

func main() {
	 r := gin.Default()

	 r.GET(
		  "/", func(c *gin.Context) {
				c.JSON(
					 http.StatusOK, gin.H{
						  "Service": "Core",
					 },
				)
		  },
	 )

	 log.Fatal(r.Run(Port))
}
