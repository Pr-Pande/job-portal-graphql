package repo

import (
	"errors"
	"graphql/models"

	"gorm.io/gorm"
)

type Repo struct {
	Db *gorm.DB
}
type UserRepo interface {
	CreateUser(newUser models.User) (models.User, error)
	CreateCompany(newCompany models.Company) (models.Company, error)
	ViewAllCompany() ([]models.Company, error)
	FetchCompanyById(companyId string) (models.Company, error)
	CreateJob(newJob models.Job) (models.Job, error)
	ViewJobsForCompany(companyId string) ([]models.Job, error)
	ViewAllJobs() ([]models.Job, error)
	ViewJobById(jobId string) (models.Job, error)
}

func NewRepo(db *gorm.DB) (UserRepo, error) {
	if db == nil {
		return nil, errors.New("db cannot be nil")

	}
	return &Repo{
		Db: db,
	}, nil
}
