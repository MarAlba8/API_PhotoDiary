package api

import (
	"PhotoDiary/models"
	"PhotoDiary/service"
	"fmt"
	"github.com/dgrijalva/jwt-go"
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
	router.GET("/", TokenAuthMiddleware(), h.MainHandler)
	router.GET("/profile", TokenAuthMiddleware(), h.ProfileHandler)
	router.POST("/login", h.LoginHandler)
	router.POST("/register", h.RegisterHandler)
	router.PUT("/updateaccount", TokenAuthMiddleware(), h.UpdateAccountHandler)

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
	identifier := c.GetString("identifier")
	account, err := h.Usecase.GetAccount(identifier)
	if err != nil {
		c.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "Error getting data"})
		return
	}
	c.IndentedJSON(http.StatusOK, account)
}

func (h *GinHandler) LoginHandler(c *gin.Context) {
	var account models.LoginCredentials
	if err := c.BindJSON(&account); err != nil {
		c.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"message": "Incorrect credentials"})
		return
	}
	token, err := h.Usecase.Login(account)
	if err != nil {
		c.IndentedJSON(
			http.StatusInternalServerError, err.Error())
		return
	}
	c.IndentedJSON(
		http.StatusOK, token)
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
	identifier := c.GetString("identifier")
	var newData models.CredentialsToUpdate
	if err := c.BindJSON(&newData); err != nil {
		c.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"message": "Wrong data"})
		return
	}

	err := h.Usecase.Update(identifier, newData)
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

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var jwtKey = []byte("my_secret_key")
		const prefix = "Bearer "

		header := c.GetHeader("Authorization")
		if header == "" {
			c.AbortWithStatusJSON(400, gin.H{"error": "Authorization header missing or empty"})
		}

		token := header[len(prefix):]
		tkn, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.IndentedJSON(
					http.StatusUnauthorized,
					gin.H{"message": "Signature invalid"})
				return
			}
			c.AbortWithStatusJSON(401, gin.H{"error": fmt.Sprintf("unable to verify token with error: %v", err)})
		}

		if claims, ok := tkn.Claims.(jwt.MapClaims); ok && tkn.Valid {
			fmt.Println(claims["identifier"])
			fmt.Println(claims)
			c.Set("identifier", claims["identifier"])
		} else {
			fmt.Println(err)
		}

		c.Next()
	}
}
