package models

import "gorm.io/gorm"

type Ticket struct {
	gorm.Model
	EventName string
	EventID string
	OrganiserID string
	NameOfBooker string
	UserId string
	BookerEmail string
	Price uint64
	SeatNumber uint64
	PaymentId string

}