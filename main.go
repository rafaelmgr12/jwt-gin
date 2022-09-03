package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rafaelmgr12/jwt-gin/controllers"
)

func main() {

	r := gin.Default()

	public := r.Group("/api")

	public.POST("/register", controllers.Register)

	r.Run(":8080")
}
