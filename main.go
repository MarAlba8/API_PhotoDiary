package main

import (
	"PhotoDiary/driver"
	"PhotoDiary/models"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

var DB = driver.InitDatabase()

func MainHandler(c *gin.Context) {
	var currentAccount models.Account
	var accounts []models.Account

	rows, err := DB.Query("SELECT * FROM account;")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error getting data"})
		return
	}

	for rows.Next() {
		err := rows.Scan(&currentAccount.ID, &currentAccount.Nickname, &currentAccount.Password, &currentAccount.ProfilePicture, &currentAccount.Username, &currentAccount.Email)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error getting data"})
			return
		}
		accounts = append(accounts, currentAccount)
	}
	c.IndentedJSON(http.StatusOK, accounts)

}

func LoginHandler(c *gin.Context) {
	var receiveData models.Account
	var savedData models.Account

	if err := c.BindJSON(&receiveData); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error reading data"})
		return
	}

	err := DB.QueryRow("SELECT * FROM account WHERE account.nickname = ?", receiveData.Nickname).Scan(&savedData.ID, &savedData.Nickname, &savedData.Password, &savedData.ProfilePicture, &savedData.Username, &savedData.Email)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "User not found"})
		return
	}

	if savedData.Password == receiveData.Password {
		c.IndentedJSON(http.StatusOK, savedData)
		return
	}
	c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "User not found"})
}

func RegisterHandler(c *gin.Context) {
	var newAccount models.Account
	if err := c.BindJSON(&newAccount); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error reading data"})
		return
	}

	err := DB.QueryRow("INSERT INTO account (nickname, password, profilePicture, username, email) VALUES (?,?,?,?,?)", newAccount.Nickname, newAccount.Password, newAccount.ProfilePicture, newAccount.Username, newAccount.Email)
	if err.Err() != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error saving data"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "User registered"})
}

func ProfileHandler(c *gin.Context) {
	id := c.Param("id")
	var currentAccount models.Account

	err := DB.QueryRow("SELECT * FROM account WHERE account.id = ?", id).Scan(&currentAccount.ID, &currentAccount.Username, &currentAccount.Password, &currentAccount.ProfilePicture, &currentAccount.Nickname, &currentAccount.Email)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "User data not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, currentAccount)

}

func UpdateAccountHandler(c *gin.Context) {
	var newData models.Account
	if err := c.BindJSON(&newData); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Wrong data"})
		return
	}

	err := DB.QueryRow("UPDATE account SET account.username=?, account.password=?, account.profilePicture=? WHERE account.id =? ", newData.Username, newData.Password, newData.ProfilePicture, newData.ID)
	if err.Err() != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error while saving data"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Account updated"})
}

func main() {
	defer DB.Close()

	router := gin.Default()
	router.GET("/", MainHandler)
	router.GET("/profile/:id", ProfileHandler)
	router.POST("/login", LoginHandler)
	router.POST("/register", RegisterHandler)
	router.PUT("/updateaccount", UpdateAccountHandler)

	log.Fatal(router.Run("localhost:8080"))
}
