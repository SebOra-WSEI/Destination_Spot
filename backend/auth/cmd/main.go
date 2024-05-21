package main

import (
	"fmt"
	"github.com/SebOra-WSEI/Destination_spot/auth/database"
	"github.com/SebOra-WSEI/Destination_spot/auth/internal/route"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

const Port = ":8081"

func main() {
	if err := godotenv.Load("../.env"); err != nil {
		fmt.Println(err.Error())
		log.Fatal("Error loading .env file")
	}

	database.Start()

	r := gin.Default()
	route.Init(r)

	log.Fatal(r.Run(Port))
}
