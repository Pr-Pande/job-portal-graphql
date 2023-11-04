package repo

import (
	"graphql/models"
	"github.com/rs/zerolog/log"
)

func (r *Repo) CreateUser(newUser models.User) (models.User, error) {

	err := r.Db.Create(&newUser).Error
	if err != nil {
		log.Info().Err(err).Send()
		return models.User{}, err
	}
	return newUser, nil

}
