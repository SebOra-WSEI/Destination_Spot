package main

import (
	"fmt"
	"github.com/SebOra-WSEI/auth/internal/route"
	"github.com/SebOra-WSEI/auth/start"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

const Port = ":8081"

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(err.Error())
		log.Fatal("Error loading .env file")
	}

	start.Database()

	r := gin.Default()
	route.InitRoutes(r)

	log.Fatal(r.Run(Port))
}
