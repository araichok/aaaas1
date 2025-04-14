package main

import (
	"github.com/gin-gonic/gin"
	"inventory-service/database"
	"inventory-service/internal/router"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	db := database.InitMongo()
	r := router.NewRouter(db)
	r.Run(":8081")
}
