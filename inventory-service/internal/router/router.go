package router

// не используется
import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"inventory-service/internal/handler"
	"inventory-service/internal/repository"
	"inventory-service/internal/usecase"
)

func NewRouter(db *mongo.Database) *gin.Engine {
	repo := repository.NewMongoProductRepo(db)
	usecase := usecase.NewProductUsecase(repo)
	handler := handler.NewProductHandler(usecase)

	r := gin.Default()

	r.POST("/products", handler.CreateProduct)
	r.GET("/products/:id", handler.GetProductByID)
	r.PATCH("/products/:id", handler.UpdateProduct)
	r.DELETE("/products/:id", handler.DeleteProduct)
	r.GET("/products", handler.ListProducts)

	return r
}
