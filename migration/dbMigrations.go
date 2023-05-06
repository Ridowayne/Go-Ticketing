package main

import (
	"example/funtion/initializers"
	"example/funtion/models"
)
func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Event{}, &models.Ticket{},)
}