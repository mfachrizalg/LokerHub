package services

import (
	"backend/dtos"
	"backend/models"
	"backend/repositories"
	"errors"
	"github.com/gofiber/fiber/v2/log"
)

type CandidateService struct {
	repo *repositories.CandidateRepository
}

func NewCandidateService(repo *repositories.CandidateRepository) *CandidateService {
	return &CandidateService{
		repo: repo,
	}
}

func (s *CandidateService) RegisterCandidate(req *dtos.RegisterCandidateRequest) (*dtos.MessageResponse, error) {
	tx := s.repo.BeginTransaction()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	candidate := models.Candidate{
		Name:      req.Name,
		Education: req.Education,
		Handphone: req.Handphone,
	}

	if err := s.repo.Update(&candidate); err != nil {
		tx.Rollback()
		log.Error("Error creating candidate: ", err)
		return nil, errors.New("failed to process registration")
	}

	if err := tx.Commit().Error; err != nil {
		log.Error("Error committing transaction: ", err)
		return nil, errors.New("failed to process registration")
	}

	return &dtos.MessageResponse{
		Message: "Candidate registered successfully",
	}, nil
}
