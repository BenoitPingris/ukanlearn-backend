package auth

import (
	"net/http"
	"ukanlearn/app/models"

	"github.com/BenoitPingris/validation-request"
	"gorm.io/gorm"
)

// Controller struct
type Controller struct {
	DB *gorm.DB
}

// New returns a new controller
func New(db *gorm.DB) *Controller {
	return &Controller{db}
}

// Register handler
func (c *Controller) Register(w http.ResponseWriter, r *http.Request) {
	body := validation.FromContext(r.Context()).(*RegisterRequest)

	var tmp models.User
	res := c.DB.First(&tmp, "email = ?", body.Email)
	if res.RowsAffected > 0 {
		http.Error(w, "Email already registered.", http.StatusBadRequest)
		return
	}
	c.DB.Create(&models.User{Email: body.Email, Password: body.Password})
	w.Write([]byte("ok"))
}

// Login handler
func (c *Controller) Login(w http.ResponseWriter, r *http.Request) {
	body := validation.FromContext(r.Context()).(*LoginRequest)

	var u models.User
	res := c.DB.First(&u, "email = ?", body.Email)
	if res.RowsAffected == 0 {
		http.Error(w, "Email not found.", http.StatusNotFound)
		return
	}
	if err := u.ComparePasswords(body.Password); err != nil {
		http.Error(w, "Invalid credentials.", http.StatusBadRequest)
		return
	}
	w.Write([]byte("ok"))
}
