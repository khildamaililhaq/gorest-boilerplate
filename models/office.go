package models

import (
	"gorm.io/gorm"
)

type Office struct {
	gorm.Model
	Name       	string	`gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Address   	string	`gorm:"->;<-;notNull" json:"address"`
	Description	string	`gorm:"type:text;->;<-" json:"description"`
	Users		[]User	`json:"users"`
}