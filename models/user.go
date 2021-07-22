package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name     	 string			`gorm:"->;<-;notNull" json:"name"`
	Email        string			`gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password     string			`gorm:"->;<-;notNull" json:"-"`
	Birthday     *time.Time		`gorm:"->;<-;notNull" json:"birthday"`
	OfficeID     string			`gorm:"->;<-;index" json:"officeID"`
	Token		 string			`gorm:"-" json:"token,omitempty"`
	Office		 Office 		`gorm:"foreignKey:OfficeID" json:"office"`
}
