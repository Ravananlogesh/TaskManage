package login

import (
	"errors"
	"net/http"
	"os"
	"tasks/internal/models"
	"tasks/internal/utils"
	database "tasks/migrations"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Login(c *gin.Context) {
	log := new(utils.Logger)
	log.SetSid(c.Request)
	var login models.Login

	if err := c.ShouldBindJSON(&login); err != nil {
		log.Log(utils.ERROR, "Login 01", err.Error())
		utils.JSONErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	tokenString, err := CheckAndCompare(log, login)
	if err != nil {
		log.Log(utils.ERROR, "Login 02", err.Error())
		utils.JSONErrorResponse(c, http.StatusUnauthorized, err)
		return
	}

	log.Log(utils.INFO, "Login successful for user: "+login.UserName)
	utils.JSONResponse(c, http.StatusOK, true, "Login Successful", gin.H{"token": tokenString})
}

func CheckAndCompare(log *utils.Logger, login models.Login) (string, error) {
	var user models.User

	if err := database.GDB.Where("username = ?", login.UserName).First(&user).Error; err != nil {
		log.Log(utils.ERROR, "CAC001", "User not found: "+login.UserName)
		return "", err
	}

	if !user.CheckPassword(log, login.Password) {
		log.Log(utils.ERROR, "CAC002", "Password mismatch for user: "+login.UserName)
		return "", errors.New("Password is  incorrect")
	}

	mapClaims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		log.Log(utils.ERROR, "CAC003", "Error signing token: "+err.Error())
		return "", err
	}
	return tokenString, nil
}
