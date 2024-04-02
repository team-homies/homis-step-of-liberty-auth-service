package patient

import (
	"main/database/entity"

	"gorm.io/gorm"
)

type PatientRepository interface {
	GetAll() ([]entity.Patient, error)
	GetByID(id uint) (*entity.Patient, error)
	Create(patient *entity.Patient) error
	Update(patient *entity.Patient) error
	Delete(id uint) error
}

type gormPatientRepository struct {
	db *gorm.DB
}

func NewPatientRepository(db *gorm.DB) PatientRepository {
	return &gormPatientRepository{db}
}

func (g *gormPatientRepository) Create(patient *entity.Patient) error {
	panic("")
}

func (g *gormPatientRepository) Delete(id uint) error {
	panic("")
}

func (g *gormPatientRepository) GetAll() ([]entity.Patient, error) {
	panic("")
}

func (g *gormPatientRepository) GetByID(id uint) (*entity.Patient, error) {
	panic("")
}

func (g *gormPatientRepository) Update(patient *entity.Patient) error {
	panic("")
}
