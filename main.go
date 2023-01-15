package main

import (
	"gin-vue/controllers"
	"gin-vue/middlewares"
	"gin-vue/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDatabase()

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "GET", "DELETE"},
		AllowHeaders: []string{"Content-Type, Access-Control-Allow-Origin, Access-Control-Allow-Headers, x-access-token, x-api-key , Accept"},
	}))
	public := r.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/auth/signin", controllers.Login)
	public.GET("/profile", controllers.GetUserProfile)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)

	r.Run(":8080")
}
