package routes

import (
	"github.com/gin-gonic/gin"
	c "api/controller"
	m "api/middleware"
)

func Router(r *gin.Engine) *gin.Engine {
	r.Use(m.SetHeader())
	v1 := r.Group("/v1")
	{
		v1.GET("/login", c.Login)
	}
    return r
}
