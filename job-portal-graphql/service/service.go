package service

import (
	"errors"
	"graphql/graph/model"
	"graphql/repo"
)

type UserService interface {
	SignUp(input model.NewUser) (*model.User, error)
	CreateCompany(input model.NewCompany) (*model.Company, error)
	ViewAllCompanies() ([]*model.Company, error)
	ViewCompanyByID(cid string) (*model.Company, error)
	CreateJob(input model.NewJob) (*model.Job, error)
	GetJobsForCompany(cid string) ([]*model.Job, error)
	GetAllJobs() ([]*model.Job, error)
	GetJobByID(jid string) (*model.Job, error) 
}
type Service struct {
	R repo.UserRepo
}

func NewService(r repo.UserRepo) (UserService, error) {
	if r == nil {
		return nil, errors.New("the service layer is nil")

	}
	return &Service{
		R: r,
	}, nil
}
