package app

import (
	"ukanlearn/app/controllers/auth"
	"ukanlearn/app/controllers/ping"

	"github.com/BenoitPingris/validation-request"
	"github.com/go-chi/chi"
	"gorm.io/gorm"
)

// SetupRouter setups all the route of the application
func SetupRouter(db *gorm.DB) *chi.Mux {
	r := chi.NewRouter()
	authCtrl := auth.New(db)

	r.Get("/ping", ping.Ping)
	r.With(validation.Validate(auth.LoginRequest{})).Post("/login", authCtrl.Login)
	r.With(validation.Validate(auth.RegisterRequest{})).Post("/register", authCtrl.Register)
	return r
}
