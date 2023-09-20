package entity

import (
	"github.com/ArthurRF/runners-api/pkg/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       entity.ID `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"-"` // password nunca será exibido para o usuário final
}

func NewUser(name, email, password string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		Name:     name,
		Email:    email,
		Password: string(hash),
	}, nil
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
