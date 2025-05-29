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
func (r *JobRepository) FindAll() ([]models.Job, error) {
	var jobs []models.Job

	if err := r.DB.Find(&jobs).Error; err != nil {
		return nil, err
	}

	return jobs, nil
}

func (r *JobRepository) FindAllByRecruiterID(id uuid.UUID) ([]models.Job, error) {
	var jobs []models.Job

	if err := r.DB.Where("recruiter_id = ?", id).Find(&jobs).Error; err != nil {
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
	var recruiter models.Recruiter

	err := r.DB.Where("id = ?", id).First(&recruiter).Error
	if err != nil {
		return uuid.Nil, err
	}
	return *recruiter.CompanyID, nil
}

func (r *JobRepository) GetRecruiterIDByUserID(id uuid.UUID) (uuid.UUID, error) {
	var recruiter models.Recruiter

	err := r.DB.Where("user_id = ?", id).First(&recruiter).Error
	if err != nil {
		return uuid.Nil, err
	}

	return recruiter.ID, nil
}

func (r *JobRepository) GetCandidateIDByUserID(id uuid.UUID) (uuid.UUID, error) {
	var candidate models.Candidate

	err := r.DB.Where("user_id = ?", id).First(&candidate).Error
	if err != nil {
		return uuid.Nil, err
	}

	return candidate.ID, nil
}

func (r *JobRepository) GetCompanyNameAndLocationByID(id uuid.UUID) (string, string, string, error) {
	type CompanyInfo struct {
		Name     string
		Location string
		Logo     string
	}

	var info CompanyInfo
	err := r.DB.Table("companies").Select("name, location, logo").Where("id = ?", id).First(&info).Error
	if err != nil {
		return "", "", "", err
	}
	return info.Name, info.Location, info.Logo, nil
}
