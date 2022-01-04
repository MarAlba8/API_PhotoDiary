package service

import (
	"PhotoDiary/models"
)

type UseCase interface {
	Register(account models.Credentials) (err error)
	Login(account models.LoginCredentials) (string, error)
	Update(identifier string, account models.CredentialsToUpdate) (err error)
	GetAll() (accounts []models.Account, err error)
	GetAccount(id string) (account models.Account, err error)
	IsUserRegistered(identifier models.Credentials) (rsp bool, err error)
}
