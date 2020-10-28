package app

import (
	"net/http"

	"gorm.io/gorm"
)

func Init(db *gorm.DB) {
	r := SetupRouter(db)
	http.ListenAndServe(":3001", r)
}
