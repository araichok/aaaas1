package router

import (
	inventorypb "api-gateway/api-gateway/proto/inventorypb"
	orderpb "api-gateway/api-gateway/proto/orderpb"
	"api-gateway/internal/middleware"
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

var (
	inventoryClient inventorypb.InventoryServiceClient
	orderClient     orderpb.OrderServiceClient
)

func InitGRPCClients() {
	conn1, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
	inventoryClient = inventorypb.NewInventoryServiceClient(conn1)

	conn2, _ := grpc.Dial("localhost:50052", grpc.WithInsecure())
	orderClient = orderpb.NewOrderServiceClient(conn2)
}

func NewRouter(r *gin.Engine) {
	r.Use(middleware.AuthMiddleware())

	r.GET("/health", healthCheck)

	inventoryServiceGroup := r.Group("/inventory-service")
	{
		inventoryServiceGroup.POST("/products", func(c *gin.Context) {
			var req inventorypb.CreateProductRequest
			if err := c.BindJSON(&req); err != nil {
				c.JSON(400, gin.H{"error": "invalid request"})
				return
			}

			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			res, err := inventoryClient.CreateProduct(ctx, &req)
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}

			c.JSON(200, res)
		})

		inventoryServiceGroup.GET("/products/:id", func(c *gin.Context) {
			id := c.Param("id")
			req := &inventorypb.GetProductRequest{Id: id}

			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			res, err := inventoryClient.GetProductByID(ctx, req)
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}

			c.JSON(200, res)
		})
		inventoryServiceGroup.PUT("/products/:id", func(c *gin.Context) {
			id := c.Param("id")
			var req inventorypb.UpdateProductRequest
			if err := c.BindJSON(&req); err != nil {
				c.JSON(400, gin.H{"error": "invalid request"})
				return
			}

			req.Product.Id = id

			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			res, err := inventoryClient.UpdateProduct(ctx, &req)
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}

			c.JSON(200, res)
		})
		inventoryServiceGroup.DELETE("/products/:id", func(c *gin.Context) {
			id := c.Param("id")
			req := &inventorypb.DeleteProductRequest{Id: id}

			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			_, err := inventoryClient.DeleteProduct(ctx, req)
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}

			c.JSON(200, gin.H{"message": "Product deleted"})
		})

		inventoryServiceGroup.GET("/products", func(c *gin.Context) {
			req := &inventorypb.ListProductsRequest{}

			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			res, err := inventoryClient.ListProducts(ctx, req)
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}

			c.JSON(200, res)
		})
	}

	orderServiceGroup := r.Group("/order-service")
	{
		orderServiceGroup.POST("/orders", func(c *gin.Context) {
			var req orderpb.CreateOrderRequest
			if err := c.BindJSON(&req); err != nil {
				c.JSON(400, gin.H{"error": "invalid request"})
				return
			}

			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			res, err := orderClient.CreateOrder(ctx, &req)
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}

			c.JSON(200, res)
		})

		orderServiceGroup.GET("/orders/:id", func(c *gin.Context) {
			id := c.Param("id")
			req := &orderpb.GetOrderRequest{Id: id}

			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			res, err := orderClient.GetOrderByID(ctx, req)
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}

			c.JSON(200, res)
		})

		orderServiceGroup.PUT("/orders/:id/status", func(c *gin.Context) {
			id := c.Param("id")
			var req orderpb.UpdateOrderStatusRequest
			if err := c.BindJSON(&req); err != nil {
				c.JSON(400, gin.H{"error": "invalid request"})
				return
			}

			req.Id = id

			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			res, err := orderClient.UpdateOrderStatus(ctx, &req)
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}

			c.JSON(200, res)
		})

		orderServiceGroup.GET("/orders/user/:user_id", func(c *gin.Context) {
			userId := c.Param("user_id")
			req := &orderpb.ListOrdersRequest{UserId: userId}

			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			res, err := orderClient.ListUserOrders(ctx, req)
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}

			c.JSON(200, res)
		})
	}
}

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{"status": "OK"})
}
