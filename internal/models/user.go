package models

import (
	"tasks/internal/utils"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           uint      `gorm:"primaryKey"`
	Username     string    `gorm:"unique;not null"`
	PasswordHash string    `gorm:"not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
}

// HashPassword hashes the password before saving
func (user *User) PasswordHASH(log *utils.Logger, password string) error {
	log.Log(utils.INFO, "PasswordHASH started...")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Log(utils.ERROR, "HPD001", err.Error())
		return err
	}
	user.PasswordHash = string(hashedPassword)
	log.Log(utils.INFO, "PasswordHASH ended...")
	return nil
}

// CheckPassword verifies if the given password is correct
func (user *User) CheckPassword(log *utils.Logger, password string) bool {
	log.Log(utils.INFO, "checkpassword +")
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	log.Log(utils.INFO, "checkpassword -")
	return err == nil
}
