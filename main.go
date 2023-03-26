package main

import (
	"log"
	"net/http"
	"runtime"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/carbans/netdebug/config"
	"github.com/carbans/netdebug/docs"
	"github.com/carbans/netdebug/middlewares"
	"github.com/carbans/netdebug/models"
	"github.com/carbans/netdebug/routers"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	// Set Configuration
	config, err := config.LoadConfig("")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// Connect to database
	models.ConnectDatabase(config)

	// Routes
	v1 := r.Group("/v1")
	docs.SwaggerInfo.BasePath = "/v1"
	routers.UserRoute(v1)
	v1.Use(middlewares.JwtAuthMiddleware())
	routers.AgentRoute(v1)
	routers.ApiKeyRoute(v1)

	r.LoadHTMLGlob("./public/html/*")
	r.Static("/public", "./public")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"packageVersion": "v0.01",
			"goVersion":      runtime.Version(),
		})
	})
	//docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{})
	})

	// Run the server
	r.Run()
}
