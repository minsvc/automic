package routers

import (
	"automic/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/scripts", v1.Script{}.Create)
		apiv1.DELETE("/scripts/:id", v1.Script{}.Delete)
		apiv1.PUT("/scripts/:id", v1.Script{}.Update)
		apiv1.GET("/scripts/:id", v1.Script{}.Get)
		apiv1.GET("/scripts", v1.Script{}.List)
	}

	return r
}
