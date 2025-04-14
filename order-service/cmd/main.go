package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"order-service/database"
	"order-service/internal/router"
)

func main() {

	gin.SetMode(gin.ReleaseMode)

	db, err := database.Connect()
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	r := router.NewRouter(db)
	if err := r.Run(":8082"); err != nil {
		log.Fatal("Server failed:", err)
	}
}
