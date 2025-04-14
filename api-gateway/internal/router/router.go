package router

import (
	"api-gateway/internal/middleware"
	"bytes"

	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

const inventoryServiceURL = "http://localhost:8081"
const orderServiceURL = "http://localhost:8082"

func NewRouter(r *gin.Engine) {
	r.Use(middleware.AuthMiddleware())

	r.GET("/health", healthCheck)

	inventoryServiceGroup := r.Group("/inventory-service")
	{
		inventoryServiceGroup.POST("/products", func(c *gin.Context) {
			forwardRequest(c, "POST", inventoryServiceURL+"/products", c.Request.Body)
		})

	}

	orderServiceGroup := r.Group("/order-service")
	{
		orderServiceGroup.POST("/orders", func(c *gin.Context) {
			forwardRequest(c, "POST", orderServiceURL+"/orders", c.Request.Body)
		})

	}

}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func forwardRequest(c *gin.Context, method, url string, body io.Reader) {
	client := &http.Client{}
	var reqBody []byte

	// Чтение тела запроса, если оно есть
	if body != nil {
		var err error
		reqBody, err = ioutil.ReadAll(body)
		if err != nil {
			log.Println("Error reading request body:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
			return
		}
	}

	// Создание нового запроса
	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		log.Println("Error creating new request:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	// Копируем заголовки оригинального запроса
	for key, values := range c.Request.Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	// Отправка запроса в нужный микросервис
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error forwarding request:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to forward request"})
		return
	}
	defer resp.Body.Close()

	// Чтение тела ответа
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response"})
		return
	}

	// Отправка ответа клиенту
	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), respBody)
}
