package controllers_test

import (
	"agmc_d6/controllers"
	"agmc_d6/database"
	"agmc_d6/repositories"
	"agmc_d6/routes"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var db *gorm.DB
var ur *repositories.UserRepository
func init() {
	godotenv.Load("../.env")
	db = database.NewMySQL()
	ur = repositories.NewUserRepository(db)
}

func TestUserIndex(t *testing.T) {
	e := echo.New()
	routes.SetRouter(e)

	uc := controllers.NewUser(ur)
	t.Run("success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/v1/users", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		err := uc.Index(e.NewContext(req, rec))
		
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}

func TestUserStore(t *testing.T) {
	e := echo.New()
	routes.SetRouter(e)
	uc := controllers.NewUser(ur)

	t.Run("success", func(t *testing.T) {
		body := map[string]string {
			"name": "Test #1",
			"email": "mailtest@mail.com",
			"address": "jakarta",
			"gender": "L",
			"phone_number": "123123123",
		}
		bodyStr, _ := json.Marshal(body)
		req := httptest.NewRequest(http.MethodPost, "/v1/users", bytes.NewBuffer(bodyStr))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		err := uc.Store(e.NewContext(req, rec))

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
	})	
}

func TestUserShow(t *testing.T) {
	e := echo.New()
	routes.SetRouter(e)
	uc := controllers.NewUser(ur)

	t.Run("success", func(t *testing.T) {
		body := map[string]string {
			"name": "Test #1",
			"email": "mailtest@mail.com",
			"address": "jakarta",
			"gender": "L",
			"phone_number": "123123123",
		}
		bodyStr, _ := json.Marshal(body)
		req := httptest.NewRequest(http.MethodPost, "/v1/users", bytes.NewBuffer(bodyStr))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		err := uc.Store(e.NewContext(req, rec))
		if err != nil {
			panic("error creating sample by API")
		}
		
		res := map[string]json.RawMessage{}
		json.Unmarshal(rec.Body.Bytes(), &res)
		data := map[string]json.RawMessage{}
		json.Unmarshal(res["data"], &data)
		id := string(data["id"])
		
		body = map[string]string {
			"name": "Test Updated #1",
			"email": "mailtest@mail.com",
			"address": "jakarta",
			"gender": "L",
			"phone_number": "123123123",
		}
		bodyStr, _ = json.Marshal(body)
		req = httptest.NewRequest(http.MethodGet, "/v1/users", bytes.NewBuffer(bodyStr))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec = httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("userID")
		c.SetParamValues(id)
		err = uc.Show(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}

func TestUserUpdate(t *testing.T) {
	e := echo.New()
	routes.SetRouter(e)
	uc := controllers.NewUser(ur)

	t.Run("success", func(t *testing.T) {
		body := map[string]string {
			"name": "Test #1",
			"email": "mailtest@mail.com",
			"address": "jakarta",
			"gender": "L",
			"phone_number": "123123123",
		}
		bodyStr, _ := json.Marshal(body)
		req := httptest.NewRequest(http.MethodPost, "/v1/users", bytes.NewBuffer(bodyStr))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		err := uc.Store(e.NewContext(req, rec))
		if err != nil {
			panic("error creating sample by API")
		}
		
		res := map[string]json.RawMessage{}
		json.Unmarshal(rec.Body.Bytes(), &res)
		data := map[string]json.RawMessage{}
		json.Unmarshal(res["data"], &data)
		id := string(data["id"])
		
		body = map[string]string {
			"name": "Test Updated #1",
			"email": "mailtest@mail.com",
			"address": "jakarta",
			"gender": "L",
			"phone_number": "123123123",
		}
		bodyStr, _ = json.Marshal(body)
		req = httptest.NewRequest(http.MethodPut, "/v1/users", bytes.NewBuffer(bodyStr))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec = httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("userID")
		c.SetParamValues(id)
		err = uc.Update(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}


func TestUserDelete(t *testing.T) {
	e := echo.New()
	routes.SetRouter(e)
	uc := controllers.NewUser(ur)

	t.Run("success", func(t *testing.T) {
		body := map[string]string {
			"name": "Test #1",
			"email": "mailtest@mail.com",
			"address": "jakarta",
			"gender": "L",
			"phone_number": "123123123",
		}
		bodyStr, _ := json.Marshal(body)
		req := httptest.NewRequest(http.MethodPost, "/v1/users", bytes.NewBuffer(bodyStr))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		err := uc.Store(e.NewContext(req, rec))
		if err != nil {
			panic("error creating sample by API")
		}
		
		res := map[string]json.RawMessage{}
		json.Unmarshal(rec.Body.Bytes(), &res)
		data := map[string]json.RawMessage{}
		json.Unmarshal(res["data"], &data)
		id := string(data["id"])
		
		body = map[string]string {
			"name": "Test Updated #1",
			"email": "mailtest@mail.com",
			"address": "jakarta",
			"gender": "L",
			"phone_number": "123123123",
		}
		bodyStr, _ = json.Marshal(body)
		req = httptest.NewRequest(http.MethodDelete, "/v1/users", bytes.NewBuffer(bodyStr))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec = httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("userID")
		c.SetParamValues(id)
		err = uc.Delete(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}