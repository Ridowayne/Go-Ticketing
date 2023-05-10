package main

import (
	"example/funtion/controllers"
	"example/funtion/initializers"

	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnv()
	initializers.ConnectDB()
	
}

func main() {
	router:= gin.Default()
	router.GET("/")
	router.POST("/singup", controllers.CreateUser)
	router.POST("/login", controllers.SignIn)
	router.POST("/events", controllers.CreateEvent)
	router.GET("/events", controllers.GetAllEvents)
	router.GET("/events/id", controllers.GetEvent)
	router.Run()
	
}