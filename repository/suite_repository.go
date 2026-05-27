package repository

import (
	"p2-ip-hotel-rental/models"

	"gorm.io/gorm"
)

type SuiteRepository interface {
	FindAll() ([]models.Suite, error)
	FindByID(id uint) (*models.Suite, error)
}

type suiteRepository struct {
	db *gorm.DB
}

func NewSuiteRepository(db *gorm.DB) SuiteRepository {
	return &suiteRepository{db}
}

func (r *suiteRepository) FindAll() ([]models.Suite, error) {
	var suites []models.Suite
	err := r.db.Find(&suites).Error
	return suites, err
}

func (r *suiteRepository) FindByID(id uint) (*models.Suite, error) {
	var suite models.Suite
	err := r.db.First(&suite, id).Error
	return &suite, err
}
