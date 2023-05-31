package routers

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/scripts", script.Post)
		apiv1.DELETE("/scripts/:id", script.Delete)
		apiv1.PUT("/scripts/:id", script.Put)
		apiv1.GET("/scripts/:id", script.Get)
		apiv1.GET("/scripts", script.Get)
	}

	return r
}
