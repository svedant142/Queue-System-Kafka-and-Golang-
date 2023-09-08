package webhttp

import (
	"message-queue-system/business/product/usecase"
	Requests "message-queue-system/domain/dto/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

var productUC usecase.IProductUC

func Init(router *gin.Engine, puc usecase.IProductUC) {
	productUC = puc
	router.POST("/user/item",insertProduct)
}


func insertProduct(c *gin.Context) {
	var req Requests.InsertProduct
	if err := c.ShouldBindQuery(&req) ; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : "invalid payload",
			"status" : false,
		})
		return
	}
	if err := req.Validate() ; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message" : "validation failed",
			"error" : err.Error(),
			"status" : false,
		})
		return
	}

	err := productUC.InsertProduct(c,req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message" : "insertion failed",
			"status" : false,
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"message" : "successful insertion",
		"status" : true,
	})
}