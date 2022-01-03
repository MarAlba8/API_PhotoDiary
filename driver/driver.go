package driver

import (
	"database/sql"
	"fmt"
)

const (
	username = "root"
	password = ""
	hostname = "127.0.0.1:3306"
)

func dsn(dbName string) string {
	//username:password@protocol(address)/dbname?param=value
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

func InitDatabase() *sql.DB {
	dsn := dsn("db-pd")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Error connecting db")
		return nil
	}
	return db
}

func CloseDatabase(connection *sql.DB) {
	connection.Close()
}
