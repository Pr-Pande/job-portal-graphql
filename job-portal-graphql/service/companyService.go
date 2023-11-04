package service

import (
	"errors"
	"graphql/graph/model"
	"graphql/models"
	"strconv"
	"github.com/rs/zerolog/log"
)

func (s *Service) CreateCompany(input model.NewCompany) (*model.Company, error) {
	companyDetails := models.Company{
		Name:     input.Name,
		Location: input.Location,
	}
	companyDetails, err := s.R.CreateCompany(companyDetails)
	if err != nil {
		log.Error().Err(err).Msg("could not create company")
		return nil, errors.New("company creation failed")
	}

	cid := strconv.FormatUint(uint64(companyDetails.ID), 10)
	return &model.Company{
		ID:       cid,
		Name:     companyDetails.Name,
		Location: companyDetails.Location,
	}, nil

}

func (s *Service) ViewAllCompanies() ([]*model.Company, error) {
	companyDetails, err := s.R.ViewAllCompany()
	if err != nil {
		log.Error().Err(err).Msg("could not find any company")
		return nil, errors.New("data of companies not found")
	}
	cDatas := []*model.Company{}

	for _, v := range companyDetails {
		cData := model.Company{
			ID:       strconv.FormatUint(uint64(v.ID), 10),
			Name:     v.Name,
			Location: v.Location,
		}
		cDatas = append(cDatas, &cData)
	}
	return cDatas, nil

}

func (s *Service) ViewCompanyByID(companyId string) (*model.Company, error) {

	companyDetails, err := s.R.FetchCompanyById(companyId)
	if err != nil {
		log.Error().Err(err).Msg("could not find the company")
		return nil, errors.New("company data not found")
	}
	return &model.Company{
		ID:       companyId,
		Name:     companyDetails.Name,
		Location: companyDetails.Location,
	}, nil

}
