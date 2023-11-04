package repo

import (
	"errors"
	"graphql/models"

	"github.com/rs/zerolog/log"
)

func (r *Repo) CreateCompany(newCompany models.Company) (models.Company, error) {

	err := r.Db.Create(&newCompany).Error
	if err != nil {
		log.Info().Err(err).Send()
		return models.Company{}, err
	}
	return newCompany, nil

}

func (r *Repo) ViewAllCompany() ([]models.Company, error) {
	var companyDetails []models.Company

	result := r.Db.Find(&companyDetails)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return []models.Company{}, errors.New("could not found companies in record")
	}
	return companyDetails, nil
}

func (r *Repo) FetchCompanyById(compId string) (models.Company, error) {
	var companyData models.Company
	result := r.Db.Where("id=?", compId).First(&companyData)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return models.Company{}, errors.New("company not found")
	}
	return companyData, nil

}
