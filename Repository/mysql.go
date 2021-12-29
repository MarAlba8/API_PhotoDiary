package Repository

import (
	"PhotoDiary/models"
	"database/sql"
	"errors"
)

//TODO: Check models and parameters

type MysqlDB struct {
	DB *sql.DB
}

func LoadDB(db *sql.DB) *MysqlDB {
	return &MysqlDB{
		DB: db,
	}
}

func (s *MysqlDB) Insert(account models.Account) (err error) {
	_, err = s.DB.Exec("INSERT INTO account (nickname, password, profilePicture, username, email) VALUES (?,?,?,?,?)",
		account.Nickname,
		account.Password,
		account.ProfilePicture,
		account.Username,
		account.Email)
	if err != nil {
		println(err.Error())
		return errors.New("error inserting data")
	}
	return nil
}

func (s *MysqlDB) Update(account models.Account) (err error) {
	_, err = s.DB.Exec("UPDATE account SET account.username=?, account.password=?, account.profilePicture=? WHERE account.id =? ",
		account.Username,
		account.Password,
		account.ProfilePicture,
		account.ID)
	if err != nil {
		return errors.New("Error updating data")
	}
	return nil
}

func (s *MysqlDB) GetAccount(id string) (account models.Account, err error) {
	row := s.DB.QueryRow("SELECT * FROM account WHERE account.id = ?",
		id).Scan(
		&account.ID,
		&account.Username,
		&account.Password,
		&account.ProfilePicture,
		&account.Nickname,
		&account.Email,
	)
	if row != nil {
		return account, errors.New("Error getting data")
	}
	return account, nil
}

func (s *MysqlDB) GetAccountByNickname(nickname string) (currentAccount models.Account, err error) {
	row := s.DB.QueryRow("SELECT * FROM account WHERE account.nickname = ?",
		nickname).Scan(
		&currentAccount.ID,
		&currentAccount.Username,
		&currentAccount.Password,
		&currentAccount.ProfilePicture,
		&currentAccount.Nickname,
		&currentAccount.Email,
	)
	if row != nil {
		return currentAccount, errors.New("Error getting data")
	}
	return currentAccount, nil
}

func (s *MysqlDB) GetAll() (accounts []models.Account, err error) {
	rows, err := s.DB.Query("SELECT * FROM account;")
	if err != nil {
		return nil, errors.New("Error getting data")
	}

	var currentAccount models.Account
	for rows.Next() {
		err := rows.Scan(
			&currentAccount.ID,
			&currentAccount.Nickname,
			&currentAccount.Password,
			&currentAccount.ProfilePicture,
			&currentAccount.Username,
			&currentAccount.Email,
		)
		if err != nil {
			return nil, errors.New("Error getting data")
		}
		accounts = append(accounts, currentAccount)
	}
	return accounts, nil
}
