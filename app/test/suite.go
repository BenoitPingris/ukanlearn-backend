package test

import (
	"log"
	"net/http/httptest"
	"ukanlearn/app"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Suite struct to manage test with testify
type Suite struct {
	suite.Suite
	Router   *chi.Mux
	Response *httptest.ResponseRecorder
}

func setupDbOrDie() *gorm.DB {
	dsn := "host=localhost user=gorm password=gorm dbname=db_test port=5432 sslmode=disable TimeZone=Europe/Paris"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}

// SetupSuite called before testing
func (s *Suite) SetupSuite() {
	db := setupDbOrDie()
	s.Router = app.SetupRouter(db)
}

// SetupTest called before each test
func (s *Suite) SetupTest() {
	s.Response = httptest.NewRecorder()
}
