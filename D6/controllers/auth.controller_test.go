package controllers_test

import (
	"agmc_d4/controllers"
	"agmc_d4/database"
	"agmc_d4/lib"
	"agmc_d4/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joho/godotenv"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func init() {
	godotenv.Load("../.env")
}


func TestLogin(t *testing.T) {

	// sample user
	sampleUser := map[string]interface{} {
		"email": "badar.wildanie@gmail.com",
		"password": "qweqweqwe",
	}
	
	// setup
	e := echo.New()
	db := database.NewMySQL()
	userRepo := lib.NewUserRepository(db)

	_, err := userRepo.FindByUsername(sampleUser["email"].(string))
	if err != nil {
		_, err := userRepo.Create(models.User{
			Name: "Badar Wildani",
			Email: sampleUser["email"].(string),
			Gender: "L",
			Address: "Jakarta",
			PhoneNumber: "082123123123",
		})
		if err != nil {
			panic("error creating samples")
		}
	}

	t.Run("success login", func(t *testing.T) {

		bodyStr, _ := json.Marshal(sampleUser)
		req := httptest.NewRequest(http.MethodPost, "/v1/auth/login", bytes.NewBuffer(bodyStr))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		authController := controllers.NewAuth(userRepo)

		err = authController.Login(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("failed login", func(t *testing.T) {

		bodyStr, _ := json.Marshal(map[string]string{
			"email": "salah@gmail.com",
			"password": "qweqweqwe",
		})
		req := httptest.NewRequest(http.MethodPost, "/v1/auth/login", bytes.NewBuffer(bodyStr))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		
		authController := controllers.NewAuth(userRepo)
		err := authController.Login(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
	
	
}