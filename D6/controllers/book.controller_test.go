package controllers_test

import (
	"agmc_d6/controllers"
	"agmc_d6/routes"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestBookIndex(t *testing.T) {
	// setup
	e := echo.New()
	
	req := httptest.NewRequest(http.MethodGet, "/v1/books", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	
	bc := controllers.NewBook()
	err := bc.Index(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}


func TestBookCreate(t *testing.T) {
	
	// setup
	e := echo.New()

	reqBody := map[string]interface{} {
		"title": "Title #1",
		"author": "Author  #1",
		"category": "Category #1",
		"year": "2000",
		"stock": 22,
	}

	reqBodyStr, _ := json.Marshal(reqBody)
	
	req := httptest.NewRequest(http.MethodPost, "/v1/books", bytes.NewBuffer(reqBodyStr))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	
	bc := controllers.NewBook()
	err := bc.Store(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestBookShow(t *testing.T) {
	
	// setup
	e := echo.New()
	routes.SetRouter(e)
	bc := controllers.NewBook()

	t.Run("invalid format", func(t *testing.T) {		
		req := httptest.NewRequest(http.MethodPut, "/v1/books", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("bookID")
		c.SetParamValues("salah")
		
		err := bc.Show(c)

		assert.NoError(t, err)
		// Empty data code
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("failed not found", func(t *testing.T) {		
		req := httptest.NewRequest(http.MethodPut, "/v1/books", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("bookID")
		c.SetParamValues("12")
		
		err := bc.Show(c)

		fmt.Println("----   INI-- ", rec.Body.String())

		assert.NoError(t, err)
		// Empty data code
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("success", func(t *testing.T) {
		// Create sample by API
		reqBody := map[string]interface{} {
			"title": "Title #1",
			"author": "Author  #1",
			"category": "Category #1",
			"year": "2000",
			"stock": 22,
		}
		reqBodyStr, _ := json.Marshal(reqBody)
		req := httptest.NewRequest(http.MethodPost, "/v1/books", bytes.NewBuffer(reqBodyStr))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		err := bc.Store(e.NewContext(req, rec))
		if err != nil {
			panic("cannot create sample by api")
		}

		// Test
		req = httptest.NewRequest(http.MethodPut, "/v1/books", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		
		rec = httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("bookID")
		c.SetParamValues("1")

		
		err = bc.Show(c)
		mapRes := map[string]json.RawMessage{}
		json.Unmarshal(rec.Body.Bytes(), &mapRes)
		fmt.Println(string(mapRes["message"]))
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}

func TestBookUpdate(t *testing.T) {
	e := echo.New()
	routes.SetRouter(e)
	bc := controllers.NewBook()

	t.Run("success", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodPost, "/v1/books", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		err := bc.Store(e.NewContext(req, rec))
		if err != nil {
			panic("cannot create sample by API")
		}

		req = httptest.NewRequest(http.MethodPut, "/v1/books", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec = httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("bookID")
		c.SetParamValues("0")
		err = bc.Update(c)
		
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
	})
	t.Run("invalid param id type", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/v1/books", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("bookID")
		c.SetParamValues("asdasd")
		err := bc.Update(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest ,rec.Code)
	})
	t.Run("invalid param id", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/v1/books", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("bookID")
		c.SetParamValues("5000")
		err := bc.Update(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest ,rec.Code)
	})
}

func TestBookDelete(t *testing.T) {
	e := echo.New()
	routes.SetRouter(e)
	bc := controllers.NewBook()

	t.Run("success", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodPost, "/v1/books", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		err := bc.Store(e.NewContext(req, rec))
		if err != nil {
			panic("cannot create sample by API")
		}

		req = httptest.NewRequest(http.MethodDelete, "/v1/books", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec = httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("bookID")
		c.SetParamValues("0")
		err = bc.Delete(c)
		
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
	})
	t.Run("invalid param id type", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/v1/books", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("bookID")
		c.SetParamValues("asdasd")
		err := bc.Delete(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest ,rec.Code)
	})
	t.Run("invalid param id", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/v1/books", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("bookID")
		c.SetParamValues("5000")
		err := bc.Delete(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest ,rec.Code)
	})
}