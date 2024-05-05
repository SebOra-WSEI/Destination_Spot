package main

import (
	"fmt"
	"github.com/SebastianOraczek/auth/internal/route"
	"github.com/SebastianOraczek/auth/startup"
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

	startup.Database()

	r := gin.Default()
	route.InitRoutes(r)

	log.Fatal(r.Run(Port))
}
