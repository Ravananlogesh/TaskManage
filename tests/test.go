package main

import (
	"fmt"
	"tasks/internal/models"
	"tasks/internal/utils"
)

func main() {

	login := models.User{
		Username:     "logesh",
		PasswordHash: "loge@123",
	}
	log := new(utils.Logger)
	login.PasswordHASH(log, login.PasswordHash)
	fmt.Println(login.PasswordHash)
}
