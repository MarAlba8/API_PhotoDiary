package service

import (
	"PhotoDiary/models"
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
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

	hashPassword, err := GenerateHashPassword(userData.Password)
	if err != nil {
		return err
	}

	userData.Password = hashPassword
	err = serv.Repository.Insert(userData)
	if err != nil {
		return err
	}
	return nil
}

func (serv *Service) Login(inputCredentials models.LoginCredentials) (string, error) {
	savedAccount, err := serv.Repository.GetAccount(inputCredentials.Identifier)
	if err != nil {
		return "", errors.New("wrong nickname or email")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(savedAccount.Password),
		[]byte(inputCredentials.Password))
	if err != nil {
		return "", errors.New("wrong Password")
	}

	token, err := GenerateToken(inputCredentials.Identifier)
	if err != nil {
		return "", errors.New("failed to create token with error: " + err.Error())
	}
	return token, nil
}

func (serv *Service) Update(identifier string, account models.CredentialsToUpdate) (err error) {
	hashPassword, err := GenerateHashPassword(account.Password)
	if err != nil {
		return err
	}
	account.Password = hashPassword
	err = serv.Repository.Update(identifier, account)
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

func (serv *Service) GetAccount(identifier string) (account models.Account, err error) {
	account, err = serv.Repository.GetAccount(identifier)
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

func GenerateHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func GenerateToken(identifier string) (string, error) {
	token := jwt.New(jwt.GetSigningMethod("HS256"))

	now := time.Now().Local()
	token.Claims = jwt.MapClaims{
		"authorized": true,
		"identifier": identifier,
		"exp":        now.Add(time.Hour * time.Duration(1)).Unix(),
	}
	//var jwtKey = []byte("my_secret_key")
	tokenString, err := token.SignedString([]byte("my_secret_key"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
