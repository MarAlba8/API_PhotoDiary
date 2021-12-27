package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

const (
	username = "root"
	password = ""
	hostname = "127.0.0.1:3306"
	dbname   = "db"
)

var db = initDb()

func dsn(dbName string) string {
	//username:password@protocol(address)/dbname?param=value
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

func initDb() *sql.DB {
	dsn := dsn("db")
	db, err := sql.Open("mysql", dsn)
	if err == nil {
		return db
	}
	return nil
}

type account struct {
	ID             string `json:"id"`
	Nickname       string `json:"nickname"`
	Password       string `json:"password"`
	ProfilePicture string `json:"profilepicture"`
	idUser         string `json:"iduser"`
}

var accounts = []account{
	{
		ID:             "1",
		Nickname:       "Andre1",
		Password:       "****",
		ProfilePicture: "www.picture.com",
	},
	{
		ID:             "2",
		Nickname:       "jose1",
		Password:       "****",
		ProfilePicture: "www.picture2.com",
	},
}

var count = len(accounts)

func mainHandler(c *gin.Context) {
	var currentAccount account
	var accounts []account

	rows, err := db.Query("SELECT * FROM account;")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error getting data"})
		return
	}

	for rows.Next() {
		err := rows.Scan(&currentAccount.ID, &currentAccount.Nickname, &currentAccount.Password, &currentAccount.ProfilePicture, &currentAccount.idUser)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error getting data"})
			return
		}
		accounts = append(accounts, currentAccount)
	}
	c.IndentedJSON(http.StatusOK, accounts)

}

func loginHandler(c *gin.Context) {
	values := c.Request.URL.Query()
	nickname := values["nickname"][0]
	password := values["password"][0]

	for _, data := range accounts {
		if data.Nickname == nickname && (data.Password) == password {
			c.IndentedJSON(http.StatusOK, data)
			return
		}
	}

	c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "User not found"})
}

func registerHandler(c *gin.Context) {
	var newAccount account
	if err := c.BindJSON(&newAccount); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error reading data"})
		return
	}
	count += 1
	newAccount.ID = fmt.Sprint(count)
	accounts = append(accounts, newAccount)
	c.IndentedJSON(http.StatusAccepted, "User registered")
}

func profileHandler(c *gin.Context) {
	id := c.Param("id")

	for _, data := range accounts {
		if data.ID == id {
			c.IndentedJSON(http.StatusOK, data)
			return
		}
	}
	c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "User data not found"})
}

func updateAccountHandler(c *gin.Context) {
	var newData account
	if err := c.BindJSON(&newData); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Wrong data"})
		return
	}

	for i, account := range accounts {
		if newData.ID == account.ID {
			accounts[i] = newData
			c.IndentedJSON(http.StatusOK, newData)
			return
		}
	}
}

func main() {
	defer db.Close()

	router := gin.Default()
	router.GET("/", mainHandler)
	router.GET("/profile/:id", profileHandler)
	router.POST("/login", loginHandler)
	router.POST("/register", registerHandler)
	router.PUT("/updateAccount", updateAccountHandler)

	router.Run("localhost:8080")
}
