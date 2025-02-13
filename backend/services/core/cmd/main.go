package main

import (
	"github.com/SebOra-WSEI/Destination_spot/core/database"
	"github.com/SebOra-WSEI/Destination_spot/core/internal/route"
	"github.com/SebOra-WSEI/Destination_spot/core/swo"
	"github.com/SebOra-WSEI/Destination_spot/shared/env"
	"github.com/gin-gonic/gin"
	"log"
)

const Port = ":8080"

func main() {
	if err := env.Load("../.env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	cb, err := swo.Start()
	defer cb()

	if err != nil {
		log.Fatalf("Error while connecting with SWO: %v", err)
	}

	database.Start()

	r := gin.Default()
	route.Init(r)

	log.Fatal(r.Run(Port))
}
