package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rafaelmgr12/jwt-gin/controllers"
	"github.com/rafaelmgr12/jwt-gin/middlewares"
	"github.com/rafaelmgr12/jwt-gin/models"
)

func main() {

	models.ConnectDataBase()

	r := gin.Default()

	public := r.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)

	r.Run(":8080")

}
