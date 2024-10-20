package models


import (
	"golang.org/x/crypto/bcrypt"
)



type User struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Email string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

func (u *User) HashPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    u.Password = string(hashedPassword)
    return nil
}

func (u *User) CheckPassword(providedPassword string) error {
    return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(providedPassword))
}