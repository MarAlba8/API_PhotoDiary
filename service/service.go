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

func (serv *Service) Register(userData models.Credentials) (err error) {

	if userData.Nickname == "" || userData.Email == "" || userData.Username == "" || userData.Password == "" {
		return errors.New("Wrong data")
	}

	userRegistered, err := serv.IsUserRegistered(userData)
	if userRegistered {
		return err
	}

	err = serv.Repository.Insert(userData)
	if err != nil {
		return err
	}
	return nil
}

func (serv *Service) Login(account models.Account) (err error) {
	savedAccount, err := serv.Repository.GetAccount(account.Nickname)
	if err != nil {
		return errors.New("wrong Nickname")
	}
	if savedAccount.Password == account.Password {
		return nil
	}
	return errors.New("wrong Password")
}

func (serv *Service) Update(account models.UpdateCredentials) (err error) {
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

func (serv *Service) IsUserRegistered(identifier models.Credentials) (rsp bool, err error) {

	account, _ := serv.Repository.GetAccount(identifier.Email)
	if account.ID != "" {
		return true, errors.New("This email has an account already")
	}

	account, _ = serv.Repository.GetAccount(identifier.Nickname)
	if account.ID != "" {
		return true, errors.New("Nickname already taken")
	}

	return false, nil
}
