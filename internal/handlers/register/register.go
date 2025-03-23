package register

import (
	"net/http"
	"tasks/internal/models"
	"tasks/internal/repo"
	"tasks/internal/utils"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	log := new(utils.Logger)
	log.SetSid(c.Request)

	log.Log(utils.INFO, "Register start")
	var register models.Login
	err := c.ShouldBindBodyWithJSON(&register)
	if err != nil {
		log.Log(utils.ERROR, "Register 001", err.Error())
		utils.JSONErrorResponse(c, http.StatusBadRequest, err)
		return
	}
	userData := models.User{
		Username: register.UserName,
	}
	err = userData.PasswordHASH(log, register.Password)
	if err != nil {
		log.Log(utils.ERROR, "Register 002", err.Error())
		utils.JSONErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	userRepo := repo.NewUserRepo()
	err = userRepo.CreateUser(&userData)
	if err != nil {
		log.Log(utils.ERROR, "Register 002", err.Error())
		utils.JSONErrorResponse(c, http.StatusInternalServerError, err)
		return
	}
	log.Log(utils.INFO, "Register end")
	utils.JSONResponse(c, http.StatusOK, true, "Registeer Successful", gin.H{"sucess": "Registeer Successful"})

}
