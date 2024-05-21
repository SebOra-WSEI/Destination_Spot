package main

import (
	 "github.com/SebOra-WSEI/Destination_spot/shared/example"
	 "github.com/gin-gonic/gin"
	 "log"
	 "net/http"
)

const Port = ":8080"

func main() {
	 example.Example()
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
