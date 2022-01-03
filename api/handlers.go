package api

import (
	"PhotoDiary/models"
	"PhotoDiary/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GinHandler struct {
	Usecase service.UseCase
}

func NewGinHandler(usecase service.UseCase) *gin.Engine {
	h := &GinHandler{
		Usecase: usecase,
	}

	router := gin.Default()
	router.GET("/", h.MainHandler)
	router.GET("/profile/:id", h.ProfileHandler)
	router.POST("/login", h.LoginHandler)
	router.POST("/register", h.RegisterHandler)
	router.PUT("/updateaccount", h.UpdateAccountHandler)

	return router
}

func (h *GinHandler) MainHandler(c *gin.Context) {
	accounts, err := h.Usecase.GetAll()
	if err != nil {
		c.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "Error getting data"})
	}
	c.IndentedJSON(
		http.StatusOK,
		accounts)
}

func (h *GinHandler) ProfileHandler(c *gin.Context) {
	id := c.Param("id")
	account, err := h.Usecase.GetAccount(id)
	if err != nil {
		c.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "Error getting data"})
	}
	c.IndentedJSON(http.StatusOK, account)
}

func (h *GinHandler) LoginHandler(c *gin.Context) {
	var account models.Account
	if err := c.BindJSON(&account); err != nil {
		c.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"message": "Incorrect credentials"})
		return
	}
	err := h.Usecase.Login(account)
	if err != nil {
		c.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "Wrong password or nickname"})
		return
	}
	c.IndentedJSON(
		http.StatusOK,
		gin.H{"message": "User logged successfully"})
}

func (h *GinHandler) RegisterHandler(c *gin.Context) {
	var userData models.Credentials

	if err := c.BindJSON(&userData); err != nil {
		c.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"message": "Error reading data"})
		return
	}

	err := h.Usecase.Register(userData)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	c.IndentedJSON(
		http.StatusOK,
		gin.H{"message": "User registered successfully"})
}

func (h *GinHandler) UpdateAccountHandler(c *gin.Context) {
	var newData models.UpdateCredentials
	if err := c.BindJSON(&newData); err != nil {
		c.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"message": "Wrong data"})
		return
	}

	err := h.Usecase.Update(newData)
	if err != nil {
		c.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "Error while saving data"})
		return
	}
	c.IndentedJSON(
		http.StatusOK,
		gin.H{"message": "Account updated"})
}
