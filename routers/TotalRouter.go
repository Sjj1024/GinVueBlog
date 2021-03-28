package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shen.com/ginvue/controller"
)

func TotalRouter(r *gin.Engine) *gin.Engine {
	userGroup := r.Group("/user")
	{
		userGroup.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "chenggong")
		})
		userGroup.POST("/register", controller.Register)
		userGroup.POST("/login", controller.Login)
	}
	return r
}
