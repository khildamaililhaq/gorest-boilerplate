package dto

import "time"

type UserUpdateDTO struct {
	ID uint64 `json:"id" form:"id" binding:"required"`
	Name string `json:"name" form:"name" binding:"required"`
	Email string `json:"email" form:"email" binding:"required" validate:"email"`
	Password string `json:"password,omitempty" form:"password,omitempty" validate:"min:6"`
	Birthday *time.Time `json:"birthday" form:"birthday" binding:"required"`
	OfficeID string `json:"officeID" form:"officeID" binding:"required"`
}

type UserCreatedDTO struct {
	Name string `json:"name" form:"name" binding:"required"`
	Email string `json:"email" form:"email" binding:"required" validate:"email"`
	Password string `json:"password,omitempty" form:"password,omitempty"  validate:"min:6" binding:"required"`
	Birthday *time.Time `json:"birthday" form:"birthday" binding:"required"`
	OfficeID string `json:"officeID" form:"officeID" binding:"required"`
}
