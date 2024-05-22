package main

import (
	 "github.com/SebOra-WSEI/Destination_spot/auth/database"
	 "github.com/SebOra-WSEI/Destination_spot/auth/internal/route"
	 "github.com/SebOra-WSEI/Destination_spot/shared/env"
	 "github.com/gin-gonic/gin"
	 "log"
)

const Port = ":8081"

func main() {
	 if err := env.Load("../.env"); err != nil {
		  log.Fatal("Error loading .env file")
	 }

	 database.Start()

	 r := gin.Default()
	 route.Init(r)

	 log.Fatal(r.Run(Port))
}
