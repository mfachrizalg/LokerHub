package services

import (
	"backend/dtos"
	"backend/models"
	"backend/repositories"
	"errors"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
)

type CompanyService struct {
	companyRepo *repositories.CompanyRepository
}

func NewCompanyService(companyRepo *repositories.CompanyRepository) *CompanyService {
	return &CompanyService{
		companyRepo: companyRepo,
	}
}

func (s *CompanyService) RegisterCompany(req *dtos.RegisterCompanyRequest) (*dtos.MessageResponse, error) {
	tx := s.companyRepo.BeginTransaction()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	company := models.Company{
		Name:     req.Name,
		Location: req.Location,
		Industry: req.Industry,
		Logo:     req.Logo,
	}

	if err := s.companyRepo.Create(&company); err != nil {
		tx.Rollback()
		log.Error("Error creating company: ", err)
		return nil, errors.New("failed to create company")
	}

	if err := tx.Commit().Error; err != nil {
		log.Error("Error committing transaction: ", err)
		return nil, errors.New("failed to create company")
	}

	return &dtos.MessageResponse{
		Message: "Company created successfully",
	}, nil
}

func (s *CompanyService) GetAllCompany() (*dtos.GetAllCompanyResponse, error) {
	var companies []models.Company
	if err := s.companyRepo.FindAll(&companies); err != nil {
		log.Error("Error getting all companies: ", err)
		return nil, errors.New("failed to get all companies")
	}

	response := make(dtos.GetAllCompanyResponse, len(companies))
	for i, company := range companies {
		response[i] = struct {
			ID   uuid.UUID `json:"id"`
			Name string    `json:"name"`
		}{
			ID:   company.ID,
			Name: company.Name,
		}
	}

	return &response, nil
}
