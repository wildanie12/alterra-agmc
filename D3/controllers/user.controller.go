package controllers

import (
	"agmc_d2/lib"
	"agmc_d2/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)


type UserController struct {
	repo *lib.UserRepository
}

func NewUser(repo *lib.UserRepository) *UserController {
	return &UserController{
		repo: repo,
	}
}

func (uc *UserController) Index(c echo.Context) error {
	users, err := uc.repo.FindAll()
	if err != nil {
		return c.JSON(http.StatusOK, map[string]interface{} {
			"status": "error",
			"code": http.StatusInternalServerError,
			"message": "error creating a user: " + err.Error(),
			"data": nil,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{} {
		"status": "success",
		"code": http.StatusOK,
		"message": "Success getting a list of users",
		"data": users,
	})
} 

func (uc *UserController) Show(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		return c.JSON(http.StatusOK, map[string]interface{} {
			"status": "error",
			"code": http.StatusBadRequest,
			"message": "userID is invalid",
			"data": nil,
		})
	}

	user, err := uc.repo.Find(id)
	if err != nil {
		code := http.StatusInternalServerError
		if err.Error() == "cannot get user detail: record not found" {
			code = http.StatusBadRequest
		}
		return c.JSON(http.StatusOK, map[string]interface{} {
			"status": "error",
			"code": code,
			"message": "error finding a user: " + err.Error(),
			"data": nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{} {
		"status": "success",
		"code": http.StatusOK,
		"message": "Success getting a user",
		"data": user,
	})
} 

func (uc *UserController) Store(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	user, err := uc.repo.Create(user)
	if err != nil {
		return c.JSON(http.StatusOK, map[string]interface{} {
			"status": "error",
			"code": http.StatusInternalServerError,
			"message": "error creating a user: " + err.Error(),
			"data": nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{} {
		"status": "success",
		"code": http.StatusOK,
		"message": "Success creating a user",
		"data": user,
	})
} 

func (uc *UserController) Update(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	id, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		return c.JSON(http.StatusOK, map[string]interface{} {
			"status": "error",
			"code": http.StatusBadRequest,
			"message": "userID is invalid",
			"data": nil,
		})
	}

	user, err = uc.repo.Update(user, id)
	if err != nil {
		code := http.StatusInternalServerError
		if err.Error() == "cannot get user detail: record not found" {
			code = http.StatusBadRequest
		}
		return c.JSON(http.StatusOK, map[string]interface{} {
			"status": "error",
			"code": code,
			"message": "error updating a user: " + err.Error(),
			"data": nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{} {
		"status": "success",
		"code": http.StatusOK,
		"message": "Success updating a user",
		"data": map[string]interface{} {
			"id": id,
		},
	})
} 

func (uc *UserController) Delete(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		return c.JSON(http.StatusOK, map[string]interface{} {
			"status": "error",
			"code": http.StatusBadRequest,
			"message": "userID is invalid",
			"data": nil,
		})
	}

	err = uc.repo.Delete(id)
	if err != nil {
		code := http.StatusInternalServerError
		if err.Error() == "cannot get user detail: record not found" {
			code = http.StatusBadRequest
		}
		return c.JSON(http.StatusOK, map[string]interface{} {
			"status": "error",
			"code": code,
			"message": "error deleting a user: " + err.Error(),
			"data": nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{} {
		"status": "success",
		"code": http.StatusOK,
		"message": "Success deleting a user",
		"data": map[string]interface{} {
			"id": id,
		},
	})
} 