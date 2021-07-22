package services

import (
	"github.com/khildamaililhaq/gorest-boilerplate/dto"
	"github.com/khildamaililhaq/gorest-boilerplate/models"
)

type OfficeService interface {
	Index() []models.Office
	Create(office dto.OfficeCreatedDTO) models.Office
	Update(officeID string, office dto.OfficeUpdateDTO) models.Office
	Get(officeID string) models.Office
	Destroy(officeID string) models.Office
}
