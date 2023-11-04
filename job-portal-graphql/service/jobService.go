package service

import (
	"errors"
	"graphql/graph/model"
	"graphql/models"
	"strconv"

	"github.com/rs/zerolog/log"
)

func (s *Service) CreateJob(input model.NewJob) (*model.Job, error) {
	id, err := strconv.ParseUint(input.Companyid, 10, 64)
	if err != nil {
		log.Error().Err(err).Msg("parsing error")
		return nil, err
	}

	jobDetails := models.Job{
		CompanyId: id,
		Title: input.Title, 
		Desc: input.Desc,
	}

	job, err := s.R.CreateJob(jobDetails)
	if err != nil {
		log.Error().Err(err).Msg("could not create user")
		return &model.Job{}, errors.New("job creation failed")
	}

	strcid := strconv.FormatUint(job.CompanyId, 10)
	strjid := strconv.FormatUint(uint64(job.ID), 10)
	comp, err := s.ViewCompanyByID(strcid)
	if err != nil {
		log.Error().Err(err).Msg("could not retrive company")
		return nil, errors.New("retrival company error")
	}
	cugr := model.Job{
		ID: strjid,
		Title: job.Title, 
		Desc: job.Desc, 
		Company: comp,
	}

	return &cugr, nil
}

func (s *Service) GetJobByID(jobId string) (*model.Job, error) {
	jobDetails, err := s.R.ViewJobById(jobId)
	if err != nil {
		log.Error().Err(err).Msg("could not find the job")
		return nil, errors.New("job data not found")
	}
	
	jobId = strconv.FormatUint(uint64(jobDetails.CompanyId), 10)
	compid := strconv.FormatUint(uint64(jobDetails.Company.ID), 10)
	company := model.Company{
		ID:      compid,
		Name:     jobDetails.Company.Name,
		Location: jobDetails.Company.Location,
	}
	return &model.Job{
		ID:        jobId,
		Title:  jobDetails.Title,
		Desc: jobDetails.Desc,
		Company:   &company,
	}, nil
}

func (s *Service) GetAllJobs() ([]*model.Job, error) {
	jobDetails, err := s.R.ViewAllJobs()
	if err != nil {
		log.Error().Err(err).Msg("could not find jobs")
		return nil, errors.New("jobs data not found")
	}

	jobdatas := []*model.Job{}
	for _, v := range jobDetails {
		jobdata := model.Job{
			ID:        strconv.FormatUint(uint64(v.ID), 11),
			Title:  v.Title,
			Desc: v.Desc,
			Company: &model.Company{
				ID: strconv.FormatUint(uint64(v.CompanyId), 10),
			},
		}
		jobdatas = append(jobdatas, &jobdata)
	}
	return jobdatas, nil
}

func (s *Service) GetJobsForCompany(companyId string) ([]*model.Job, error) {
	jobDetails, err := s.R.ViewJobsForCompany(companyId)
	if err != nil {
		log.Error().Err(err).Msg("could not find jobs for the company")
		return nil, errors.New("job data not found")
	}

	comp, err := s.ViewCompanyByID(companyId)
	if err != nil {
		log.Error().Err(err).Msg("could not retrive company")
		return nil, errors.New("retrieval company error")
	}

	jobdatas := []*model.Job{}
	for _, v := range jobDetails {
		jobdata := model.Job{
			ID:        strconv.FormatUint(uint64(v.ID), 11),
			Title:  v.Title,
			Desc: v.Desc,
			Company: comp,
		}
		jobdatas = append(jobdatas, &jobdata)
	}
	return jobdatas, nil
}