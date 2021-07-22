package dto

type OfficeUpdateDTO struct {
	ID uint64 `json:"id" form:"id" binding:"required"`
	Name string `json:"name" form:"name" binding:"required"`
	Address string `json:"address" form:"address" binding:"required"`
	Description string `json:"address" form:"address"  binding:"required"`
}

type OfficeCreatedDTO struct {
	Name string `json:"name" form:"name" binding:"required"`
	Address string `json:"address" form:"address" binding:"required"`
	Description string `json:"address" form:"address"  binding:"required"`
}
