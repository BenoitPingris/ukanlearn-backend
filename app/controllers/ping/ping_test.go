package ping_test

import (
	"net/http"
	"testing"
	"ukanlearn/app/test"

	"github.com/stretchr/testify/suite"
)

type PingSuite struct {
	test.Suite
}

func (p *PingSuite) TestGet() {
	req, _ := http.NewRequest("GET", "/ping", nil)
	p.Router.ServeHTTP(p.Response, req)
	p.Equal(http.StatusOK, p.Response.Code)
	p.Equal("Pong!", p.Response.Body.String())
}

func TestPingSuite(t *testing.T) {
	suite.Run(t, &PingSuite{})
}
