package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mfhmiii/go-rest-api-1/controllers/productcontroller"
	"github.com/mfhmiii/go-rest-api-1/models"
)

func main() {
	r := gin.Default()
	models.DBconn()

	//membuat routes
	r.GET("api/products", productcontrollers.Index)
	r.GET("api/products/:id", productcontrollers.Show)
	r.POST("api/product", productcontrollers.Create)
	r.PUT("api/product/:id", productcontrollers.Update)
	r.DELETE("api/product", productcontrollers.Delete)

	r.Run()
}
