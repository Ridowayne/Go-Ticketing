package controllers

import (
	"example/funtion/initializers"
	"example/funtion/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateEvent(c *gin.Context) {
	userMail,  _ := c.Get("user")
	var body struct{
	EventName string
	EventOrganiser string
	Venue string
	EventDate uint64
	Price uint16
	SeatsAvailable uint64	
	}
	if c.Bind(body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Kindly provide all the necessary ino"})
	}
	event := models.Event{EventName: body.EventName, EventOrganiser: userMail.(models.User).FirstName, OrganiserEmail: userMail.(models.User).Email, Venue: body.Venue, EventDate: body.EventDate, Price: body.Price, SeatsAvailable: body.SeatsAvailable, SeatsBooked: 0 }

	newEvent := initializers.DB.Create(&event)

	if newEvent.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating event"})
	}

}
func GetAllEvents (c *gin.Context) {
	var events [] models.Event
	initializers.DB.Find(events)
	c.JSON(http.StatusOK, gin.H{"events": events})

}

func GetEvent (c *gin.Context) {
	id:= c.Param("id")
	var event models.Event
	initializers.DB.First(&event, id)
	c.JSON(http.StatusOK, gin.H{"event": event})
}

func GetEventByDate (c *gin.Context) {

}

func GetEventVenue (c *gin.Context) {
	venue:= c.Query("venue")
	var events []gorm.Model
	initializers.DB.Find(&events, "venue = ?", venue)


}

func GetEventByOrganiser (c *gin.Context) {}