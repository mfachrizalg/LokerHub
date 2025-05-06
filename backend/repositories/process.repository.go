package repositories

import (
	"backend/config"
	"backend/models"
	"github.com/google/uuid"

	"gorm.io/gorm"
)

type ProcessRepository struct {
	DB *gorm.DB
}

func NewProcessRepository() *ProcessRepository {
	return &ProcessRepository{
		DB: config.DB,
	}
}

func (r *ProcessRepository) BeginTransaction() *gorm.DB {
	return r.DB.Begin()
}

// Create inserts a new process record using the provided transaction or the default DB
func (r *ProcessRepository) Create(tx *gorm.DB, process *models.Process) error {
	db := r.DB
	if tx != nil {
		db = tx // Use the transaction if provided
	}
	return db.Create(process).Error
}

// CheckIfExists checks if a candidate has already applied for a specific job
func (r *ProcessRepository) CheckIfExists(jobID, candidateID uuid.UUID) (bool, error) {
	var count int64
	err := r.DB.Model(&models.Process{}).Where("job_id = ? AND candidate_id = ?", jobID, candidateID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// FindAllByJobID retrieves all process records for a given job ID
func (r *ProcessRepository) FindAllByJobID(jobID uuid.UUID) ([]models.Process, error) {
	var processes []models.Process
	// Consider adding Preload("Candidate") if you need candidate details
	if err := r.DB.Where("job_id = ?", jobID).Find(&processes).Error; err != nil {
		return nil, err
	}
	return processes, nil
}
