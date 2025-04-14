package main

import (
	"api-gateway/internal/router"
	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"time"
)

func main() {
	// Инициализация логирования
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal("unable to create logger", zap.Error(err))
	}
	defer logger.Sync()

	// Инициализация Gin с логированием
	r := gin.Default()
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true)) // Логирование запросов

	// Инициализация маршрутов
	router.NewRouter(r)

	// Запуск сервера
	r.Run(":8080")
}
