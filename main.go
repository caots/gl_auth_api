// main.go

package main

import (
	"auth_api/controllers"
	"auth_api/database"
	"auth_api/middlewares"
	"auth_api/models"
	"log"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	//here
	api := r.Group("/api")
	{
		public := api.Group("/v1")
		{
			public.POST("/login", controllers.Login)
			public.POST("/signup", controllers.Signup)
		}
	}

	// here
	protected := api.Group("/protected").Use(middlewares.Authz())
	{
		protected.GET("/profile", controllers.Profile)
	}

	return r
}

func main() {
	err := database.InitDatabase()
	if err != nil {
		log.Fatalln("could not create database", err)
	}

	database.GlobalDB.AutoMigrate(&models.User{})

	r := setupRouter()
	r.Run(":4000")
}
