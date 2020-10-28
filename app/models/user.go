package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User model
type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"-"`
}

// ComparePasswords compares passwords
func (u *User) ComparePasswords(plaintext string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plaintext))
}

// BeforeCreate hook
func (u *User) BeforeCreate(tx *gorm.DB) error {
	return u.bcryptPasswords(tx)
}

// BeforeUpdate hook
func (u *User) BeforeUpdate(tx *gorm.DB) error {
	if tx.Statement.Changed("Password") {
		return u.bcryptPasswords(tx)
	}
	return nil
}

func (u *User) bcryptPasswords(tx *gorm.DB) error {
	var newPass string
	switch u := tx.Statement.Dest.(type) {
	case map[string]interface{}:
		newPass = u["password"].(string)
	case *User:
		newPass = u.Password
	case []*User:
		newPass = u[tx.Statement.CurDestIndex].Password
	}
	b, err := bcrypt.GenerateFromPassword([]byte(newPass), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	tx.Statement.SetColumn("password", b)
	return nil
}
