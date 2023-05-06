package models

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	EventName string
	EventOrganiser string
	OrganiserEmail string
	Venue string
	EventDateate uint64
	Price uint16
	SeatsAvailable uint64
	SeatsBooked uint64
}