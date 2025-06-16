package routes

import (
	"github.com/Adrwaan/stock-api-go/config"

	"github.com/gin-gonic/gin"
)

var logger *config.Logger = config.GetLogger()

func Bootstrap() {
	r := gin.Default()

	initializeRoutes(r)

	r.Run(":9090")
}

// http://localhost:9090/products
