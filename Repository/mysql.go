package Repository

import (
	"PhotoDiary/models"
	"database/sql"
)

type MysqlDB struct {
	DB *sql.DB
}

func LoadDB(db *sql.DB) *MysqlDB {
	return &MysqlDB{
		DB: db,
	}
}

func (s *MysqlDB) Insert(account models.Credentials) (err error) {
	_, err = s.DB.Exec("INSERT INTO account (nickname, password, username, email) VALUES (?,?,?,?)",
		account.Nickname,
		account.Password,
		account.Username,
		account.Email)
	if err != nil {
		return err
	}
	return nil
}

func (s *MysqlDB) Update(account models.UpdateCredentials) (err error) {
	_, err = s.DB.Exec("UPDATE account SET account.username=?, "+
		"account.password=?, account.profilePicture=? WHERE account.id =? ",
		account.Username,
		account.Password,
		account.ProfilePicture,
		account.ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *MysqlDB) GetAccount(identifier string) (account models.Account, err error) {
	var profilePicture sql.NullString
	err = s.DB.QueryRow("SELECT * FROM account WHERE "+
		"account.id = ? OR account.nickname = ? OR account.email = ?",
		identifier, identifier, identifier).Scan(
		&account.ID,
		&account.Username,
		&account.Password,
		&profilePicture,
		&account.Nickname,
		&account.Email,
	)
	if err != nil {
		return account, err
	}
	account.ProfilePicture = profilePicture.String
	return account, nil
}

/*func (s *MysqlDB) GetAccountByNickname(nickname string) (currentAccount models.Account, err error) {
	err = s.DB.QueryRow("SELECT * FROM account WHERE account.nickname = ?",
		nickname).Scan(
		&currentAccount.ID,
		&currentAccount.Username,
		&currentAccount.Password,
		&currentAccount.ProfilePicture,
		&currentAccount.Nickname,
		&currentAccount.Email,
	)
	if err != nil {
		return currentAccount, err
	}
	return currentAccount, nil
}*/

func (s *MysqlDB) GetAll() (accounts []models.Account, err error) {
	rows, err := s.DB.Query("SELECT * FROM account;")
	if err != nil {
		return nil, err
	}
	var profilePicture sql.NullString
	var currentAccount models.Account
	for rows.Next() {
		err := rows.Scan(
			&currentAccount.ID,
			&currentAccount.Nickname,
			&currentAccount.Password,
			&profilePicture,
			&currentAccount.Username,
			&currentAccount.Email,
		)
		if err != nil {
			return nil, err
		}
		currentAccount.ProfilePicture = profilePicture.String
		accounts = append(accounts, currentAccount)
	}
	return accounts, nil
}

/*func (s *MysqlDB) IsUserRegistered(identifier string) (rsp bool, err error) {
	var account models.Account
	err = s.DB.QueryRow("SELECT * FROM account WHERE "+
		"account.nickname = ? OR account.email = ?",
		identifier, identifier).Scan(
		&account.ID,
		&account.Username,
		&account.Password,
		&account.ProfilePicture,
		&account.Nickname,
		&account.Email,
	)
	if err != nil {
		return false, err
	}
	if account.ID != "" {
		return true, nil
	}
	return false, nil
}*/
