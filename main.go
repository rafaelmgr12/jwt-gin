package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rafaelmgr12/jwt-gin/controllers"
	"github.com/rafaelmgr12/jwt-gin/models"
)

func main() {

	models.ConnectDataBase()

	r := gin.Default()

	public := r.Group("/api")

	public.POST("/register", controllers.Register)

	r.Run(":8080")
}
