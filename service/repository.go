package service

import (
	"PhotoDiary/models"
)

type Repository interface {
	Insert(account models.Account) (err error)
	Update(account models.Account) (err error)
	GetAccount(id string) (account models.Account, err error)
	GetAccountByNickname(nickname string) (account models.Account, err error)
	GetAll() (accounts []models.Account, err error)
}
