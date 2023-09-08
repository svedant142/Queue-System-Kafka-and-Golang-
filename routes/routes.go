package routes

import (
	"message-queue-system/business/product/repository"
	"message-queue-system/business/product/usecase"
	productHTTP "message-queue-system/business/product/webhttp"
	"message-queue-system/db"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	productRepo := repository.NewProductRepo(db.DbClient)
	productUC :=  usecase.NewProductUC(productRepo)
	productHTTP.Init(router, productUC)
}
