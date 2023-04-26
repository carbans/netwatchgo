package routers

import (
	"github.com/carbans/netdebug/controllers"
	"github.com/gin-gonic/gin"
)

func ApiKeyRoute(r *gin.RouterGroup) {
	r.POST("/apikeys", controllers.CreateApiKey)
}
