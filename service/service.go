package service

import (
	"PhotoDiary/models"
	"errors"
)

type Service struct {
	Repository Repository
}

func LoadService(rep Repository) *Service {
	return &Service{
		Repository: rep,
	}
}

func (serv *Service) Register(account models.Account) (err error) {
	err = serv.Repository.Insert(account)
	if err != nil {
		return err
	}
	return nil
}

func (serv *Service) Login(account models.Account) (err error) {
	savedAccount, err := serv.Repository.GetAccountByNickname(account.Nickname)
	if err != nil {
		return errors.New("wrong Nickname")
	}
	if savedAccount.Password == account.Password {
		return nil
	}
	return errors.New("wrong Password")
}

func (serv *Service) Update(account models.Account) (err error) {
	err = serv.Repository.Update(account)
	if err != nil {
		return errors.New("wrong data")
	}
	return nil
}

func (serv *Service) GetAll() (accounts []models.Account, err error) {
	accounts, err = serv.Repository.GetAll()
	if err != nil {
		return nil, errors.New("error getting data")
	}
	return accounts, nil
}

func (serv *Service) GetAccount(id string) (account models.Account, err error) {
	account, err = serv.Repository.GetAccount(id)
	if err != nil {
		return account, errors.New("error getting account")
	}
	return account, nil
}
