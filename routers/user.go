package routers

import (
	"github.com/carbans/netdebug/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.RouterGroup) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.GET("/user", controllers.CurrentUser)
}
