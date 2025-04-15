package router

// not used
import (
	"order-service/internal/handler"
	"order-service/internal/repository"
	"order-service/internal/usecase"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewRouter(db *mongo.Database) *gin.Engine {
	r := gin.Default()

	orderRepo := repository.NewOrderRepository(db)
	orderUC := usecase.NewOrderUsecase(orderRepo)
	orderHandler := handler.NewOrderHandler(orderUC)

	orderGroup := r.Group("/orders")
	{
		orderGroup.POST("", orderHandler.CreateOrder)
		orderGroup.GET("/:id", orderHandler.GetOrder)
		orderGroup.PATCH("/:id", orderHandler.UpdateOrderStatus)
		orderGroup.GET("", orderHandler.GetUserOrders)
	}

	return r
}
