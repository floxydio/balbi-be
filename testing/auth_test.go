package testing

import (
	"balbibe/handler"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var jsonRegister = `{"nama" : "Dio Rovelino", "email" : "d@d.com", "username" : "dio", "password" : "12345"}`

func TestRegister(t *testing.T) {
	e := echo.New()

	req, err := http.NewRequest(http.MethodPost, "/register", strings.NewReader(jsonRegister))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	res := rec.Result()
	defer res.Body.Close()

	if assert.NoError(t, handler.SignUp(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "["+jsonRegister+"]", rec.Body.String())
	}
}
