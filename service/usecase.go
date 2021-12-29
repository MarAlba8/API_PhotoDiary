package service

import (
	"PhotoDiary/models"
)

type UseCase interface {
	Register(account models.Account) (err error)
	Login(account models.Account) (err error)
	Update(account models.Account) (err error)
	GetAll() (accounts []models.Account, err error)
	GetAccount(id string) (account models.Account, err error)
}
