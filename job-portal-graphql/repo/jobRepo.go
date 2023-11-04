package repo

import (
	"errors"
	"fmt"
	"graphql/models"
	"github.com/rs/zerolog/log"
)

func (r *Repo) CreateJob(newUser models.Job) (models.Job, error) {
	err := r.Db.Create(&newUser).Error
	if err != nil {
		log.Info().Err(err).Send()
		return models.Job{}, fmt.Errorf("error in creating job %w", err)
	}
	return newUser, nil
}

func (r *Repo) ViewAllJobs() ([]models.Job, error) {
	var jobDetails []models.Job
	result := r.Db.Find(&jobDetails)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return []models.Job{}, errors.New("jobs data not found")
	}
	return jobDetails, nil
}

func (r *Repo) ViewJobById(jobId string) (models.Job, error) {
	var jobDetails models.Job
	result := r.Db.Where("id=?", jobId).Find(&jobDetails)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return models.Job{}, errors.New("error in finding job records")
	}
	return jobDetails, nil
}

func (r *Repo) ViewJobsForCompany(companyId string) ([]models.Job, error) {
	var jobDetails []models.Job

	result := r.Db.Where("company_id=?", companyId).Find(&jobDetails)

	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return nil, errors.New("job not found")
	}
	return jobDetails, nil

}
