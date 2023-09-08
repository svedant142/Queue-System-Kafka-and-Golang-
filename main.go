package main

import (
	"fmt"
	"message-queue-system/clients/kafka"
	"message-queue-system/db"
	"message-queue-system/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	err := db.InitMYSQL()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	
	go kafka.Consume()

	router := gin.Default()
	routes.InitRoutes(router)
	
	err =	router.Run(":8080")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}