package routes

import (
	"net/http"

	"github.com/Adrwaan/stock-api-go/models"

	"github.com/gin-gonic/gin"
)

func productRoutes(r *gin.RouterGroup) {
	r.GET("/products", func(c *gin.Context) {
		c.JSON(http.StatusOK, models.GetProducts())
	})

	r.GET("/product/:id", func(c *gin.Context) {
		id := c.Param("id")
		product, err := models.GetProductWithID(id)

		if err != nil {
			logger.Error(err)
			c.Status(http.StatusInternalServerError)
			c.Abort()
		}

		c.JSON(http.StatusOK, product)
	})

	r.POST("/product", func(c *gin.Context) {
		var body models.Product
		c.ShouldBindJSON(&body)

		product, err := models.AddNewProduct(body.Name, body.Qtd)

		if err != nil {
			logger.Error(err)
			c.Status(http.StatusInternalServerError)
			c.Abort()
		}

		c.JSON(http.StatusCreated, product)
	})

	r.PUT("/product/:id", func(c *gin.Context) {
		id := c.Param("id")

		var body models.Product
		c.ShouldBindJSON(&body)

		product, err := models.UpdateProduct(id, body.Name, body.Qtd)

		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			c.Abort()
		}

		c.JSON(http.StatusAccepted, product)
	})

	r.DELETE("/product/:id", func(c *gin.Context) {
		id := c.Param("id")

		err := models.DeleteProduct(id)

		if err != nil {
			logger.Error(err)
			c.Status(http.StatusInternalServerError)
			c.Abort()
		}

		c.Status(http.StatusAccepted)
	})
}
