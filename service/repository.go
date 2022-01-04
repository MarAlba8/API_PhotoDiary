package service

import (
	"PhotoDiary/models"
)

type Repository interface {
	Insert(account models.Credentials) (err error)
	Update(identifier string, account models.CredentialsToUpdate) (err error)
	GetAccount(id string) (account models.Account, err error)
	GetAll() (accounts []models.Account, err error)
}
