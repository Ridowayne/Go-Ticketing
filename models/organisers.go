package models

import "gorm.io/gorm"

type Organiser struct {
	gorm.Model
	FullName string
	Email string
	OrganiserEvents []Event


}