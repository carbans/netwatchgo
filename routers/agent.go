package routers

import (
	"github.com/carbans/netdebug/controllers"
	"github.com/gin-gonic/gin"
)

func AgentRoute(r *gin.RouterGroup) {
	r.GET("/agents", controllers.FindAgents)
	r.GET("/agents/:id", controllers.FindAgent)
	r.POST("/agents", controllers.CreateAgent)
	r.PATCH("/agents/:id", controllers.UpdateAgent)
	r.DELETE("/agents/:id", controllers.DeleteAgent)

}
