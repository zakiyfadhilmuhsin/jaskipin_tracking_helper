package models

import "github.com/jinzhu/gorm"

type Destination struct {
	gorm.Model
	Country  string
	Whatsapp string
	Email    string
	Facebook string
	Group    string
	Trainer  string
}
