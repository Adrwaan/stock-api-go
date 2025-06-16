package routes

import "github.com/gin-gonic/gin"

func initializeRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		productRoutes(v1)
	}
}
