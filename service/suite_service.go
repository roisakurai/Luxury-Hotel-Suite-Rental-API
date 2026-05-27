package service

import (
	"p2-ip-hotel-rental/models"
	"p2-ip-hotel-rental/repository"
)

type SuiteService interface {
	GetAllSuites() ([]models.Suite, error)
}

type suiteService struct {
	suiteRepo repository.SuiteRepository
}

func NewSuiteService(suiteRepo repository.SuiteRepository) SuiteService {
	return &suiteService{suiteRepo}
}

func (s *suiteService) GetAllSuites() ([]models.Suite, error) {
	return s.suiteRepo.FindAll()
}
