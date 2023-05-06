package main

import (
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
	router.Run()
	
}