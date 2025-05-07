package repositories

import (
	"backend/config"
	"backend/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type JobRepository struct {
	DB *gorm.DB
}

func NewJobRepository() *JobRepository {
	return &JobRepository{
		DB: config.DB,
	}
}

func (r *JobRepository) BeginTransaction() *gorm.DB {
	return r.DB.Begin()
}

// FindAll retrieves all jobs
// FindAll retrieves all jobs without pagination
func (r *JobRepository) FindAll() ([]models.Job, error) {
	var jobs []models.Job

	if err := r.DB.Find(&jobs).Error; err != nil {
		return nil, err
	}

	return jobs, nil
}

func (r *JobRepository) FindById(id uuid.UUID) (*models.Job, error) {
	var job models.Job

	if err := r.DB.Where("id = ?", id).First(&job).Error; err != nil {
		return nil, err
	}

	return &job, nil
}

func (r *JobRepository) Create(job *models.Job) error {
	return r.DB.Create(job).Error
}

func (r *JobRepository) Update(tx *gorm.DB, job *models.Job) error {
	db := r.DB
	if tx != nil {
		db = tx // Use the transaction if provided
	}
	// Use Save to update all fields, or Updates to update specific fields
	// Save ensures all fields are updated, including zero values.
	return db.Save(job).Error
}

func (r *JobRepository) Delete(id uuid.UUID) error {
	return r.DB.Delete(&models.Job{}, id).Error
}

func (r *JobRepository) GetCompanyIDByRecruiterID(id uuid.UUID) (uuid.UUID, error) {
	var companyID uuid.UUID
	err := r.DB.Table("jobs").Select("company_id").Where("recruiter_id = ?", id).First(&companyID).Error
	if err != nil {
		return uuid.Nil, err
	}
	return companyID, nil
}

func (r *JobRepository) GetRecruiterIDByUserID(id uuid.UUID) (uuid.UUID, error) {
	var recruiterID uuid.UUID
	err := r.DB.Table("recruiters").Select("id").Where("user_id = ?", id).First(&recruiterID).Error
	if err != nil {
		return uuid.Nil, err
	}
	return recruiterID, nil
}

func (r *JobRepository) GetCandidateIDByUserID(id uuid.UUID) (uuid.UUID, error) {
	var candidateID uuid.UUID
	err := r.DB.Table("candidates").Select("id").Where("job_id = ?", id).First(&candidateID).Error
	if err != nil {
		return uuid.Nil, err
	}
	return candidateID, nil
}

func (r *JobRepository) GetCompanyNameAndLocationByID(id uuid.UUID) (string, string, error) {
	var companyName, location string
	err := r.DB.Table("companies").Select("name, location").Where("id = ?", id).First(&companyName, &location).Error
	if err != nil {
		return "", "", err
	}
	return companyName, location, nil
}
